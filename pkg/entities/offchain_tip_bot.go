package entities

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/defipod/mochi/pkg/consts"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
)

func (e *Entity) TransferToken(req request.OffchainTransferRequest) ([]response.OffchainTipBotTransferToken, error) {
	// check supported tokens
	amountEachRecipient := req.Amount / float64(len(req.Recipients))
	supportedToken, err := e.repo.OffchainTipBotTokens.GetBySymbol(strings.ToUpper(req.Token))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			e.repo.OffchainTipBotActivityLogs.CreateActivityLog(&model.OffchainTipBotActivityLog{
				UserID:          req.Sender,
				GuildID:         req.GuildID,
				ChannelID:       req.ChannelID,
				Action:          &req.TransferType,
				Receiver:        req.Recipients,
				NumberReceivers: len(req.Recipients),
				Duration:        &req.Duration,
				Amount:          amountEachRecipient,
				Status:          consts.OffchainTipBotTrasferStatusFail,
				FullCommand:     &req.FullCommand,
				FailReason:      consts.OffchainTipBotFailReasonTokenNotSupported,
			})
			return []response.OffchainTipBotTransferToken{}, errors.New(consts.OffchainTipBotFailReasonTokenNotSupported)
		}
		e.log.Fields(logger.Fields{"token": req.Token}).Error(err, "[repo.OffchainTipBotTokens.GetBySymbol] - failed to get check supported token")
		return nil, err
	}

	tokenPrice, err := e.svc.CoinGecko.GetCoinPrice([]string{supportedToken.CoinGeckoID}, "usd")
	if err != nil {
		e.log.Fields(logger.Fields{"token": supportedToken.CoinGeckoID}).Error(err, "[svc.CoinGecko.GetCoinPrice] - failed to get coin price from Coingecko")
		return nil, err
	}

	// check user bals, both not have record user_bals + amount in record user_bals = 0 -> return not enough bals
	modelNotEnoughBalance := &model.OffchainTipBotActivityLog{
		UserID:          req.Sender,
		GuildID:         req.GuildID,
		ChannelID:       req.ChannelID,
		Action:          &req.TransferType,
		Receiver:        req.Recipients,
		NumberReceivers: len(req.Recipients),
		TokenID:         supportedToken.ID.String(),
		Duration:        &req.Duration,
		Amount:          amountEachRecipient,
		Status:          consts.OffchainTipBotTrasferStatusFail,
		FullCommand:     &req.FullCommand,
		FailReason:      consts.OffchainTipBotFailReasonNotEnoughBalance,
	}
	userBal, err := e.repo.OffchainTipBotUserBalances.GetUserBalanceByTokenID(req.Sender, supportedToken.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			e.repo.OffchainTipBotActivityLogs.CreateActivityLog(modelNotEnoughBalance)
			return []response.OffchainTipBotTransferToken{}, errors.New(consts.OffchainTipBotFailReasonNotEnoughBalance)
		}
		e.log.Fields(logger.Fields{"token": req.Token, "user": req.Sender}).Error(err, "[repo.OffchainTipBotUserBalances.GetUserBalanceByTokenID] - failed to get user balance")
		return nil, err
	}

	if req.All {
		if userBal.Amount == 0 {
			e.repo.OffchainTipBotActivityLogs.CreateActivityLog(modelNotEnoughBalance)
			return []response.OffchainTipBotTransferToken{}, errors.New(consts.OffchainTipBotFailReasonNotEnoughBalance)
		}
		req.Amount = userBal.Amount
		amountEachRecipient = req.Amount / float64(len(req.Recipients))
	}

	if float64(userBal.Amount) < req.Amount {
		e.repo.OffchainTipBotActivityLogs.CreateActivityLog(modelNotEnoughBalance)
		return []response.OffchainTipBotTransferToken{}, errors.New(consts.OffchainTipBotFailReasonNotEnoughBalance)
	}

	// case record offchain_tip_bot_user_balances if not exist yet, CreateIfNotExists here to ensure that
	// TODO(trkhoi): gorm not support upsert batch model, temp do like this will find another way later
	for _, recipient := range req.Recipients {
		e.repo.OffchainTipBotUserBalances.CreateIfNotExists(&model.OffchainTipBotUserBalance{
			UserID:  recipient,
			TokenID: supportedToken.ID,
		})
	}

	// create activity log
	al, err := e.repo.OffchainTipBotActivityLogs.CreateActivityLog(&model.OffchainTipBotActivityLog{
		UserID:          req.Sender,
		GuildID:         req.GuildID,
		ChannelID:       req.ChannelID,
		Action:          &req.TransferType,
		Receiver:        req.Recipients,
		TokenID:         supportedToken.ID.String(),
		NumberReceivers: len(req.Recipients),
		Duration:        &req.Duration,
		Amount:          amountEachRecipient,
		Status:          consts.OffchainTipBotTrasferStatusSuccess,
		FullCommand:     &req.FullCommand,
		FailReason:      "",
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[repo.OffchainTipBotActivityLogs.CreateActivityLog] - failed to create activity log")
		return nil, err
	}

	// create transfer histories for each transfer
	listTransferHistories := make([]model.OffchainTipBotTransferHistory, 0)
	for _, recipient := range req.Recipients {
		listTransferHistories = append(listTransferHistories, model.OffchainTipBotTransferHistory{
			SenderID:   req.Sender,
			ReceiverID: recipient,
			GuildID:    req.GuildID,
			LogID:      al.ID.String(),
			Status:     consts.OffchainTipBotTrasferStatusSuccess,
			Amount:     amountEachRecipient,
			Token:      supportedToken.TokenSymbol,
			Action:     req.TransferType,
		})
	}
	transferHistories, err := e.repo.OffchainTipBotTransferHistories.CreateTransferHistories(listTransferHistories)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[repo.OffchainTipBotTransferHistories.CreateTransferHistories] - failed to create transfer histories")
		return nil, err
	}

	// update recipients balances
	err = e.repo.OffchainTipBotUserBalances.UpdateListUserBalances(req.Recipients, supportedToken.ID, amountEachRecipient)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[repo.OffchainTipBotUserBalances.UpdateListUserBalances] - failed to update recipients balances")
		return nil, err
	}

	// update sender balanace
	err = e.repo.OffchainTipBotUserBalances.UpdateUserBalance(&model.OffchainTipBotUserBalance{UserID: req.Sender, TokenID: supportedToken.ID, Amount: userBal.Amount - req.Amount})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[repo.OffchainTipBotUserBalances.UpdateUserBalance] - failed to update sender balance")
		return nil, err
	}

	return e.MappingTransferTokenResponse(req.Token, amountEachRecipient, tokenPrice[supportedToken.CoinGeckoID], transferHistories), nil
}

func (e *Entity) MappingTransferTokenResponse(tokenSymbol string, amount float64, price float64, transferHistories []model.OffchainTipBotTransferHistory) (res []response.OffchainTipBotTransferToken) {
	for _, transferHistory := range transferHistories {
		res = append(res, response.OffchainTipBotTransferToken{
			SenderID:    transferHistory.SenderID,
			RecipientID: transferHistory.ReceiverID,
			Amount:      amount,
			Symbol:      tokenSymbol,
			AmountInUSD: amount * price,
		})
	}
	return res
}

func (e *Entity) OffchainTipBotWithdraw(req request.OffchainWithdrawRequest) (*response.OffchainTipBotWithdraw, error) {
	// check supported tokens
	offchainToken, err := e.repo.OffchainTipBotTokens.GetBySymbol(req.Token)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			e.repo.OffchainTipBotActivityLogs.CreateActivityLog(&model.OffchainTipBotActivityLog{
				UserID:          req.Recipient,
				GuildID:         req.GuildID,
				ChannelID:       req.ChannelID,
				Action:          &req.TransferType,
				Receiver:        []string{req.Recipient},
				NumberReceivers: 1,
				Duration:        &req.Duration,
				Amount:          req.Amount,
				Status:          consts.OffchainTipBotTrasferStatusFail,
				FullCommand:     &req.FullCommand,
				FailReason:      consts.OffchainTipBotFailReasonTokenNotSupported,
			})
			return nil, errors.New(consts.OffchainTipBotFailReasonTokenNotSupported)
		}
		e.log.Fields(logger.Fields{"token": req.Token}).Error(err, "[repo.OffchainTipBotTokens.GetBySymbol] - failed to get check supported token")
		return nil, err
	}

	// check recipient balance
	modelNotEnoughBalance := &model.OffchainTipBotActivityLog{
		UserID:          req.Recipient,
		GuildID:         req.GuildID,
		ChannelID:       req.ChannelID,
		Action:          &req.TransferType,
		Receiver:        []string{req.Recipient},
		NumberReceivers: 1,
		TokenID:         offchainToken.ID.String(),
		Duration:        &req.Duration,
		Amount:          req.Amount,
		Status:          consts.OffchainTipBotTrasferStatusFail,
		FullCommand:     &req.FullCommand,
		FailReason:      consts.OffchainTipBotFailReasonNotEnoughBalance,
		ServiceFee:      offchainToken.ServiceFee,
		FeeAmount:       offchainToken.ServiceFee * req.Amount,
	}
	recipientBal, err := e.repo.OffchainTipBotUserBalances.GetUserBalanceByTokenID(req.Recipient, offchainToken.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			e.repo.OffchainTipBotActivityLogs.CreateActivityLog(modelNotEnoughBalance)
			return nil, errors.New(consts.OffchainTipBotFailReasonNotEnoughBalance)
		}
		e.log.Fields(logger.Fields{"token": req.Token, "user": req.Recipient}).Error(err, "[repo.OffchainTipBotUserBalances.GetUserBalanceByTokenID] - failed to get user balance")
		return nil, err
	}

	if float64(recipientBal.Amount) < req.Amount+offchainToken.ServiceFee*req.Amount {
		e.repo.OffchainTipBotActivityLogs.CreateActivityLog(modelNotEnoughBalance)
		return nil, errors.New(consts.OffchainTipBotFailReasonNotEnoughBalance)
	}

	// temp get from old tokens because not have flow from coingecko to offchain_tip_bot_tokens yet
	token, err := e.repo.Token.GetBySymbol(strings.ToLower(req.Token), true)
	if err != nil {
		return nil, err
	}

	// execute tx
	signedTx, transferredAmount, err := e.transferOffchain(req.Amount,
		accounts.Account{Address: common.HexToAddress(req.RecipientAddress)},
		req.Amount,
		token, -1, req.All)
	if err != nil {
		err = fmt.Errorf("error transfer: %v", err)
		return nil, err
	}

	// execute tx success -> create offchain_tip_bot_activity_logs + offchain_tip_bot_transfer_histories and update offchain_tip_bot_user_balances
	al, err := e.repo.OffchainTipBotActivityLogs.CreateActivityLog(&model.OffchainTipBotActivityLog{
		UserID:          req.Recipient,
		GuildID:         req.GuildID,
		ChannelID:       req.ChannelID,
		Action:          &req.TransferType,
		Receiver:        []string{req.RecipientAddress},
		TokenID:         offchainToken.ID.String(),
		NumberReceivers: 1,
		Duration:        &req.Duration,
		Amount:          transferredAmount,
		Status:          consts.OffchainTipBotTrasferStatusSuccess,
		FullCommand:     &req.FullCommand,
		FailReason:      "",
		ServiceFee:      offchainToken.ServiceFee,
		FeeAmount:       offchainToken.ServiceFee * req.Amount,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[repo.OffchainTipBotActivityLogs.CreateActivityLog] - failed to create activity log")
		return nil, err
	}

	_, err = e.repo.OffchainTipBotTransferHistories.CreateTransferHistories([]model.OffchainTipBotTransferHistory{{
		SenderID:   req.Recipient,
		ReceiverID: req.Recipient,
		GuildID:    req.GuildID,
		LogID:      al.ID.String(),
		Status:     consts.OffchainTipBotTrasferStatusSuccess,
		Amount:     transferredAmount,
		Token:      offchainToken.TokenSymbol,
		Action:     req.TransferType,
		ServiceFee: offchainToken.ServiceFee,
		FeeAmount:  offchainToken.ServiceFee * req.Amount,
	}})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[repo.OffchainTipBotTransferHistories.CreateTransferHistories] - failed to create transfer histories")
		return nil, err
	}

	err = e.repo.OffchainTipBotUserBalances.UpdateUserBalance(&model.OffchainTipBotUserBalance{UserID: req.Recipient, TokenID: offchainToken.ID, Amount: recipientBal.Amount - transferredAmount - offchainToken.ServiceFee*req.Amount})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[repo.OffchainTipBotUserBalances.UpdateUserBalance] - failed to update sender balance")
		return nil, err
	}

	withdrawalAmount := util.WeiToEther(signedTx.Value())
	transactionFee, _ := util.WeiToEther(new(big.Int).Sub(signedTx.Cost(), signedTx.Value())).Float64()

	return &response.OffchainTipBotWithdraw{
		UserDiscordID:  req.Recipient,
		ToAddress:      req.RecipientAddress,
		Amount:         transferredAmount,
		Symbol:         req.Token,
		TxHash:         signedTx.Hash().Hex(),
		TxUrl:          fmt.Sprintf("%s/%s", token.Chain.TxBaseURL, signedTx.Hash().Hex()),
		WithdrawAmount: withdrawalAmount,
		TransactionFee: transactionFee,
	}, nil
}

func (e *Entity) TotalBalances() ([]response.TotalBalances, error) {
	tokens, err := e.repo.Token.GetAll()
	if err != nil {
		e.log.Error(err, "[entities.migrateBalance] - failed to get supported tokens")
		return nil, err
	}
	balances, err := e.balances("0x4ec16127E879464bEF6ab310084FAcEC1E4Fe465", tokens)
	if err != nil {
		e.log.Error(err, "[entities.migrateBalance] - failed to get balance")
		return nil, err
	}

	totalBalances := make([]response.TotalBalances, 0)

	for key, value := range balances {
		coingeckoID := ""
		// select coingecko id
		for _, t := range tokens {
			if t.Symbol == key {
				coingeckoID = t.CoinGeckoID
			}
		}
		// get coingecko price
		var tokenPrice float64

		tokenPrices, err := e.svc.CoinGecko.GetCoinPrice([]string{coingeckoID}, "usd")
		if err != nil {
			e.log.Fields(logger.Fields{"token": key}).Error(err, "[svc.CoinGecko.GetCoinPrice] - failed to get coin price from Coingecko")
			return nil, err
		}
		if len(tokenPrices) > 0 {
			tokenPrice = tokenPrices[coingeckoID]
		} else {
			tokenPrice = 1.0
		}

		if value > 0 {
			totalBalances = append(totalBalances, response.TotalBalances{
				Symbol:      key,
				Amount:      value,
				AmountInUsd: value * tokenPrice,
			})
		}

	}

	return totalBalances, nil
}

func (e *Entity) TotalOffchainBalances() ([]response.TotalOffchainBalances, error) {
	// cal exist offchain balance
	totalOffchainBalances, err := e.repo.OffchainTipBotUserBalances.SumAmountByTokenId()
	if err != nil {
		e.log.Error(err, "[e.repo.TotalOffchainBalances] - failed to get total offchain balance")
		return nil, err
	}

	mappingTotalOffchainBalances := make([]response.TotalOffchainBalances, 0)

	for _, totalOffchainBalance := range totalOffchainBalances {
		// get coingecko price
		var tokenPrice float64
		tokenPrices, err := e.svc.CoinGecko.GetCoinPrice([]string{totalOffchainBalance.CoinGeckoId}, "usd")
		if err != nil {
			e.log.Fields(logger.Fields{"token": totalOffchainBalance.CoinGeckoId}).Error(err, "[svc.CoinGecko.GetCoinPrice] - failed to get coin price from Coingecko")
			return nil, err
		}
		if len(tokenPrices) > 0 {
			tokenPrice = tokenPrices[totalOffchainBalance.CoinGeckoId]
		} else {
			tokenPrice = 1.0
		}

		mappingTotalOffchainBalances = append(mappingTotalOffchainBalances, response.TotalOffchainBalances{
			Symbol:      totalOffchainBalance.TokenSymbol,
			Amount:      totalOffchainBalance.Total,
			AmountInUsd: totalOffchainBalance.Total * tokenPrice,
		})
	}
	return mappingTotalOffchainBalances, nil
}

func (e *Entity) TotalFee() ([]response.TotalFeeWithdraw, error) {
	return e.repo.OffchainTipBotTransferHistories.TotalFeeFromWithdraw()
}

func (e *Entity) UpdateTokenFee(req request.OffchainUpdateTokenFee) error {
	return e.repo.OffchainTipBotTokens.UpdateTokenFee(req.Symbol, req.ServiceFee)
}
