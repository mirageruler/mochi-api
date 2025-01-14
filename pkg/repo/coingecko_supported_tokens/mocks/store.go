// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/repo/coingecko_supported_tokens/store.go

// Package mock_coingeckosupportedtokens is a generated GoMock package.
package mock_coingeckosupportedtokens

import (
	model "github.com/defipod/mochi/pkg/model"
	coingecko_supported_tokens "github.com/defipod/mochi/pkg/repo/coingecko_supported_tokens"
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

// GetOne mocks base method
func (m *MockStore) GetOne(id string) (*model.CoingeckoSupportedTokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOne", id)
	ret0, _ := ret[0].(*model.CoingeckoSupportedTokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOne indicates an expected call of GetOne
func (mr *MockStoreMockRecorder) GetOne(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOne", reflect.TypeOf((*MockStore)(nil).GetOne), id)
}

// List mocks base method
func (m *MockStore) List(q coingecko_supported_tokens.ListQuery) ([]model.CoingeckoSupportedTokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", q)
	ret0, _ := ret[0].([]model.CoingeckoSupportedTokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockStoreMockRecorder) List(q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockStore)(nil).List), q)
}

// Upsert mocks base method
func (m *MockStore) Upsert(token *model.CoingeckoSupportedTokens) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", token)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upsert indicates an expected call of Upsert
func (mr *MockStoreMockRecorder) Upsert(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockStore)(nil).Upsert), token)
}
