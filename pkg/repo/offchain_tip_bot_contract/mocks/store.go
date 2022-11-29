// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/repo/offchain_tip_bot_contract/store.go

// Package mock_offchain_tip_bot_contract is a generated GoMock package.
package mock_offchain_tip_bot_contract

import (
	reflect "reflect"

	model "github.com/defipod/mochi/pkg/model"
	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateAssignContract mocks base method.
func (m *MockStore) CreateAssignContract(ac *model.OffchainTipBotAssignContract) (*model.OffchainTipBotAssignContract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAssignContract", ac)
	ret0, _ := ret[0].(*model.OffchainTipBotAssignContract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAssignContract indicates an expected call of CreateAssignContract.
func (mr *MockStoreMockRecorder) CreateAssignContract(ac interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAssignContract", reflect.TypeOf((*MockStore)(nil).CreateAssignContract), ac)
}

// DeleteExpiredAssignContract mocks base method.
func (m *MockStore) DeleteExpiredAssignContract() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExpiredAssignContract")
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteExpiredAssignContract indicates an expected call of DeleteExpiredAssignContract.
func (mr *MockStoreMockRecorder) DeleteExpiredAssignContract() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExpiredAssignContract", reflect.TypeOf((*MockStore)(nil).DeleteExpiredAssignContract))
}

// GetByAddress mocks base method.
func (m *MockStore) GetByAddress(addr string) (model.OffchainTipBotContract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByAddress", addr)
	ret0, _ := ret[0].(model.OffchainTipBotContract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByAddress indicates an expected call of GetByAddress.
func (mr *MockStoreMockRecorder) GetByAddress(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByAddress", reflect.TypeOf((*MockStore)(nil).GetByAddress), addr)
}

// GetByChainID mocks base method.
func (m *MockStore) GetByChainID(id string) ([]model.OffchainTipBotContract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByChainID", id)
	ret0, _ := ret[0].([]model.OffchainTipBotContract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByChainID indicates an expected call of GetByChainID.
func (mr *MockStoreMockRecorder) GetByChainID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByChainID", reflect.TypeOf((*MockStore)(nil).GetByChainID), id)
}

// GetByID mocks base method.
func (m *MockStore) GetByID(id string) (model.OffchainTipBotContract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(model.OffchainTipBotContract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockStoreMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockStore)(nil).GetByID), id)
}