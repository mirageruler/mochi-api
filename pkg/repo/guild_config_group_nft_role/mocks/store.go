// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/repo/guild_config_group_nft_role/store.go

// Package mock_guildconfiggroupnftrole is a generated GoMock package.
package mock_guildconfiggroupnftrole

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

// Create mocks base method.
func (m *MockStore) Create(config model.GuildConfigGroupNFTRole) (*model.GuildConfigGroupNFTRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", config)
	ret0, _ := ret[0].(*model.GuildConfigGroupNFTRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockStoreMockRecorder) Create(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockStore)(nil).Create), config)
}

// Delete mocks base method.
func (m *MockStore) Delete(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockStoreMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockStore)(nil).Delete), id)
}

// GetByRoleID mocks base method.
func (m *MockStore) GetByRoleID(guildID, roleID string) (*model.GuildConfigGroupNFTRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByRoleID", guildID, roleID)
	ret0, _ := ret[0].(*model.GuildConfigGroupNFTRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByRoleID indicates an expected call of GetByRoleID.
func (mr *MockStoreMockRecorder) GetByRoleID(guildID, roleID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByRoleID", reflect.TypeOf((*MockStore)(nil).GetByRoleID), guildID, roleID)
}

// ListByGuildID mocks base method.
func (m *MockStore) ListByGuildID(guildID string) ([]model.GuildConfigGroupNFTRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByGuildID", guildID)
	ret0, _ := ret[0].([]model.GuildConfigGroupNFTRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByGuildID indicates an expected call of ListByGuildID.
func (mr *MockStoreMockRecorder) ListByGuildID(guildID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByGuildID", reflect.TypeOf((*MockStore)(nil).ListByGuildID), guildID)
}