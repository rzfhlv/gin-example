// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// IAuth is an autogenerated mock type for the IAuth type
type IAuth struct {
	mock.Mock
}

// Bearer provides a mock function with given fields:
func (_m *IAuth) Bearer() gin.HandlerFunc {
	ret := _m.Called()

	var r0 gin.HandlerFunc
	if rf, ok := ret.Get(0).(func() gin.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gin.HandlerFunc)
		}
	}

	return r0
}

// NewIAuth creates a new instance of IAuth. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIAuth(t interface {
	mock.TestingT
	Cleanup(func())
}) *IAuth {
	mock := &IAuth{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}