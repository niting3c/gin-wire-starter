// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	models "starter/internal/app/models"

	mock "github.com/stretchr/testify/mock"

	utils "starter/internal/app/utils"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// GetUserByEmail provides a mock function with given fields: emailId
func (_m *UserService) GetUserByEmail(emailId string) (*models.User, *utils.ErrorMessage) {
	ret := _m.Called(emailId)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByEmail")
	}

	var r0 *models.User
	var r1 *utils.ErrorMessage
	if rf, ok := ret.Get(0).(func(string) (*models.User, *utils.ErrorMessage)); ok {
		return rf(emailId)
	}
	if rf, ok := ret.Get(0).(func(string) *models.User); ok {
		r0 = rf(emailId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) *utils.ErrorMessage); ok {
		r1 = rf(emailId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*utils.ErrorMessage)
		}
	}

	return r0, r1
}

// NewUserService creates a new instance of UserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserService(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserService {
	mock := &UserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
