// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/service/ronin/service.go

// Package mock_ronin is a generated GoMock package.
package mock_ronin

import (
	ronin "github.com/defipod/mochi/pkg/service/ronin"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// GetAxsStakingAmount mocks base method
func (m *MockService) GetAxsStakingAmount(address string) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAxsStakingAmount", address)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAxsStakingAmount indicates an expected call of GetAxsStakingAmount
func (mr *MockServiceMockRecorder) GetAxsStakingAmount(address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAxsStakingAmount", reflect.TypeOf((*MockService)(nil).GetAxsStakingAmount), address)
}

// GetAxsPendingRewards mocks base method
func (m *MockService) GetAxsPendingRewards(address string) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAxsPendingRewards", address)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAxsPendingRewards indicates an expected call of GetAxsPendingRewards
func (mr *MockServiceMockRecorder) GetAxsPendingRewards(address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAxsPendingRewards", reflect.TypeOf((*MockService)(nil).GetAxsPendingRewards), address)
}

// GetRonStakingAmount mocks base method
func (m *MockService) GetRonStakingAmount(address string) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRonStakingAmount", address)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRonStakingAmount indicates an expected call of GetRonStakingAmount
func (mr *MockServiceMockRecorder) GetRonStakingAmount(address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRonStakingAmount", reflect.TypeOf((*MockService)(nil).GetRonStakingAmount), address)
}

// GetRonPendingRewards mocks base method
func (m *MockService) GetRonPendingRewards(address string) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRonPendingRewards", address)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRonPendingRewards indicates an expected call of GetRonPendingRewards
func (mr *MockServiceMockRecorder) GetRonPendingRewards(address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRonPendingRewards", reflect.TypeOf((*MockService)(nil).GetRonPendingRewards), address)
}

// GetLpPendingRewards mocks base method
func (m *MockService) GetLpPendingRewards(address string) (map[string]ronin.LpRewardData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLpPendingRewards", address)
	ret0, _ := ret[0].(map[string]ronin.LpRewardData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLpPendingRewards indicates an expected call of GetLpPendingRewards
func (mr *MockServiceMockRecorder) GetLpPendingRewards(address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLpPendingRewards", reflect.TypeOf((*MockService)(nil).GetLpPendingRewards), address)
}
