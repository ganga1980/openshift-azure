// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/openshift-azure/pkg/util/azureclient/managedapplications (interfaces: ApplicationsClient)

// Package mock_managedapplications is a generated GoMock package.
package mock_managedapplications

import (
	context "context"
	reflect "reflect"

	managedapplications "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-06-01/managedapplications"
	autorest "github.com/Azure/go-autorest/autorest"
	gomock "github.com/golang/mock/gomock"
)

// MockApplicationsClient is a mock of ApplicationsClient interface
type MockApplicationsClient struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationsClientMockRecorder
}

// MockApplicationsClientMockRecorder is the mock recorder for MockApplicationsClient
type MockApplicationsClientMockRecorder struct {
	mock *MockApplicationsClient
}

// NewMockApplicationsClient creates a new mock instance
func NewMockApplicationsClient(ctrl *gomock.Controller) *MockApplicationsClient {
	mock := &MockApplicationsClient{ctrl: ctrl}
	mock.recorder = &MockApplicationsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApplicationsClient) EXPECT() *MockApplicationsClientMockRecorder {
	return m.recorder
}

// Client mocks base method
func (m *MockApplicationsClient) Client() autorest.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Client")
	ret0, _ := ret[0].(autorest.Client)
	return ret0
}

// Client indicates an expected call of Client
func (mr *MockApplicationsClientMockRecorder) Client() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockApplicationsClient)(nil).Client))
}

// Get mocks base method
func (m *MockApplicationsClient) Get(arg0 context.Context, arg1, arg2 string) (managedapplications.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(managedapplications.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockApplicationsClientMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockApplicationsClient)(nil).Get), arg0, arg1, arg2)
}

// GetByID mocks base method
func (m *MockApplicationsClient) GetByID(arg0 context.Context, arg1 string) (managedapplications.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0, arg1)
	ret0, _ := ret[0].(managedapplications.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockApplicationsClientMockRecorder) GetByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockApplicationsClient)(nil).GetByID), arg0, arg1)
}

// ListByResourceGroup mocks base method
func (m *MockApplicationsClient) ListByResourceGroup(arg0 context.Context, arg1 string) (managedapplications.ApplicationListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByResourceGroup", arg0, arg1)
	ret0, _ := ret[0].(managedapplications.ApplicationListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByResourceGroup indicates an expected call of ListByResourceGroup
func (mr *MockApplicationsClientMockRecorder) ListByResourceGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByResourceGroup", reflect.TypeOf((*MockApplicationsClient)(nil).ListByResourceGroup), arg0, arg1)
}