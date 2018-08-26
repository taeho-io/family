// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import config "github.com/taeho-io/family/services/auth/config"
import mock "github.com/stretchr/testify/mock"

// IFace is an autogenerated mock type for the IFace type
type IFace struct {
	mock.Mock
}

// Env provides a mock function with given fields:
func (_m *IFace) Env() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Prefix provides a mock function with given fields:
func (_m *IFace) Prefix() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ProductName provides a mock function with given fields:
func (_m *IFace) ProductName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ServiceName provides a mock function with given fields:
func (_m *IFace) ServiceName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Settings provides a mock function with given fields:
func (_m *IFace) Settings() config.Settings {
	ret := _m.Called()

	var r0 config.Settings
	if rf, ok := ret.Get(0).(func() config.Settings); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(config.Settings)
	}

	return r0
}
