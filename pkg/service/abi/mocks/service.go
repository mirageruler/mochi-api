// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/service/abi/service.go

// Package mock_abi is a generated GoMock package.
package mock_abi

import (
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

// GetNameAndSymbol mocks base method
func (m *MockService) GetNameAndSymbol(address string, chainId int64) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNameAndSymbol", address, chainId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetNameAndSymbol indicates an expected call of GetNameAndSymbol
func (mr *MockServiceMockRecorder) GetNameAndSymbol(address, chainId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNameAndSymbol", reflect.TypeOf((*MockService)(nil).GetNameAndSymbol), address, chainId)
}
