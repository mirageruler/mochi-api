package vaultwallet

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"

	"github.com/defipod/mochi/pkg/chain"
	"github.com/defipod/mochi/pkg/config"
	erc20contract "github.com/defipod/mochi/pkg/contracts/erc20"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/repo"
	"github.com/defipod/mochi/pkg/repo/pg"
	"github.com/defipod/mochi/pkg/service/mochipay"
)

type IVaultWallet interface {
	GetAccountByWalletNumber(i int) (accounts.Account, error)
	GetHDWallet() *hdwallet.Wallet
	Chain(chainID int) *chain.Chain
	GetPrivateKeyByAccount(acc accounts.Account) (string, error)
	Balance(token *mochipay.Token, account string) (balance *big.Int, err error)
}

type VauleWallet struct {
	log      logger.Logger
	cfg      config.Config
	repo     *repo.Repo
	hdwallet *hdwallet.Wallet
	chains   map[int]*chain.Chain
}

// New will return an instance of DiscordWallet struct
func New(cfg config.Config, l logger.Logger, s repo.Store) (*VauleWallet, error) {
	r := pg.NewRepo(s.DB())

	wallet, err := hdwallet.NewFromMnemonic(cfg.VaultMnemonic)
	if err != nil {
		return nil, err
	}

	chainMap := make(map[int]*chain.Chain)
	chains, err := r.Chain.GetAll()
	if err != nil {
		return nil, err
	}

	for _, c := range chains {
		if c.RPC == "" {
			continue
		}
		chain, err := chain.NewClient(&cfg, wallet, l, c.RPC, c.APIKey, c.APIBaseURL)
		if err != nil {
			return nil, err
		}
		chainMap[c.ID] = chain
	}

	return &VauleWallet{
		log:      l,
		cfg:      cfg,
		repo:     r,
		chains:   chainMap,
		hdwallet: wallet,
	}, nil
}

func (d *VauleWallet) GetAccountByWalletNumber(i int) (accounts.Account, error) {
	path := hdwallet.MustParseDerivationPath(fmt.Sprintf("m/44'/60'/0'/0/%d", i))
	return d.hdwallet.Derive(path, false)
}

func (d *VauleWallet) GetHDWallet() *hdwallet.Wallet {
	return d.hdwallet
}

func (d *VauleWallet) Chain(chainID int) *chain.Chain {
	chain, _ := d.chains[chainID]
	return chain
}

func (d *VauleWallet) GetPrivateKeyByAccount(acc accounts.Account) (string, error) {
	return d.hdwallet.PrivateKeyHex(acc)
}

func (d *VauleWallet) Balance(token *mochipay.Token, account string) (balance *big.Int, err error) {
	if token.Native {
		balance, err = d.nativeBalance(token, account)
	} else {
		balance, err = d.erc20Balance(token, account)
	}
	return
}

func (d *VauleWallet) nativeBalance(token *mochipay.Token, account string) (*big.Int, error) {
	client, err := ethclient.Dial(token.Chain.Rpc)
	if err != nil {
		return nil, nil
	}
	return client.BalanceAt(context.Background(), common.HexToAddress(account), nil)
}

func (d *VauleWallet) erc20Balance(token *mochipay.Token, account string) (*big.Int, error) {
	client, err := ethclient.Dial(token.Chain.Rpc)
	if err != nil {
		return nil, nil
	}

	instance, err := erc20contract.NewErc20(common.HexToAddress(token.Address), client)
	if err != nil {
		return nil, err
	}

	accountAddress := common.HexToAddress(account)
	return instance.BalanceOf(&bind.CallOpts{}, accountAddress)
}
