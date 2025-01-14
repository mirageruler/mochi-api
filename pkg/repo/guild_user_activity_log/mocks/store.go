// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/repo/guild_user_activity_log/store.go

// Package mock_guild_user_activity_log is a generated GoMock package.
package mock_guild_user_activity_log

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

// CreateOne mocks base method
func (m *MockStore) CreateOne(record model.GuildUserActivityLog) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOne", record)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOne indicates an expected call of CreateOne
func (mr *MockStoreMockRecorder) CreateOne(record interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOne", reflect.TypeOf((*MockStore)(nil).CreateOne), record)
}

// CreateBatch mocks base method
func (m *MockStore) CreateBatch(records []model.GuildUserActivityLog) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBatch", records)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBatch indicates an expected call of CreateBatch
func (mr *MockStoreMockRecorder) CreateBatch(records interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBatch", reflect.TypeOf((*MockStore)(nil).CreateBatch), records)
}

// UpdateInvalidRecords mocks base method
func (m *MockStore) UpdateInvalidRecords(userID, profileID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInvalidRecords", userID, profileID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateInvalidRecords indicates an expected call of UpdateInvalidRecords
func (mr *MockStoreMockRecorder) UpdateInvalidRecords(userID, profileID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInvalidRecords", reflect.TypeOf((*MockStore)(nil).UpdateInvalidRecords), userID, profileID)
}
