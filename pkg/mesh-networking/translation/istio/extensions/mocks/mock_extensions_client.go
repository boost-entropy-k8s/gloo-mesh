// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/extensions/v1beta1 (interfaces: NetworkingExtensionsClient,NetworkingExtensions_WatchPushNotificationsClient)

// Package mock_extensions is a generated GoMock package.
package mock_extensions

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1 "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/extensions/v1beta1"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockNetworkingExtensionsClient is a mock of NetworkingExtensionsClient interface.
type MockNetworkingExtensionsClient struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkingExtensionsClientMockRecorder
}

// MockNetworkingExtensionsClientMockRecorder is the mock recorder for MockNetworkingExtensionsClient.
type MockNetworkingExtensionsClientMockRecorder struct {
	mock *MockNetworkingExtensionsClient
}

// NewMockNetworkingExtensionsClient creates a new mock instance.
func NewMockNetworkingExtensionsClient(ctrl *gomock.Controller) *MockNetworkingExtensionsClient {
	mock := &MockNetworkingExtensionsClient{ctrl: ctrl}
	mock.recorder = &MockNetworkingExtensionsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworkingExtensionsClient) EXPECT() *MockNetworkingExtensionsClientMockRecorder {
	return m.recorder
}

// GetExtensionPatches mocks base method.
func (m *MockNetworkingExtensionsClient) GetExtensionPatches(arg0 context.Context, arg1 *v1beta1.ExtensionPatchRequest, arg2 ...grpc.CallOption) (*v1beta1.ExtensionPatchResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetExtensionPatches", varargs...)
	ret0, _ := ret[0].(*v1beta1.ExtensionPatchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExtensionPatches indicates an expected call of GetExtensionPatches.
func (mr *MockNetworkingExtensionsClientMockRecorder) GetExtensionPatches(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExtensionPatches", reflect.TypeOf((*MockNetworkingExtensionsClient)(nil).GetExtensionPatches), varargs...)
}

// WatchPushNotifications mocks base method.
func (m *MockNetworkingExtensionsClient) WatchPushNotifications(arg0 context.Context, arg1 *v1beta1.WatchPushNotificationsRequest, arg2 ...grpc.CallOption) (v1beta1.NetworkingExtensions_WatchPushNotificationsClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WatchPushNotifications", varargs...)
	ret0, _ := ret[0].(v1beta1.NetworkingExtensions_WatchPushNotificationsClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchPushNotifications indicates an expected call of WatchPushNotifications.
func (mr *MockNetworkingExtensionsClientMockRecorder) WatchPushNotifications(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchPushNotifications", reflect.TypeOf((*MockNetworkingExtensionsClient)(nil).WatchPushNotifications), varargs...)
}

// MockNetworkingExtensions_WatchPushNotificationsClient is a mock of NetworkingExtensions_WatchPushNotificationsClient interface.
type MockNetworkingExtensions_WatchPushNotificationsClient struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkingExtensions_WatchPushNotificationsClientMockRecorder
}

// MockNetworkingExtensions_WatchPushNotificationsClientMockRecorder is the mock recorder for MockNetworkingExtensions_WatchPushNotificationsClient.
type MockNetworkingExtensions_WatchPushNotificationsClientMockRecorder struct {
	mock *MockNetworkingExtensions_WatchPushNotificationsClient
}

// NewMockNetworkingExtensions_WatchPushNotificationsClient creates a new mock instance.
func NewMockNetworkingExtensions_WatchPushNotificationsClient(ctrl *gomock.Controller) *MockNetworkingExtensions_WatchPushNotificationsClient {
	mock := &MockNetworkingExtensions_WatchPushNotificationsClient{ctrl: ctrl}
	mock.recorder = &MockNetworkingExtensions_WatchPushNotificationsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworkingExtensions_WatchPushNotificationsClient) EXPECT() *MockNetworkingExtensions_WatchPushNotificationsClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockNetworkingExtensions_WatchPushNotificationsClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockNetworkingExtensions_WatchPushNotificationsClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockNetworkingExtensions_WatchPushNotificationsClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockNetworkingExtensions_WatchPushNotificationsClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockNetworkingExtensions_WatchPushNotificationsClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockNetworkingExtensions_WatchPushNotificationsClient)(nil).Context))
}

// Header mocks base method.
func (m *MockNetworkingExtensions_WatchPushNotificationsClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockNetworkingExtensions_WatchPushNotificationsClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockNetworkingExtensions_WatchPushNotificationsClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockNetworkingExtensions_WatchPushNotificationsClient) Recv() (*v1beta1.PushNotification, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*v1beta1.PushNotification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockNetworkingExtensions_WatchPushNotificationsClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockNetworkingExtensions_WatchPushNotificationsClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m *MockNetworkingExtensions_WatchPushNotificationsClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockNetworkingExtensions_WatchPushNotificationsClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockNetworkingExtensions_WatchPushNotificationsClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method.
func (m *MockNetworkingExtensions_WatchPushNotificationsClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockNetworkingExtensions_WatchPushNotificationsClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockNetworkingExtensions_WatchPushNotificationsClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method.
func (m *MockNetworkingExtensions_WatchPushNotificationsClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockNetworkingExtensions_WatchPushNotificationsClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockNetworkingExtensions_WatchPushNotificationsClient)(nil).Trailer))
}
