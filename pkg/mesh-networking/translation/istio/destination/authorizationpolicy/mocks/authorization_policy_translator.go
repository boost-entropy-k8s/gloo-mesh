// Code generated by MockGen. DO NOT EDIT.
// Source: ./authorization_policy_translator.go

// Package mock_authorizationpolicy is a generated GoMock package.
package mock_authorizationpolicy

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1"
	input "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/input"
	reporting "github.com/solo-io/gloo-mesh/pkg/mesh-networking/reporting"
	v1beta1 "istio.io/client-go/pkg/apis/security/v1beta1"
)

// MockTranslator is a mock of Translator interface
type MockTranslator struct {
	ctrl     *gomock.Controller
	recorder *MockTranslatorMockRecorder
}

// MockTranslatorMockRecorder is the mock recorder for MockTranslator
type MockTranslatorMockRecorder struct {
	mock *MockTranslator
}

// NewMockTranslator creates a new mock instance
func NewMockTranslator(ctrl *gomock.Controller) *MockTranslator {
	mock := &MockTranslator{ctrl: ctrl}
	mock.recorder = &MockTranslatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTranslator) EXPECT() *MockTranslatorMockRecorder {
	return m.recorder
}

// Translate mocks base method
func (m *MockTranslator) Translate(in input.LocalSnapshot, destination *v1.Destination, reporter reporting.Reporter) *v1beta1.AuthorizationPolicy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Translate", in, destination, reporter)
	ret0, _ := ret[0].(*v1beta1.AuthorizationPolicy)
	return ret0
}

// Translate indicates an expected call of Translate
func (mr *MockTranslatorMockRecorder) Translate(in, destination, reporter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Translate", reflect.TypeOf((*MockTranslator)(nil).Translate), in, destination, reporter)
}
