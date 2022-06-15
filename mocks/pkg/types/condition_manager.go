// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	v1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
)

// ConditionManager is an autogenerated mock type for the ConditionManager type
type ConditionManager struct {
	mock.Mock
}

// Conditions provides a mock function with given fields:
func (_m *ConditionManager) Conditions() []*v1alpha1.Condition {
	ret := _m.Called()

	var r0 []*v1alpha1.Condition
	if rf, ok := ret.Get(0).(func() []*v1alpha1.Condition); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1alpha1.Condition)
		}
	}

	return r0
}

// ReplaceConditions provides a mock function with given fields: _a0
func (_m *ConditionManager) ReplaceConditions(_a0 []*v1alpha1.Condition) {
	_m.Called(_a0)
}

type mockConstructorTestingTNewConditionManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewConditionManager creates a new instance of ConditionManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewConditionManager(t mockConstructorTestingTNewConditionManager) *ConditionManager {
	mock := &ConditionManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
