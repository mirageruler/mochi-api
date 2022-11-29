package entities

import (
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/response"
)

func (e *Entity) OffchainTipBotCreateAssignContract(ac *model.OffchainTipBotAssignContract) (userAssignedContract *model.OffchainTipBotAssignContract, err error) {
	return e.repo.OffchainTipBotContract.CreateAssignContract(ac)
}

func (e *Entity) OffchainTipBotDeleteExpiredAssignContract() (err error) {
	return e.repo.OffchainTipBotContract.DeleteExpiredAssignContract()
}

func (e *Entity) GetUserBalances(userID string) (bals []response.GetUserBalances, err error) {
	userBals, err := e.repo.OffchainTipBotUserBalances.GetUserBalances(userID)
	if err != nil {
		e.log.Fields(logger.Fields{"userID": userID}).Error(err, "[repo.OffchainTipBotUserBalances.GetUserBalances] - failed to get user balances")
		return []response.GetUserBalances{}, err
	}

	listCoinIDs := []string{}
	for _, userBal := range userBals {
		coinID := userBal.Token.CoinGeckoID
		listCoinIDs = append(listCoinIDs, coinID)
		bals = append(bals, response.GetUserBalances{
			ID:       coinID,
			Name:     userBal.Token.TokenName,
			Symbol:   userBal.Token.TokenSymbol,
			Balances: userBal.Amount,
		})

	}

	tokenPrices, err := e.svc.CoinGecko.GetCoinPrice(listCoinIDs, "usd")
	if err != nil {
		e.log.Fields(logger.Fields{"listCoinIDs": listCoinIDs}).Error(err, "[svc.CoinGecko.GetCoinPrice] - failed to get coin price from Coingecko")
		return []response.GetUserBalances{}, err
	}

	for i, bal := range bals {
		bals[i].RateInUSD = tokenPrices[bal.ID]
		bals[i].BalancesInUSD = tokenPrices[bal.ID] * bal.Balances
	}

	return bals, nil
}