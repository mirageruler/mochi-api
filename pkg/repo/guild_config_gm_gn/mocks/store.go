// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/repo/guild_config_gm_gn/store.go

// Package mock_guild_config_gm_gn is a generated GoMock package.
package mock_guild_config_gm_gn

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
func (m *MockStore) GetByGuildID(guildID string) (*model.GuildConfigGmGn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByGuildID", guildID)
	ret0, _ := ret[0].(*model.GuildConfigGmGn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByGuildID indicates an expected call of GetByGuildID
func (mr *MockStoreMockRecorder) GetByGuildID(guildID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByGuildID", reflect.TypeOf((*MockStore)(nil).GetByGuildID), guildID)
}

// GetAllByGuildID mocks base method
func (m *MockStore) GetAllByGuildID(guildID string) ([]model.GuildConfigGmGn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllByGuildID", guildID)
	ret0, _ := ret[0].([]model.GuildConfigGmGn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllByGuildID indicates an expected call of GetAllByGuildID
func (mr *MockStoreMockRecorder) GetAllByGuildID(guildID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByGuildID", reflect.TypeOf((*MockStore)(nil).GetAllByGuildID), guildID)
}

// GetLatestByGuildID mocks base method
func (m *MockStore) GetLatestByGuildID(guildID string) ([]model.GuildConfigGmGn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestByGuildID", guildID)
	ret0, _ := ret[0].([]model.GuildConfigGmGn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestByGuildID indicates an expected call of GetLatestByGuildID
func (mr *MockStoreMockRecorder) GetLatestByGuildID(guildID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestByGuildID", reflect.TypeOf((*MockStore)(nil).GetLatestByGuildID), guildID)
}

// UpsertOne mocks base method
func (m *MockStore) UpsertOne(config *model.GuildConfigGmGn) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertOne", config)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertOne indicates an expected call of UpsertOne
func (mr *MockStoreMockRecorder) UpsertOne(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertOne", reflect.TypeOf((*MockStore)(nil).UpsertOne), config)
}

// CreateOne mocks base method
func (m *MockStore) CreateOne(config *model.GuildConfigGmGn) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOne", config)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOne indicates an expected call of CreateOne
func (mr *MockStoreMockRecorder) CreateOne(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOne", reflect.TypeOf((*MockStore)(nil).CreateOne), config)
}
