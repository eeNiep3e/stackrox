// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/interfaces.go -source interfaces.go
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	central "github.com/stackrox/rox/generated/internalapi/central"
	gomock "go.uber.org/mock/gomock"
	metadata "google.golang.org/grpc/metadata"
)

// MockServiceCommunicateClient is a mock of ServiceCommunicateClient interface.
type MockServiceCommunicateClient struct {
	ctrl     *gomock.Controller
	recorder *MockServiceCommunicateClientMockRecorder
}

// MockServiceCommunicateClientMockRecorder is the mock recorder for MockServiceCommunicateClient.
type MockServiceCommunicateClientMockRecorder struct {
	mock *MockServiceCommunicateClient
}

// NewMockServiceCommunicateClient creates a new mock instance.
func NewMockServiceCommunicateClient(ctrl *gomock.Controller) *MockServiceCommunicateClient {
	mock := &MockServiceCommunicateClient{ctrl: ctrl}
	mock.recorder = &MockServiceCommunicateClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceCommunicateClient) EXPECT() *MockServiceCommunicateClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockServiceCommunicateClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockServiceCommunicateClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockServiceCommunicateClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockServiceCommunicateClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockServiceCommunicateClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockServiceCommunicateClient)(nil).Context))
}

// Header mocks base method.
func (m *MockServiceCommunicateClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockServiceCommunicateClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockServiceCommunicateClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockServiceCommunicateClient) Recv() (*central.MsgToSensor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*central.MsgToSensor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockServiceCommunicateClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockServiceCommunicateClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockServiceCommunicateClient) RecvMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockServiceCommunicateClientMockRecorder) RecvMsg(m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockServiceCommunicateClient)(nil).RecvMsg), m)
}

// Send mocks base method.
func (m *MockServiceCommunicateClient) Send(arg0 *central.MsgFromSensor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockServiceCommunicateClientMockRecorder) Send(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockServiceCommunicateClient)(nil).Send), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockServiceCommunicateClient) SendMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockServiceCommunicateClientMockRecorder) SendMsg(m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockServiceCommunicateClient)(nil).SendMsg), m)
}

// Trailer mocks base method.
func (m *MockServiceCommunicateClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockServiceCommunicateClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockServiceCommunicateClient)(nil).Trailer))
}