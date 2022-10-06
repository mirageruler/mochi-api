// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/repo/guild_config_welcome_channel/store.go

// Package mock_guild_config_welcome_channel is a generated GoMock package.
package mock_guild_config_welcome_channel

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

// DeleteOne mocks base method.
func (m *MockStore) DeleteOne(config *model.GuildConfigWelcomeChannel) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOne", config)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOne indicates an expected call of DeleteOne.
func (mr *MockStoreMockRecorder) DeleteOne(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOne", reflect.TypeOf((*MockStore)(nil).DeleteOne), config)
}

// GetByGuildID mocks base method.
func (m *MockStore) GetByGuildID(guildID string) (*model.GuildConfigWelcomeChannel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByGuildID", guildID)
	ret0, _ := ret[0].(*model.GuildConfigWelcomeChannel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByGuildID indicates an expected call of GetByGuildID.
func (mr *MockStoreMockRecorder) GetByGuildID(guildID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByGuildID", reflect.TypeOf((*MockStore)(nil).GetByGuildID), guildID)
}

// UpsertOne mocks base method.
func (m *MockStore) UpsertOne(config *model.GuildConfigWelcomeChannel) (*model.GuildConfigWelcomeChannel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertOne", config)
	ret0, _ := ret[0].(*model.GuildConfigWelcomeChannel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsertOne indicates an expected call of UpsertOne.
func (mr *MockStoreMockRecorder) UpsertOne(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertOne", reflect.TypeOf((*MockStore)(nil).UpsertOne), config)
}
