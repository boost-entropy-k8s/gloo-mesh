// Code generated by MockGen. DO NOT EDIT.
// Source: ./smi_destination_translator.go

// Package mock_destination is a generated GoMock package.
package mock_destination

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1"
	input "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/input"
	smi "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/output/smi"
	reporting "github.com/solo-io/gloo-mesh/pkg/mesh-networking/reporting"
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
func (m *MockTranslator) Translate(ctx context.Context, in input.LocalSnapshot, destination *v1.Destination, outputs smi.Builder, reporter reporting.Reporter) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Translate", ctx, in, destination, outputs, reporter)
}

// Translate indicates an expected call of Translate
func (mr *MockTranslatorMockRecorder) Translate(ctx, in, destination, outputs, reporter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Translate", reflect.TypeOf((*MockTranslator)(nil).Translate), ctx, in, destination, outputs, reporter)
}
