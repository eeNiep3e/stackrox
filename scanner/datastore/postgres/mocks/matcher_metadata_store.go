// Code generated by MockGen. DO NOT EDIT.
// Source: matcher_metadata_store.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/matcher_metadata_store.go -source matcher_metadata_store.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"
)

// MockMatcherMetadataStore is a mock of MatcherMetadataStore interface.
type MockMatcherMetadataStore struct {
	ctrl     *gomock.Controller
	recorder *MockMatcherMetadataStoreMockRecorder
	isgomock struct{}
}

// MockMatcherMetadataStoreMockRecorder is the mock recorder for MockMatcherMetadataStore.
type MockMatcherMetadataStoreMockRecorder struct {
	mock *MockMatcherMetadataStore
}

// NewMockMatcherMetadataStore creates a new mock instance.
func NewMockMatcherMetadataStore(ctrl *gomock.Controller) *MockMatcherMetadataStore {
	mock := &MockMatcherMetadataStore{ctrl: ctrl}
	mock.recorder = &MockMatcherMetadataStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMatcherMetadataStore) EXPECT() *MockMatcherMetadataStoreMockRecorder {
	return m.recorder
}

// GCVulnerabilityUpdates mocks base method.
func (m *MockMatcherMetadataStore) GCVulnerabilityUpdates(ctx context.Context, activeUpdaters []string, lastUpdate time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GCVulnerabilityUpdates", ctx, activeUpdaters, lastUpdate)
	ret0, _ := ret[0].(error)
	return ret0
}

// GCVulnerabilityUpdates indicates an expected call of GCVulnerabilityUpdates.
func (mr *MockMatcherMetadataStoreMockRecorder) GCVulnerabilityUpdates(ctx, activeUpdaters, lastUpdate any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GCVulnerabilityUpdates", reflect.TypeOf((*MockMatcherMetadataStore)(nil).GCVulnerabilityUpdates), ctx, activeUpdaters, lastUpdate)
}

// GetLastVulnerabilityUpdate mocks base method.
func (m *MockMatcherMetadataStore) GetLastVulnerabilityUpdate(ctx context.Context) (time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastVulnerabilityUpdate", ctx)
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLastVulnerabilityUpdate indicates an expected call of GetLastVulnerabilityUpdate.
func (mr *MockMatcherMetadataStoreMockRecorder) GetLastVulnerabilityUpdate(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastVulnerabilityUpdate", reflect.TypeOf((*MockMatcherMetadataStore)(nil).GetLastVulnerabilityUpdate), ctx)
}

// GetOrSetLastVulnerabilityUpdate mocks base method.
func (m *MockMatcherMetadataStore) GetOrSetLastVulnerabilityUpdate(ctx context.Context, bundle string, lastUpdate time.Time) (time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrSetLastVulnerabilityUpdate", ctx, bundle, lastUpdate)
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrSetLastVulnerabilityUpdate indicates an expected call of GetOrSetLastVulnerabilityUpdate.
func (mr *MockMatcherMetadataStoreMockRecorder) GetOrSetLastVulnerabilityUpdate(ctx, bundle, lastUpdate any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrSetLastVulnerabilityUpdate", reflect.TypeOf((*MockMatcherMetadataStore)(nil).GetOrSetLastVulnerabilityUpdate), ctx, bundle, lastUpdate)
}

// SetLastVulnerabilityUpdate mocks base method.
func (m *MockMatcherMetadataStore) SetLastVulnerabilityUpdate(ctx context.Context, bundle string, lastUpdate time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetLastVulnerabilityUpdate", ctx, bundle, lastUpdate)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetLastVulnerabilityUpdate indicates an expected call of SetLastVulnerabilityUpdate.
func (mr *MockMatcherMetadataStoreMockRecorder) SetLastVulnerabilityUpdate(ctx, bundle, lastUpdate any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLastVulnerabilityUpdate", reflect.TypeOf((*MockMatcherMetadataStore)(nil).SetLastVulnerabilityUpdate), ctx, bundle, lastUpdate)
}
