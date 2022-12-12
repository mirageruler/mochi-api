// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/repo/guild_config_level_role/store.go

// Package mock_guild_config_level_role is a generated GoMock package.
package mock_guild_config_level_role

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	model "github.com/defipod/mochi/pkg/model"
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

// DeleteOne mocks base method.
func (m *MockStore) DeleteOne(guildID string, level int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOne", guildID, level)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOne indicates an expected call of DeleteOne.
func (mr *MockStoreMockRecorder) DeleteOne(guildID, level interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOne", reflect.TypeOf((*MockStore)(nil).DeleteOne), guildID, level)
}

// GetByGuildID mocks base method.
func (m *MockStore) GetByGuildID(guildID string) ([]model.GuildConfigLevelRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByGuildID", guildID)
	ret0, _ := ret[0].([]model.GuildConfigLevelRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByGuildID indicates an expected call of GetByGuildID.
func (mr *MockStoreMockRecorder) GetByGuildID(guildID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByGuildID", reflect.TypeOf((*MockStore)(nil).GetByGuildID), guildID)
}

// GetByRoleID mocks base method.
func (m *MockStore) GetByRoleID(guildID, roleID string) (*model.GuildConfigLevelRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByRoleID", guildID, roleID)
	ret0, _ := ret[0].(*model.GuildConfigLevelRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByRoleID indicates an expected call of GetByRoleID.
func (mr *MockStoreMockRecorder) GetByRoleID(guildID, roleID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByRoleID", reflect.TypeOf((*MockStore)(nil).GetByRoleID), guildID, roleID)
}

// GetHighest mocks base method.
func (m *MockStore) GetHighest(guildID string, level int) (*model.GuildConfigLevelRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHighest", guildID, level)
	ret0, _ := ret[0].(*model.GuildConfigLevelRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHighest indicates an expected call of GetHighest.
func (mr *MockStoreMockRecorder) GetHighest(guildID, level interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHighest", reflect.TypeOf((*MockStore)(nil).GetHighest), guildID, level)
}

// UpsertOne mocks base method.
func (m *MockStore) UpsertOne(config model.GuildConfigLevelRole) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertOne", config)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertOne indicates an expected call of UpsertOne.
func (mr *MockStoreMockRecorder) UpsertOne(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertOne", reflect.TypeOf((*MockStore)(nil).UpsertOne), config)
}
