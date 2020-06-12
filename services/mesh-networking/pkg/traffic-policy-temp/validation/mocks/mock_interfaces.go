// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_traffic_policy_validation is a generated GoMock package.
package mock_traffic_policy_validation

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/solo-io/service-mesh-hub/pkg/api/core.smh.solo.io/v1alpha1/types"
	v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1"
	v1alpha10 "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha1"
)

// MockValidator is a mock of Validator interface.
type MockValidator struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorMockRecorder
}

// MockValidatorMockRecorder is the mock recorder for MockValidator.
type MockValidatorMockRecorder struct {
	mock *MockValidator
}

// NewMockValidator creates a new mock instance.
func NewMockValidator(ctrl *gomock.Controller) *MockValidator {
	mock := &MockValidator{ctrl: ctrl}
	mock.recorder = &MockValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidator) EXPECT() *MockValidatorMockRecorder {
	return m.recorder
}

// ValidateTrafficPolicy mocks base method.
func (m *MockValidator) ValidateTrafficPolicy(trafficPolicy *v1alpha10.TrafficPolicy, allMeshServices []*v1alpha1.MeshService) (*types.Status, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateTrafficPolicy", trafficPolicy, allMeshServices)
	ret0, _ := ret[0].(*types.Status)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateTrafficPolicy indicates an expected call of ValidateTrafficPolicy.
func (mr *MockValidatorMockRecorder) ValidateTrafficPolicy(trafficPolicy, allMeshServices interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateTrafficPolicy", reflect.TypeOf((*MockValidator)(nil).ValidateTrafficPolicy), trafficPolicy, allMeshServices)
}
