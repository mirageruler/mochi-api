// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/service/covalent/service.go

// Package mock_covalent is a generated GoMock package.
package mock_covalent

import (
	reflect "reflect"

	response "github.com/defipod/mochi/pkg/response"
	covalent "github.com/defipod/mochi/pkg/service/covalent"
	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// GetHistoricalTokenPrices mocks base method.
func (m *MockService) GetHistoricalTokenPrices(chainID int, currency, address string) (*response.HistoricalTokenPricesResponse, error, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHistoricalTokenPrices", chainID, currency, address)
	ret0, _ := ret[0].(*response.HistoricalTokenPricesResponse)
	ret1, _ := ret[1].(error)
	ret2, _ := ret[2].(int)
	return ret0, ret1, ret2
}

// GetHistoricalTokenPrices indicates an expected call of GetHistoricalTokenPrices.
func (mr *MockServiceMockRecorder) GetHistoricalTokenPrices(chainID, currency, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHistoricalTokenPrices", reflect.TypeOf((*MockService)(nil).GetHistoricalTokenPrices), chainID, currency, address)
}

// GetTokenBalances mocks base method.
func (m *MockService) GetTokenBalances(chainID int, address string, retry int) (*covalent.GetTokenBalancesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTokenBalances", chainID, address, retry)
	ret0, _ := ret[0].(*covalent.GetTokenBalancesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTokenBalances indicates an expected call of GetTokenBalances.
func (mr *MockServiceMockRecorder) GetTokenBalances(chainID, address, retry interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTokenBalances", reflect.TypeOf((*MockService)(nil).GetTokenBalances), chainID, address, retry)
}

// GetTransactionsByAddress mocks base method.
func (m *MockService) GetTransactionsByAddress(chainID int, address string, size, retry int) (*covalent.GetTransactionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionsByAddress", chainID, address, size, retry)
	ret0, _ := ret[0].(*covalent.GetTransactionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionsByAddress indicates an expected call of GetTransactionsByAddress.
func (mr *MockServiceMockRecorder) GetTransactionsByAddress(chainID, address, size, retry interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionsByAddress", reflect.TypeOf((*MockService)(nil).GetTransactionsByAddress), chainID, address, size, retry)
}
