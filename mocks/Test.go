// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Test is an autogenerated mock type for the Test type
type Test struct {
	mock.Mock
}

// A provides a mock function with given fields: _a0
func (_m *Test) A(_a0 int) int {
	ret := _m.Called(_a0)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}
