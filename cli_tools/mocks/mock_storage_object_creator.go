//  Copyright 2020 Google Inc. All Rights Reserved.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/GoogleCloudPlatform/compute-image-import/cli_tools/common/domain (interfaces: StorageObjectCreatorInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	domain "github.com/GoogleCloudPlatform/compute-image-import/cli_tools/common/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockStorageObjectCreatorInterface is a mock of StorageObjectCreatorInterface interface.
type MockStorageObjectCreatorInterface struct {
	ctrl     *gomock.Controller
	recorder *MockStorageObjectCreatorInterfaceMockRecorder
}

// MockStorageObjectCreatorInterfaceMockRecorder is the mock recorder for MockStorageObjectCreatorInterface.
type MockStorageObjectCreatorInterfaceMockRecorder struct {
	mock *MockStorageObjectCreatorInterface
}

// NewMockStorageObjectCreatorInterface creates a new mock instance.
func NewMockStorageObjectCreatorInterface(ctrl *gomock.Controller) *MockStorageObjectCreatorInterface {
	mock := &MockStorageObjectCreatorInterface{ctrl: ctrl}
	mock.recorder = &MockStorageObjectCreatorInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageObjectCreatorInterface) EXPECT() *MockStorageObjectCreatorInterfaceMockRecorder {
	return m.recorder
}

// GetObject mocks base method.
func (m *MockStorageObjectCreatorInterface) GetObject(arg0, arg1 string) domain.StorageObject {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObject", arg0, arg1)
	ret0, _ := ret[0].(domain.StorageObject)
	return ret0
}

// GetObject indicates an expected call of GetObject.
func (mr *MockStorageObjectCreatorInterfaceMockRecorder) GetObject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObject", reflect.TypeOf((*MockStorageObjectCreatorInterface)(nil).GetObject), arg0, arg1)
}
