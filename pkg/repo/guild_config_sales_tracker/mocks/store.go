// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/repo/guild_config_sales_tracker/store.go

// Package mock_guild_config_sales_tracker is a generated GoMock package.
package mock_guild_config_sales_tracker

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

// GetByGuildID mocks base method
func (m *MockStore) GetByGuildID(guildID string) ([]model.GuildConfigSalesTracker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByGuildID", guildID)
	ret0, _ := ret[0].([]model.GuildConfigSalesTracker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByGuildID indicates an expected call of GetByGuildID
func (mr *MockStoreMockRecorder) GetByGuildID(guildID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByGuildID", reflect.TypeOf((*MockStore)(nil).GetByGuildID), guildID)
}

// Create mocks base method
func (m *MockStore) Create(config *model.GuildConfigSalesTracker) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", config)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockStoreMockRecorder) Create(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockStore)(nil).Create), config)
}

// GetAllSalesTrackerConfig mocks base method
func (m *MockStore) GetAllSalesTrackerConfig() ([]model.GuildConfigSalesTracker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSalesTrackerConfig")
	ret0, _ := ret[0].([]model.GuildConfigSalesTracker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSalesTrackerConfig indicates an expected call of GetAllSalesTrackerConfig
func (mr *MockStoreMockRecorder) GetAllSalesTrackerConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSalesTrackerConfig", reflect.TypeOf((*MockStore)(nil).GetAllSalesTrackerConfig))
}
