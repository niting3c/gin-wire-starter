// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"

	utils "starter/internal/app/utils"
)

// RestCaller is an autogenerated mock type for the RestCaller type
type RestCaller struct {
	mock.Mock
}

// Get provides a mock function with given fields: url
func (_m *RestCaller) Get(url string) (*http.Response, error) {
	ret := _m.Called(url)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*http.Response, error)); ok {
		return rf(url)
	}
	if rf, ok := ret.Get(0).(func(string) *http.Response); ok {
		r0 = rf(url)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MakeRestCallToPartner provides a mock function with given fields: url, params
func (_m *RestCaller) MakeRestCallToPartner(url string, params string) *utils.ErrorMessage {
	ret := _m.Called(url, params)

	if len(ret) == 0 {
		panic("no return value specified for MakeRestCallToPartner")
	}

	var r0 *utils.ErrorMessage
	if rf, ok := ret.Get(0).(func(string, string) *utils.ErrorMessage); ok {
		r0 = rf(url, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*utils.ErrorMessage)
		}
	}

	return r0
}

// NewRestCaller creates a new instance of RestCaller. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRestCaller(t interface {
	mock.TestingT
	Cleanup(func())
}) *RestCaller {
	mock := &RestCaller{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
