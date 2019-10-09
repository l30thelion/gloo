// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/gloo/projects/gloo/pkg/api/grpc/validation (interfaces: ProxyValidationServiceClient)

// Package mock_validation is a generated GoMock package.
package mock_validation

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	validation "github.com/solo-io/gloo/projects/gloo/pkg/api/grpc/validation"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockProxyValidationServiceClient is a mock of ProxyValidationServiceClient interface
type MockProxyValidationServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockProxyValidationServiceClientMockRecorder
}

// MockProxyValidationServiceClientMockRecorder is the mock recorder for MockProxyValidationServiceClient
type MockProxyValidationServiceClientMockRecorder struct {
	mock *MockProxyValidationServiceClient
}

// NewMockProxyValidationServiceClient creates a new mock instance
func NewMockProxyValidationServiceClient(ctrl *gomock.Controller) *MockProxyValidationServiceClient {
	mock := &MockProxyValidationServiceClient{ctrl: ctrl}
	mock.recorder = &MockProxyValidationServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProxyValidationServiceClient) EXPECT() *MockProxyValidationServiceClientMockRecorder {
	return m.recorder
}

// ValidateProxy mocks base method
func (m *MockProxyValidationServiceClient) ValidateProxy(arg0 context.Context, arg1 *validation.ProxyValidationServiceRequest, arg2 ...grpc.CallOption) (*validation.ProxyValidationServiceResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ValidateProxy", varargs...)
	ret0, _ := ret[0].(*validation.ProxyValidationServiceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateProxy indicates an expected call of ValidateProxy
func (mr *MockProxyValidationServiceClientMockRecorder) ValidateProxy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateProxy", reflect.TypeOf((*MockProxyValidationServiceClient)(nil).ValidateProxy), varargs...)
}