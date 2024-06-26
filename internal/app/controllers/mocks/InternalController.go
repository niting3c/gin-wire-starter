// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// InternalController is an autogenerated mock type for the InternalController type
type InternalController struct {
	mock.Mock
}

// HealthCheck provides a mock function with given fields: c
func (_m *InternalController) HealthCheck(c *gin.Context) {
	_m.Called(c)
}

// SetLogLevel provides a mock function with given fields: c
func (_m *InternalController) SetLogLevel(c *gin.Context) {
	_m.Called(c)
}

// NewInternalController creates a new instance of InternalController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewInternalController(t interface {
	mock.TestingT
	Cleanup(func())
}) *InternalController {
	mock := &InternalController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
