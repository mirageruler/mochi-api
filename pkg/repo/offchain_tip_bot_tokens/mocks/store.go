// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/repo/offchain_tip_bot_tokens/store.go

// Package mock_offchain_tip_bot_tokens is a generated GoMock package.
package mock_offchain_tip_bot_tokens

import (
	model "github.com/defipod/mochi/pkg/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// GetBySymbol mocks base method
func (m *MockStore) GetBySymbol(symbol string) (*model.OffchainTipBotToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBySymbol", symbol)
	ret0, _ := ret[0].(*model.OffchainTipBotToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBySymbol indicates an expected call of GetBySymbol
func (mr *MockStoreMockRecorder) GetBySymbol(symbol interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBySymbol", reflect.TypeOf((*MockStore)(nil).GetBySymbol), symbol)
}

// Create mocks base method
func (m *MockStore) Create(arg0 *model.OffchainTipBotToken) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockStoreMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockStore)(nil).Create), arg0)
}
