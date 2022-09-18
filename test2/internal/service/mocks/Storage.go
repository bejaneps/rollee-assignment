// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Storage is an autogenerated mock type for the Storage type
type Storage struct {
	mock.Mock
}

// GetMostFrequentWord provides a mock function with given fields: beginning
func (_m *Storage) GetMostFrequentWord(beginning string) (string, error) {
	ret := _m.Called(beginning)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(beginning)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(beginning)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpsertWord provides a mock function with given fields: word
func (_m *Storage) UpsertWord(word string) error {
	ret := _m.Called(word)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(word)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewStorage interface {
	mock.TestingT
	Cleanup(func())
}

// NewStorage creates a new instance of Storage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStorage(t mockConstructorTestingTNewStorage) *Storage {
	mock := &Storage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
