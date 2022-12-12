// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	accounts "github.com/ethereum/go-ethereum/accounts"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	mock "github.com/stretchr/testify/mock"

	chain "github.com/defipod/mochi/pkg/chain"
)

// IDiscordWallet is an autogenerated mock type for the IDiscordWallet type
type IDiscordWallet struct {
	mock.Mock
}

// FTM provides a mock function with given fields:
func (_m *IDiscordWallet) FTM() chain.Chain {
	ret := _m.Called()

	var r0 chain.Chain
	if rf, ok := ret.Get(0).(func() chain.Chain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chain.Chain)
		}
	}

	return r0
}

// GetAccountByWalletNumber provides a mock function with given fields: i
func (_m *IDiscordWallet) GetAccountByWalletNumber(i int) (accounts.Account, error) {
	ret := _m.Called(i)

	var r0 accounts.Account
	if rf, ok := ret.Get(0).(func(int) accounts.Account); ok {
		r0 = rf(i)
	} else {
		r0 = ret.Get(0).(accounts.Account)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(i)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHDWallet provides a mock function with given fields:
func (_m *IDiscordWallet) GetHDWallet() *hdwallet.Wallet {
	ret := _m.Called()

	var r0 *hdwallet.Wallet
	if rf, ok := ret.Get(0).(func() *hdwallet.Wallet); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*hdwallet.Wallet)
		}
	}

	return r0
}
