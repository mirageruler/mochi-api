// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/service/marketplace/service.go

// Package mock_marketplace is a generated GoMock package.
package mock_marketplace

import (
	response "github.com/defipod/mochi/pkg/response"
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

// HandleMarketplaceLink mocks base method
func (m *MockService) HandleMarketplaceLink(contractAddress, chain string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleMarketplaceLink", contractAddress, chain)
	ret0, _ := ret[0].(string)
	return ret0
}

// HandleMarketplaceLink indicates an expected call of HandleMarketplaceLink
func (mr *MockServiceMockRecorder) HandleMarketplaceLink(contractAddress, chain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleMarketplaceLink", reflect.TypeOf((*MockService)(nil).HandleMarketplaceLink), contractAddress, chain)
}

// GetCollectionFromOpensea mocks base method
func (m *MockService) GetCollectionFromOpensea(collectionSymbol string) (*response.OpenseaGetCollectionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollectionFromOpensea", collectionSymbol)
	ret0, _ := ret[0].(*response.OpenseaGetCollectionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCollectionFromOpensea indicates an expected call of GetCollectionFromOpensea
func (mr *MockServiceMockRecorder) GetCollectionFromOpensea(collectionSymbol interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollectionFromOpensea", reflect.TypeOf((*MockService)(nil).GetCollectionFromOpensea), collectionSymbol)
}

// GetCollectionFromQuixotic mocks base method
func (m *MockService) GetCollectionFromQuixotic(collectionSymbol string) (*response.QuixoticCollectionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollectionFromQuixotic", collectionSymbol)
	ret0, _ := ret[0].(*response.QuixoticCollectionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCollectionFromQuixotic indicates an expected call of GetCollectionFromQuixotic
func (mr *MockServiceMockRecorder) GetCollectionFromQuixotic(collectionSymbol interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollectionFromQuixotic", reflect.TypeOf((*MockService)(nil).GetCollectionFromQuixotic), collectionSymbol)
}

// GetCollectionFromPaintswap mocks base method
func (m *MockService) GetCollectionFromPaintswap(address string) (*response.PaintswapCollectionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollectionFromPaintswap", address)
	ret0, _ := ret[0].(*response.PaintswapCollectionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCollectionFromPaintswap indicates an expected call of GetCollectionFromPaintswap
func (mr *MockServiceMockRecorder) GetCollectionFromPaintswap(address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollectionFromPaintswap", reflect.TypeOf((*MockService)(nil).GetCollectionFromPaintswap), address)
}

// GetOpenseaAssetContract mocks base method
func (m *MockService) GetOpenseaAssetContract(address string) (*response.OpenseaAssetContractResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOpenseaAssetContract", address)
	ret0, _ := ret[0].(*response.OpenseaAssetContractResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOpenseaAssetContract indicates an expected call of GetOpenseaAssetContract
func (mr *MockServiceMockRecorder) GetOpenseaAssetContract(address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOpenseaAssetContract", reflect.TypeOf((*MockService)(nil).GetOpenseaAssetContract), address)
}
