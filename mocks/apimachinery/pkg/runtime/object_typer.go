// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	runtime "k8s.io/apimachinery/pkg/runtime"

	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ObjectTyper is an autogenerated mock type for the ObjectTyper type
type ObjectTyper struct {
	mock.Mock
}

// ObjectKinds provides a mock function with given fields: _a0
func (_m *ObjectTyper) ObjectKinds(_a0 runtime.Object) ([]schema.GroupVersionKind, bool, error) {
	ret := _m.Called(_a0)

	var r0 []schema.GroupVersionKind
	if rf, ok := ret.Get(0).(func(runtime.Object) []schema.GroupVersionKind); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]schema.GroupVersionKind)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(runtime.Object) bool); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(runtime.Object) error); ok {
		r2 = rf(_a0)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Recognizes provides a mock function with given fields: gvk
func (_m *ObjectTyper) Recognizes(gvk schema.GroupVersionKind) bool {
	ret := _m.Called(gvk)

	var r0 bool
	if rf, ok := ret.Get(0).(func(schema.GroupVersionKind) bool); ok {
		r0 = rf(gvk)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTNewObjectTyper interface {
	mock.TestingT
	Cleanup(func())
}

// NewObjectTyper creates a new instance of ObjectTyper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewObjectTyper(t mockConstructorTestingTNewObjectTyper) *ObjectTyper {
	mock := &ObjectTyper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
