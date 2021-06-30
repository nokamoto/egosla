// Code generated by MockGen. DO NOT EDIT.
// Source: persistent.go

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/nokamoto/egosla/api"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	proto "google.golang.org/protobuf/proto"
)

// MockpersistentWatcher is a mock of persistentWatcher interface.
type MockpersistentWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockpersistentWatcherMockRecorder
}

// MockpersistentWatcherMockRecorder is the mock recorder for MockpersistentWatcher.
type MockpersistentWatcherMockRecorder struct {
	mock *MockpersistentWatcher
}

// NewMockpersistentWatcher creates a new mock instance.
func NewMockpersistentWatcher(ctrl *gomock.Controller) *MockpersistentWatcher {
	mock := &MockpersistentWatcher{ctrl: ctrl}
	mock.recorder = &MockpersistentWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockpersistentWatcher) EXPECT() *MockpersistentWatcherMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockpersistentWatcher) Create(arg0 *api.Watcher) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockpersistentWatcherMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockpersistentWatcher)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockpersistentWatcher) Delete(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockpersistentWatcherMockRecorder) Delete(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockpersistentWatcher)(nil).Delete), name)
}

// Get mocks base method.
func (m *MockpersistentWatcher) Get(name string) (*api.Watcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", name)
	ret0, _ := ret[0].(*api.Watcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockpersistentWatcherMockRecorder) Get(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockpersistentWatcher)(nil).Get), name)
}

// List mocks base method.
func (m *MockpersistentWatcher) List(offset, limit int) ([]*api.Watcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", offset, limit)
	ret0, _ := ret[0].([]*api.Watcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockpersistentWatcherMockRecorder) List(offset, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockpersistentWatcher)(nil).List), offset, limit)
}

// Update mocks base method.
func (m *MockpersistentWatcher) Update(arg0 *api.Watcher, arg1 *field_mask.FieldMask) (*api.Watcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*api.Watcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockpersistentWatcherMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockpersistentWatcher)(nil).Update), arg0, arg1)
}

// MockpersistentSubscription is a mock of persistentSubscription interface.
type MockpersistentSubscription struct {
	ctrl     *gomock.Controller
	recorder *MockpersistentSubscriptionMockRecorder
}

// MockpersistentSubscriptionMockRecorder is the mock recorder for MockpersistentSubscription.
type MockpersistentSubscriptionMockRecorder struct {
	mock *MockpersistentSubscription
}

// NewMockpersistentSubscription creates a new mock instance.
func NewMockpersistentSubscription(ctrl *gomock.Controller) *MockpersistentSubscription {
	mock := &MockpersistentSubscription{ctrl: ctrl}
	mock.recorder = &MockpersistentSubscriptionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockpersistentSubscription) EXPECT() *MockpersistentSubscriptionMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockpersistentSubscription) Create(arg0 *api.Subscription) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockpersistentSubscriptionMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockpersistentSubscription)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockpersistentSubscription) Delete(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockpersistentSubscriptionMockRecorder) Delete(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockpersistentSubscription)(nil).Delete), name)
}

// Get mocks base method.
func (m *MockpersistentSubscription) Get(name string) (*api.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", name)
	ret0, _ := ret[0].(*api.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockpersistentSubscriptionMockRecorder) Get(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockpersistentSubscription)(nil).Get), name)
}

// List mocks base method.
func (m *MockpersistentSubscription) List(offset, limit int) ([]*api.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", offset, limit)
	ret0, _ := ret[0].([]*api.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockpersistentSubscriptionMockRecorder) List(offset, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockpersistentSubscription)(nil).List), offset, limit)
}

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
func (m *Mockpersistent) Create(arg0 proto.Message) error {
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
