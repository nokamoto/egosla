// Code generated by MockGen. DO NOT EDIT.
// Source: name.go

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MocknameGenerator is a mock of nameGenerator interface.
type MocknameGenerator struct {
	ctrl     *gomock.Controller
	recorder *MocknameGeneratorMockRecorder
}

// MocknameGeneratorMockRecorder is the mock recorder for MocknameGenerator.
type MocknameGeneratorMockRecorder struct {
	mock *MocknameGenerator
}

// NewMocknameGenerator creates a new mock instance.
func NewMocknameGenerator(ctrl *gomock.Controller) *MocknameGenerator {
	mock := &MocknameGenerator{ctrl: ctrl}
	mock.recorder = &MocknameGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocknameGenerator) EXPECT() *MocknameGeneratorMockRecorder {
	return m.recorder
}

// newName mocks base method.
func (m *MocknameGenerator) newName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "newName")
	ret0, _ := ret[0].(string)
	return ret0
}

// newName indicates an expected call of newName.
func (mr *MocknameGeneratorMockRecorder) newName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "newName", reflect.TypeOf((*MocknameGenerator)(nil).newName))
}
