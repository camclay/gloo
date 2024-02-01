// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/install (interfaces: HelmClient)
//
// Generated by this command:
//
//	mockgen -destination mocks/mock_helm_client.go -package mocks github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/install HelmClient
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	install "github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/install"
	gomock "go.uber.org/mock/gomock"
	chart "helm.sh/helm/v3/pkg/chart"
	cli "helm.sh/helm/v3/pkg/cli"
)

// MockHelmClient is a mock of HelmClient interface.
type MockHelmClient struct {
	ctrl     *gomock.Controller
	recorder *MockHelmClientMockRecorder
}

// MockHelmClientMockRecorder is the mock recorder for MockHelmClient.
type MockHelmClientMockRecorder struct {
	mock *MockHelmClient
}

// NewMockHelmClient creates a new mock instance.
func NewMockHelmClient(ctrl *gomock.Controller) *MockHelmClient {
	mock := &MockHelmClient{ctrl: ctrl}
	mock.recorder = &MockHelmClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHelmClient) EXPECT() *MockHelmClientMockRecorder {
	return m.recorder
}

// DownloadChart mocks base method.
func (m *MockHelmClient) DownloadChart(arg0 string) (*chart.Chart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadChart", arg0)
	ret0, _ := ret[0].(*chart.Chart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownloadChart indicates an expected call of DownloadChart.
func (mr *MockHelmClientMockRecorder) DownloadChart(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadChart", reflect.TypeOf((*MockHelmClient)(nil).DownloadChart), arg0)
}

// NewInstall mocks base method.
func (m *MockHelmClient) NewInstall(arg0, arg1 string, arg2 bool, arg3 string) (install.HelmInstallation, *cli.EnvSettings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewInstall", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(install.HelmInstallation)
	ret1, _ := ret[1].(*cli.EnvSettings)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// NewInstall indicates an expected call of NewInstall.
func (mr *MockHelmClientMockRecorder) NewInstall(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewInstall", reflect.TypeOf((*MockHelmClient)(nil).NewInstall), arg0, arg1, arg2, arg3)
}

// NewUninstall mocks base method.
func (m *MockHelmClient) NewUninstall(arg0 string) (install.HelmUninstallation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewUninstall", arg0)
	ret0, _ := ret[0].(install.HelmUninstallation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewUninstall indicates an expected call of NewUninstall.
func (mr *MockHelmClientMockRecorder) NewUninstall(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUninstall", reflect.TypeOf((*MockHelmClient)(nil).NewUninstall), arg0)
}

// ReleaseExists mocks base method.
func (m *MockHelmClient) ReleaseExists(arg0, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReleaseExists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReleaseExists indicates an expected call of ReleaseExists.
func (mr *MockHelmClientMockRecorder) ReleaseExists(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseExists", reflect.TypeOf((*MockHelmClient)(nil).ReleaseExists), arg0, arg1)
}

// ReleaseList mocks base method.
func (m *MockHelmClient) ReleaseList(arg0 string) (install.HelmReleaseListRunner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReleaseList", arg0)
	ret0, _ := ret[0].(install.HelmReleaseListRunner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReleaseList indicates an expected call of ReleaseList.
func (mr *MockHelmClientMockRecorder) ReleaseList(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseList", reflect.TypeOf((*MockHelmClient)(nil).ReleaseList), arg0)
}
