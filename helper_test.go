// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bendrucker/terraform-credential-helper-sdk (interfaces: Helper)

// Package credentialhelper is a generated GoMock package.
package credentialhelper

import (
	flag "flag"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockHelper is a mock of Helper interface
type MockHelper struct {
	ctrl     *gomock.Controller
	recorder *MockHelperMockRecorder
}

// MockHelperMockRecorder is the mock recorder for MockHelper
type MockHelperMockRecorder struct {
	mock *MockHelper
}

// NewMockHelper creates a new mock instance
func NewMockHelper(ctrl *gomock.Controller) *MockHelper {
	mock := &MockHelper{ctrl: ctrl}
	mock.recorder = &MockHelperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHelper) EXPECT() *MockHelperMockRecorder {
	return m.recorder
}

// Forget mocks base method
func (m *MockHelper) Forget(arg0 string, arg1 *flag.FlagSet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Forget", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Forget indicates an expected call of Forget
func (mr *MockHelperMockRecorder) Forget(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Forget", reflect.TypeOf((*MockHelper)(nil).Forget), arg0, arg1)
}

// Get mocks base method
func (m *MockHelper) Get(arg0 string, arg1 *flag.FlagSet) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockHelperMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockHelper)(nil).Get), arg0, arg1)
}

// Store mocks base method
func (m *MockHelper) Store(arg0 string, arg1 []byte, arg2 *flag.FlagSet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store
func (mr *MockHelperMockRecorder) Store(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockHelper)(nil).Store), arg0, arg1, arg2)
}
