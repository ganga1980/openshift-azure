// Code generated by MockGen. DO NOT EDIT.
// Source: scaler.go

// Package mock_scaler is a generated GoMock package.
package mock_scaler

import (
	context "context"
	reflect "reflect"

	compute "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	gomock "github.com/golang/mock/gomock"
	logrus "github.com/sirupsen/logrus"

	api "github.com/openshift/openshift-azure/pkg/api"
	kubeclient "github.com/openshift/openshift-azure/pkg/cluster/kubeclient"
	scaler "github.com/openshift/openshift-azure/pkg/cluster/scaler"
	azureclient "github.com/openshift/openshift-azure/pkg/util/azureclient"
)

// MockFactory is a mock of Factory interface
type MockFactory struct {
	ctrl     *gomock.Controller
	recorder *MockFactoryMockRecorder
}

// MockFactoryMockRecorder is the mock recorder for MockFactory
type MockFactoryMockRecorder struct {
	mock *MockFactory
}

// NewMockFactory creates a new mock instance
func NewMockFactory(ctrl *gomock.Controller) *MockFactory {
	mock := &MockFactory{ctrl: ctrl}
	mock.recorder = &MockFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFactory) EXPECT() *MockFactoryMockRecorder {
	return m.recorder
}

// New mocks base method
func (m *MockFactory) New(log *logrus.Entry, ssc azureclient.VirtualMachineScaleSetsClient, vmc azureclient.VirtualMachineScaleSetVMsClient, kubeclient kubeclient.Interface, resourceGroup string, ss *compute.VirtualMachineScaleSet) scaler.Scaler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", log, ssc, vmc, kubeclient, resourceGroup, ss)
	ret0, _ := ret[0].(scaler.Scaler)
	return ret0
}

// New indicates an expected call of New
func (mr *MockFactoryMockRecorder) New(log, ssc, vmc, kubeclient, resourceGroup, ss interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockFactory)(nil).New), log, ssc, vmc, kubeclient, resourceGroup, ss)
}

// MockScaler is a mock of Scaler interface
type MockScaler struct {
	ctrl     *gomock.Controller
	recorder *MockScalerMockRecorder
}

// MockScalerMockRecorder is the mock recorder for MockScaler
type MockScalerMockRecorder struct {
	mock *MockScaler
}

// NewMockScaler creates a new mock instance
func NewMockScaler(ctrl *gomock.Controller) *MockScaler {
	mock := &MockScaler{ctrl: ctrl}
	mock.recorder = &MockScalerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockScaler) EXPECT() *MockScalerMockRecorder {
	return m.recorder
}

// Scale mocks base method
func (m *MockScaler) Scale(ctx context.Context, count int64) *api.PluginError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scale", ctx, count)
	ret0, _ := ret[0].(*api.PluginError)
	return ret0
}

// Scale indicates an expected call of Scale
func (mr *MockScalerMockRecorder) Scale(ctx, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scale", reflect.TypeOf((*MockScaler)(nil).Scale), ctx, count)
}
