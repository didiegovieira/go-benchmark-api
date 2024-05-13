// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/application/use_case/sorting_algorithm/insertion_sort_interface.go
//
// Generated by this command:
//
//	mockgen -source=./internal/application/use_case/sorting_algorithm/insertion_sort_interface.go -destination=tests/mocks/sorting_algorithm_insertion_sort_mock.go -package mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockInsertionSortUseCaseInterface is a mock of InsertionSortUseCaseInterface interface.
type MockInsertionSortUseCaseInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInsertionSortUseCaseInterfaceMockRecorder
}

// MockInsertionSortUseCaseInterfaceMockRecorder is the mock recorder for MockInsertionSortUseCaseInterface.
type MockInsertionSortUseCaseInterfaceMockRecorder struct {
	mock *MockInsertionSortUseCaseInterface
}

// NewMockInsertionSortUseCaseInterface creates a new mock instance.
func NewMockInsertionSortUseCaseInterface(ctrl *gomock.Controller) *MockInsertionSortUseCaseInterface {
	mock := &MockInsertionSortUseCaseInterface{ctrl: ctrl}
	mock.recorder = &MockInsertionSortUseCaseInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInsertionSortUseCaseInterface) EXPECT() *MockInsertionSortUseCaseInterfaceMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockInsertionSortUseCaseInterface) Execute(arr []int) []int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", arr)
	ret0, _ := ret[0].([]int)
	return ret0
}

// Execute indicates an expected call of Execute.
func (mr *MockInsertionSortUseCaseInterfaceMockRecorder) Execute(arr any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockInsertionSortUseCaseInterface)(nil).Execute), arr)
}
