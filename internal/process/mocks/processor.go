// Code generated by MockGen. DO NOT EDIT.
// Source: processor.go

// Package process is a generated GoMock package.
package process

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/rog-golang-buddies/internal/model"
)

// MockUrlProcessor is a mock of UrlProcessor interface.
type MockUrlProcessor struct {
	ctrl     *gomock.Controller
	recorder *MockUrlProcessorMockRecorder
}

// MockUrlProcessorMockRecorder is the mock recorder for MockUrlProcessor.
type MockUrlProcessorMockRecorder struct {
	mock *MockUrlProcessor
}

// NewMockUrlProcessor creates a new mock instance.
func NewMockUrlProcessor(ctrl *gomock.Controller) *MockUrlProcessor {
	mock := &MockUrlProcessor{ctrl: ctrl}
	mock.recorder = &MockUrlProcessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUrlProcessor) EXPECT() *MockUrlProcessorMockRecorder {
	return m.recorder
}

// process mocks base method.
func (m *MockUrlProcessor) process(ctx context.Context, url string) (*model.ApiSpecDoc, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "process", ctx, url)
	ret0, _ := ret[0].(*model.ApiSpecDoc)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// process indicates an expected call of process.
func (mr *MockUrlProcessorMockRecorder) process(ctx, url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "process", reflect.TypeOf((*MockUrlProcessor)(nil).process), ctx, url)
}
