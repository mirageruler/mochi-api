// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/repo/discord_guilds/store.go

// Package mock_discord_guilds is a generated GoMock package.
package mock_discord_guilds

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

// Gets mocks base method
func (m *MockStore) Gets() ([]model.DiscordGuild, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Gets")
	ret0, _ := ret[0].([]model.DiscordGuild)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Gets indicates an expected call of Gets
func (mr *MockStoreMockRecorder) Gets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Gets", reflect.TypeOf((*MockStore)(nil).Gets))
}

// GetByID mocks base method
func (m *MockStore) GetByID(id string) (*model.DiscordGuild, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(*model.DiscordGuild)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockStoreMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockStore)(nil).GetByID), id)
}

// CreateOrReactivate mocks base method
func (m *MockStore) CreateOrReactivate(guild model.DiscordGuild) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrReactivate", guild)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrReactivate indicates an expected call of CreateOrReactivate
func (mr *MockStoreMockRecorder) CreateOrReactivate(guild interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrReactivate", reflect.TypeOf((*MockStore)(nil).CreateOrReactivate), guild)
}

// Update mocks base method
func (m *MockStore) Update(guild *model.DiscordGuild) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", guild)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockStoreMockRecorder) Update(guild interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockStore)(nil).Update), guild)
}

// GetNonLeftGuilds mocks base method
func (m *MockStore) GetNonLeftGuilds() ([]model.DiscordGuild, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNonLeftGuilds")
	ret0, _ := ret[0].([]model.DiscordGuild)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNonLeftGuilds indicates an expected call of GetNonLeftGuilds
func (mr *MockStoreMockRecorder) GetNonLeftGuilds() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNonLeftGuilds", reflect.TypeOf((*MockStore)(nil).GetNonLeftGuilds))
}
