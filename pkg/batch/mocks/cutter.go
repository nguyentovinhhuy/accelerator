/*
 *    Copyright 2019 Samsung SDS
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

// Code generated by MockGen. DO NOT EDIT.
// Source: ./cutter.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	cutter "github.com/nexledger/accelerator/pkg/batch/queue/cutter"
	tx "github.com/nexledger/accelerator/pkg/batch/tx"
	reflect "reflect"
)

// MockCutter is a mock of Cutter interface
type MockCutter struct {
	ctrl     *gomock.Controller
	recorder *MockCutterMockRecorder
}

// MockCutterMockRecorder is the mock recorder for MockCutter
type MockCutterMockRecorder struct {
	mock *MockCutter
}

// NewMockCutter creates a new mock instance
func NewMockCutter(ctrl *gomock.Controller) *MockCutter {
	mock := &MockCutter{ctrl: ctrl}
	mock.recorder = &MockCutterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCutter) EXPECT() *MockCutterMockRecorder {
	return m.recorder
}

// Before mocks base method
func (m *MockCutter) Before(job *tx.Job, next *tx.Item) (cutter.Cut, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Before", job, next)
	ret0, _ := ret[0].(cutter.Cut)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Before indicates an expected call of Before
func (mr *MockCutterMockRecorder) Before(job, next interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Before", reflect.TypeOf((*MockCutter)(nil).Before), job, next)
}

// After mocks base method
func (m *MockCutter) After(job *tx.Job) cutter.Cut {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "After", job)
	ret0, _ := ret[0].(cutter.Cut)
	return ret0
}

// After indicates an expected call of After
func (mr *MockCutterMockRecorder) After(job interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "After", reflect.TypeOf((*MockCutter)(nil).After), job)
}

// Clear mocks base method
func (m *MockCutter) Clear() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Clear")
}

// Clear indicates an expected call of Clear
func (mr *MockCutterMockRecorder) Clear() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clear", reflect.TypeOf((*MockCutter)(nil).Clear))
}
