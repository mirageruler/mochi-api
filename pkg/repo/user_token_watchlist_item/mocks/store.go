// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/repo/user_token_watchlist_item/store.go

// Package mock_usertokenwatchlistitem is a generated GoMock package.
package mock_usertokenwatchlistitem

import (
	reflect "reflect"

	model "github.com/defipod/mochi/pkg/model"
	usertokenwatchlistitem "github.com/defipod/mochi/pkg/repo/user_token_watchlist_item"
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

// Count mocks base method.
func (m *MockStore) Count(arg0 usertokenwatchlistitem.CountQuery) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockStoreMockRecorder) Count(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockStore)(nil).Count), arg0)
}

// Create mocks base method.
func (m *MockStore) Create(item *model.UserTokenWatchlistItem) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", item)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockStoreMockRecorder) Create(item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockStore)(nil).Create), item)
}

// Delete mocks base method.
func (m *MockStore) Delete(profileID, symbol string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", profileID, symbol)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockStoreMockRecorder) Delete(profileID, symbol interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockStore)(nil).Delete), profileID, symbol)
}

// List mocks base method.
func (m *MockStore) List(q usertokenwatchlistitem.UserWatchlistQuery) ([]model.UserTokenWatchlistItem, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", q)
	ret0, _ := ret[0].([]model.UserTokenWatchlistItem)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockStoreMockRecorder) List(q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockStore)(nil).List), q)
}