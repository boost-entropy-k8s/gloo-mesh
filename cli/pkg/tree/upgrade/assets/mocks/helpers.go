// Code generated by MockGen. DO NOT EDIT.
// Source: ./helpers.go

// Package mock_upgrade_assets is a generated GoMock package.
package mock_upgrade_assets

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	github "github.com/google/go-github/github"
)

// MockGithubAssetClient is a mock of GithubAssetClient interface.
type MockGithubAssetClient struct {
	ctrl     *gomock.Controller
	recorder *MockGithubAssetClientMockRecorder
}

// MockGithubAssetClientMockRecorder is the mock recorder for MockGithubAssetClient.
type MockGithubAssetClientMockRecorder struct {
	mock *MockGithubAssetClient
}

// NewMockGithubAssetClient creates a new mock instance.
func NewMockGithubAssetClient(ctrl *gomock.Controller) *MockGithubAssetClient {
	mock := &MockGithubAssetClient{ctrl: ctrl}
	mock.recorder = &MockGithubAssetClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGithubAssetClient) EXPECT() *MockGithubAssetClientMockRecorder {
	return m.recorder
}

// ListReleases mocks base method.
func (m *MockGithubAssetClient) ListReleases(ctx context.Context, opt *github.ListOptions) ([]*github.RepositoryRelease, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListReleases", ctx, opt)
	ret0, _ := ret[0].([]*github.RepositoryRelease)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListReleases indicates an expected call of ListReleases.
func (mr *MockGithubAssetClientMockRecorder) ListReleases(ctx, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListReleases", reflect.TypeOf((*MockGithubAssetClient)(nil).ListReleases), ctx, opt)
}

// GetReleaseByTag mocks base method.
func (m *MockGithubAssetClient) GetReleaseByTag(ctx context.Context, tag string) (*github.RepositoryRelease, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReleaseByTag", ctx, tag)
	ret0, _ := ret[0].(*github.RepositoryRelease)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReleaseByTag indicates an expected call of GetReleaseByTag.
func (mr *MockGithubAssetClientMockRecorder) GetReleaseByTag(ctx, tag interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReleaseByTag", reflect.TypeOf((*MockGithubAssetClient)(nil).GetReleaseByTag), ctx, tag)
}

// MockAssetHelper is a mock of AssetHelper interface.
type MockAssetHelper struct {
	ctrl     *gomock.Controller
	recorder *MockAssetHelperMockRecorder
}

// MockAssetHelperMockRecorder is the mock recorder for MockAssetHelper.
type MockAssetHelperMockRecorder struct {
	mock *MockAssetHelper
}

// NewMockAssetHelper creates a new mock instance.
func NewMockAssetHelper(ctrl *gomock.Controller) *MockAssetHelper {
	mock := &MockAssetHelper{ctrl: ctrl}
	mock.recorder = &MockAssetHelperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAssetHelper) EXPECT() *MockAssetHelperMockRecorder {
	return m.recorder
}

// GetReleaseWithAsset mocks base method.
func (m *MockAssetHelper) GetReleaseWithAsset(ctx context.Context, tag, expectedAssetName string) (*github.RepositoryRelease, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReleaseWithAsset", ctx, tag, expectedAssetName)
	ret0, _ := ret[0].(*github.RepositoryRelease)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReleaseWithAsset indicates an expected call of GetReleaseWithAsset.
func (mr *MockAssetHelperMockRecorder) GetReleaseWithAsset(ctx, tag, expectedAssetName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReleaseWithAsset", reflect.TypeOf((*MockAssetHelper)(nil).GetReleaseWithAsset), ctx, tag, expectedAssetName)
}

// DownloadAsset mocks base method.
func (m *MockAssetHelper) DownloadAsset(downloadUrl, destFile string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadAsset", downloadUrl, destFile)
	ret0, _ := ret[0].(error)
	return ret0
}

// DownloadAsset indicates an expected call of DownloadAsset.
func (mr *MockAssetHelperMockRecorder) DownloadAsset(downloadUrl, destFile interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadAsset", reflect.TypeOf((*MockAssetHelper)(nil).DownloadAsset), downloadUrl, destFile)
}
