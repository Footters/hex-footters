// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/media/provider/auth/auth.go
// mockgen -source=./pkg/media/provider/auth/auth.go -destination=./pkg/media/mocks/authService.go -package=mocks
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockServiceProvider is a mock of ServiceProvider interface
type MockServiceProvider struct {
	ctrl     *gomock.Controller
	recorder *MockServiceProviderMockRecorder
}

// MockServiceProviderMockRecorder is the mock recorder for MockServiceProvider
type MockServiceProviderMockRecorder struct {
	mock *MockServiceProvider
}

// NewMockServiceProvider creates a new mock instance
func NewMockServiceProvider(ctrl *gomock.Controller) *MockServiceProvider {
	mock := &MockServiceProvider{ctrl: ctrl}
	mock.recorder = &MockServiceProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServiceProvider) EXPECT() *MockServiceProviderMockRecorder {
	return m.recorder
}

// Login mocks base method
func (m *MockServiceProvider) Login() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login")
	ret0, _ := ret[0].(string)
	return ret0
}

// Login indicates an expected call of Login
func (mr *MockServiceProviderMockRecorder) Login() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockServiceProvider)(nil).Login))
}
