// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package client is a generated GoMock package.
package client

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockC is a mock of C interface.
type MockC struct {
	ctrl     *gomock.Controller
	recorder *MockCMockRecorder
}

// MockCMockRecorder is the mock recorder for MockC.
type MockCMockRecorder struct {
	mock *MockC
}

// NewMockC creates a new mock instance.
func NewMockC(ctrl *gomock.Controller) *MockC {
	mock := &MockC{ctrl: ctrl}
	mock.recorder = &MockCMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockC) EXPECT() *MockCMockRecorder {
	return m.recorder
}

// Connect mocks base method.
func (m *MockC) Connect() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect")
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect.
func (mr *MockCMockRecorder) Connect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockC)(nil).Connect))
}

// Get mocks base method.
func (m *MockC) Get(endpoint string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", endpoint)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCMockRecorder) Get(endpoint interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockC)(nil).Get), endpoint)
}
