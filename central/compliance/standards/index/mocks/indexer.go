// Code generated by MockGen. DO NOT EDIT.
// Source: indexer.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/stackrox/rox/generated/api/v1"
	search "github.com/stackrox/rox/pkg/search"
	reflect "reflect"
)

// MockIndexer is a mock of Indexer interface
type MockIndexer struct {
	ctrl     *gomock.Controller
	recorder *MockIndexerMockRecorder
}

// MockIndexerMockRecorder is the mock recorder for MockIndexer
type MockIndexerMockRecorder struct {
	mock *MockIndexer
}

// NewMockIndexer creates a new mock instance
func NewMockIndexer(ctrl *gomock.Controller) *MockIndexer {
	mock := &MockIndexer{ctrl: ctrl}
	mock.recorder = &MockIndexerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIndexer) EXPECT() *MockIndexerMockRecorder {
	return m.recorder
}

// IndexStandard mocks base method
func (m *MockIndexer) IndexStandard(standard *v1.ComplianceStandard) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IndexStandard", standard)
	ret0, _ := ret[0].(error)
	return ret0
}

// IndexStandard indicates an expected call of IndexStandard
func (mr *MockIndexerMockRecorder) IndexStandard(standard interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IndexStandard", reflect.TypeOf((*MockIndexer)(nil).IndexStandard), standard)
}

// DeleteStandard mocks base method
func (m *MockIndexer) DeleteStandard(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStandard", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStandard indicates an expected call of DeleteStandard
func (mr *MockIndexerMockRecorder) DeleteStandard(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStandard", reflect.TypeOf((*MockIndexer)(nil).DeleteStandard), id)
}

// DeleteControl mocks base method
func (m *MockIndexer) DeleteControl(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteControl", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteControl indicates an expected call of DeleteControl
func (mr *MockIndexerMockRecorder) DeleteControl(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteControl", reflect.TypeOf((*MockIndexer)(nil).DeleteControl), id)
}

// SearchStandards mocks base method
func (m *MockIndexer) SearchStandards(q *v1.Query) ([]search.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchStandards", q)
	ret0, _ := ret[0].([]search.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchStandards indicates an expected call of SearchStandards
func (mr *MockIndexerMockRecorder) SearchStandards(q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchStandards", reflect.TypeOf((*MockIndexer)(nil).SearchStandards), q)
}

// SearchControls mocks base method
func (m *MockIndexer) SearchControls(q *v1.Query) ([]search.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchControls", q)
	ret0, _ := ret[0].([]search.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchControls indicates an expected call of SearchControls
func (mr *MockIndexerMockRecorder) SearchControls(q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchControls", reflect.TypeOf((*MockIndexer)(nil).SearchControls), q)
}
