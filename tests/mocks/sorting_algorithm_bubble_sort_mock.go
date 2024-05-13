// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/application/use_case/sorting_algorithm/bubble_sort_interface.go
//
// Generated by this command:
//
//	mockgen -source=./internal/application/use_case/sorting_algorithm/bubble_sort_interface.go -destination=tests/mocks/sorting_algorithm_bubble_sort_mock.go -package mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockBubbleSortUseCaseInterface is a mock of BubbleSortUseCaseInterface interface.
type MockBubbleSortUseCaseInterface struct {
	ctrl     *gomock.Controller
	recorder *MockBubbleSortUseCaseInterfaceMockRecorder
}

// MockBubbleSortUseCaseInterfaceMockRecorder is the mock recorder for MockBubbleSortUseCaseInterface.
type MockBubbleSortUseCaseInterfaceMockRecorder struct {
	mock *MockBubbleSortUseCaseInterface
}

// NewMockBubbleSortUseCaseInterface creates a new mock instance.
func NewMockBubbleSortUseCaseInterface(ctrl *gomock.Controller) *MockBubbleSortUseCaseInterface {
	mock := &MockBubbleSortUseCaseInterface{ctrl: ctrl}
	mock.recorder = &MockBubbleSortUseCaseInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBubbleSortUseCaseInterface) EXPECT() *MockBubbleSortUseCaseInterfaceMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockBubbleSortUseCaseInterface) Execute(arr []int) []int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", arr)
	ret0, _ := ret[0].([]int)
	return ret0
}

// Execute indicates an expected call of Execute.
func (mr *MockBubbleSortUseCaseInterfaceMockRecorder) Execute(arr any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockBubbleSortUseCaseInterface)(nil).Execute), arr)
}