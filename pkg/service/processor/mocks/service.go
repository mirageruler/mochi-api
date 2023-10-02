// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/service/processor/service.go

// Package mock_processor is a generated GoMock package.
package mock_processor

import (
	model "github.com/defipod/mochi/pkg/model"
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

// CreateUserTransaction mocks base method
func (m *MockService) CreateUserTransaction(createUserTransactionRequest model.CreateUserTransaction) (*model.CreateUserTxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserTransaction", createUserTransactionRequest)
	ret0, _ := ret[0].(*model.CreateUserTxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUserTransaction indicates an expected call of CreateUserTransaction
func (mr *MockServiceMockRecorder) CreateUserTransaction(createUserTransactionRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserTransaction", reflect.TypeOf((*MockService)(nil).CreateUserTransaction), createUserTransactionRequest)
}
