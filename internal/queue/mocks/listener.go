// Code generated by MockGen. DO NOT EDIT.
// Source: listener.go

// Package mock_queue is a generated GoMock package.
package mock_queue

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	config "github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/config"
	queue "github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/queue"
	handler "github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/queue/handler"
)

// MockListener is a mock of Listener interface.
type MockListener struct {
	ctrl     *gomock.Controller
	recorder *MockListenerMockRecorder
}

// MockListenerMockRecorder is the mock recorder for MockListener.
type MockListenerMockRecorder struct {
	mock *MockListener
}

// NewMockListener creates a new mock instance.
func NewMockListener(ctrl *gomock.Controller) *MockListener {
	mock := &MockListener{ctrl: ctrl}
	mock.recorder = &MockListenerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockListener) EXPECT() *MockListenerMockRecorder {
	return m.recorder
}

// Start mocks base method.
func (m *MockListener) Start(consumer queue.Consumer, config *config.QueueConfig, handler handler.Handler) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", consumer, config, handler)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockListenerMockRecorder) Start(consumer, config, handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockListener)(nil).Start), consumer, config, handler)
}
