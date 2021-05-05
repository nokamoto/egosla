// Code generated by MockGen. DO NOT EDIT.
// Source: persistent.go

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/nokamoto/egosla/api"
)

// Mockpersistent is a mock of persistent interface.
type Mockpersistent struct {
	ctrl     *gomock.Controller
	recorder *MockpersistentMockRecorder
}

// MockpersistentMockRecorder is the mock recorder for Mockpersistent.
type MockpersistentMockRecorder struct {
	mock *Mockpersistent
}

// NewMockpersistent creates a new mock instance.
func NewMockpersistent(ctrl *gomock.Controller) *Mockpersistent {
	mock := &Mockpersistent{ctrl: ctrl}
	mock.recorder = &MockpersistentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockpersistent) EXPECT() *MockpersistentMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *Mockpersistent) Create(arg0 *api.Watcher) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockpersistentMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*Mockpersistent)(nil).Create), arg0)
}