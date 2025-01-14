// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/repo/chain/store.go

// Package mock_chain is a generated GoMock package.
package mock_chain

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

// GetAll mocks base method
func (m *MockStore) GetAll() ([]model.Chain, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]model.Chain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockStoreMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockStore)(nil).GetAll))
}

// GetByID mocks base method
func (m *MockStore) GetByID(id int) (model.Chain, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(model.Chain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockStoreMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockStore)(nil).GetByID), id)
}

// GetByShortName mocks base method
func (m *MockStore) GetByShortName(shortName string) (*model.Chain, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByShortName", shortName)
	ret0, _ := ret[0].(*model.Chain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByShortName indicates an expected call of GetByShortName
func (mr *MockStoreMockRecorder) GetByShortName(shortName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByShortName", reflect.TypeOf((*MockStore)(nil).GetByShortName), shortName)
}
