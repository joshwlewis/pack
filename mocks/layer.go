// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/google/go-containerregistry/pkg/v1 (interfaces: Layer)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	io "io"
	reflect "reflect"
)

// MockLayer is a mock of Layer interface
type MockLayer struct {
	ctrl     *gomock.Controller
	recorder *MockLayerMockRecorder
}

// MockLayerMockRecorder is the mock recorder for MockLayer
type MockLayerMockRecorder struct {
	mock *MockLayer
}

// NewMockLayer creates a new mock instance
func NewMockLayer(ctrl *gomock.Controller) *MockLayer {
	mock := &MockLayer{ctrl: ctrl}
	mock.recorder = &MockLayerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLayer) EXPECT() *MockLayerMockRecorder {
	return m.recorder
}

// Compressed mocks base method
func (m *MockLayer) Compressed() (io.ReadCloser, error) {
	ret := m.ctrl.Call(m, "Compressed")
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Compressed indicates an expected call of Compressed
func (mr *MockLayerMockRecorder) Compressed() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Compressed", reflect.TypeOf((*MockLayer)(nil).Compressed))
}

// DiffID mocks base method
func (m *MockLayer) DiffID() (v1.Hash, error) {
	ret := m.ctrl.Call(m, "DiffID")
	ret0, _ := ret[0].(v1.Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DiffID indicates an expected call of DiffID
func (mr *MockLayerMockRecorder) DiffID() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiffID", reflect.TypeOf((*MockLayer)(nil).DiffID))
}

// Digest mocks base method
func (m *MockLayer) Digest() (v1.Hash, error) {
	ret := m.ctrl.Call(m, "Digest")
	ret0, _ := ret[0].(v1.Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Digest indicates an expected call of Digest
func (mr *MockLayerMockRecorder) Digest() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Digest", reflect.TypeOf((*MockLayer)(nil).Digest))
}

// Size mocks base method
func (m *MockLayer) Size() (int64, error) {
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Size indicates an expected call of Size
func (mr *MockLayerMockRecorder) Size() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockLayer)(nil).Size))
}

// Uncompressed mocks base method
func (m *MockLayer) Uncompressed() (io.ReadCloser, error) {
	ret := m.ctrl.Call(m, "Uncompressed")
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Uncompressed indicates an expected call of Uncompressed
func (mr *MockLayerMockRecorder) Uncompressed() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Uncompressed", reflect.TypeOf((*MockLayer)(nil).Uncompressed))
}