// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-vpc-resource-controller-k8s/pkg/k8s (interfaces: K8sWrapper)

// Package mock_k8s is a generated GoMock package.
package mock_k8s

import (
	reflect "reflect"

	v1alpha1 "github.com/aws/amazon-vpc-cni-k8s/pkg/apis/crd/v1alpha1"
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/apps/v1"
	v10 "k8s.io/api/core/v1"
	v11 "k8s.io/api/events/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockK8sWrapper is a mock of K8sWrapper interface.
type MockK8sWrapper struct {
	ctrl     *gomock.Controller
	recorder *MockK8sWrapperMockRecorder
}

// MockK8sWrapperMockRecorder is the mock recorder for MockK8sWrapper.
type MockK8sWrapperMockRecorder struct {
	mock *MockK8sWrapper
}

// NewMockK8sWrapper creates a new mock instance.
func NewMockK8sWrapper(ctrl *gomock.Controller) *MockK8sWrapper {
	mock := &MockK8sWrapper{ctrl: ctrl}
	mock.recorder = &MockK8sWrapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockK8sWrapper) EXPECT() *MockK8sWrapperMockRecorder {
	return m.recorder
}

// AddLabelToManageNode mocks base method.
func (m *MockK8sWrapper) AddLabelToManageNode(arg0 *v10.Node, arg1, arg2 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddLabelToManageNode", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddLabelToManageNode indicates an expected call of AddLabelToManageNode.
func (mr *MockK8sWrapperMockRecorder) AddLabelToManageNode(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLabelToManageNode", reflect.TypeOf((*MockK8sWrapper)(nil).AddLabelToManageNode), arg0, arg1, arg2)
}

// AdvertiseCapacityIfNotSet mocks base method.
func (m *MockK8sWrapper) AdvertiseCapacityIfNotSet(arg0, arg1 string, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdvertiseCapacityIfNotSet", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AdvertiseCapacityIfNotSet indicates an expected call of AdvertiseCapacityIfNotSet.
func (mr *MockK8sWrapperMockRecorder) AdvertiseCapacityIfNotSet(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdvertiseCapacityIfNotSet", reflect.TypeOf((*MockK8sWrapper)(nil).AdvertiseCapacityIfNotSet), arg0, arg1, arg2)
}

// BroadcastEvent mocks base method.
func (m *MockK8sWrapper) BroadcastEvent(arg0 runtime.Object, arg1, arg2, arg3 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BroadcastEvent", arg0, arg1, arg2, arg3)
}

// BroadcastEvent indicates an expected call of BroadcastEvent.
func (mr *MockK8sWrapperMockRecorder) BroadcastEvent(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BroadcastEvent", reflect.TypeOf((*MockK8sWrapper)(nil).BroadcastEvent), arg0, arg1, arg2, arg3)
}

// GetConfigMap mocks base method.
func (m *MockK8sWrapper) GetConfigMap(arg0, arg1 string) (*v10.ConfigMap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfigMap", arg0, arg1)
	ret0, _ := ret[0].(*v10.ConfigMap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfigMap indicates an expected call of GetConfigMap.
func (mr *MockK8sWrapperMockRecorder) GetConfigMap(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfigMap", reflect.TypeOf((*MockK8sWrapper)(nil).GetConfigMap), arg0, arg1)
}

// GetDaemonSet mocks base method.
func (m *MockK8sWrapper) GetDaemonSet(arg0, arg1 string) (*v1.DaemonSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDaemonSet", arg0, arg1)
	ret0, _ := ret[0].(*v1.DaemonSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDaemonSet indicates an expected call of GetDaemonSet.
func (mr *MockK8sWrapperMockRecorder) GetDaemonSet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDaemonSet", reflect.TypeOf((*MockK8sWrapper)(nil).GetDaemonSet), arg0, arg1)
}

// GetDeployment mocks base method.
func (m *MockK8sWrapper) GetDeployment(arg0, arg1 string) (*v1.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeployment", arg0, arg1)
	ret0, _ := ret[0].(*v1.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeployment indicates an expected call of GetDeployment.
func (mr *MockK8sWrapperMockRecorder) GetDeployment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeployment", reflect.TypeOf((*MockK8sWrapper)(nil).GetDeployment), arg0, arg1)
}

// GetENIConfig mocks base method.
func (m *MockK8sWrapper) GetENIConfig(arg0 string) (*v1alpha1.ENIConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetENIConfig", arg0)
	ret0, _ := ret[0].(*v1alpha1.ENIConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetENIConfig indicates an expected call of GetENIConfig.
func (mr *MockK8sWrapperMockRecorder) GetENIConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetENIConfig", reflect.TypeOf((*MockK8sWrapper)(nil).GetENIConfig), arg0)
}

// GetNode mocks base method.
func (m *MockK8sWrapper) GetNode(arg0 string) (*v10.Node, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNode", arg0)
	ret0, _ := ret[0].(*v10.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNode indicates an expected call of GetNode.
func (mr *MockK8sWrapperMockRecorder) GetNode(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNode", reflect.TypeOf((*MockK8sWrapper)(nil).GetNode), arg0)
}

// ListEvents mocks base method.
func (m *MockK8sWrapper) ListEvents(arg0 []client.ListOption) (*v11.EventList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEvents", arg0)
	ret0, _ := ret[0].(*v11.EventList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEvents indicates an expected call of ListEvents.
func (mr *MockK8sWrapperMockRecorder) ListEvents(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEvents", reflect.TypeOf((*MockK8sWrapper)(nil).ListEvents), arg0)
}

// ListNodes mocks base method.
func (m *MockK8sWrapper) ListNodes() (*v10.NodeList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNodes")
	ret0, _ := ret[0].(*v10.NodeList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNodes indicates an expected call of ListNodes.
func (mr *MockK8sWrapperMockRecorder) ListNodes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNodes", reflect.TypeOf((*MockK8sWrapper)(nil).ListNodes))
}
