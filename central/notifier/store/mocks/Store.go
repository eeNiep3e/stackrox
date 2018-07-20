// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import mock "github.com/stretchr/testify/mock"

import v1 "bitbucket.org/stack-rox/apollo/generated/api/v1"

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// AddNotifier provides a mock function with given fields: notifier
func (_m *Store) AddNotifier(notifier *v1.Notifier) (string, error) {
	ret := _m.Called(notifier)

	var r0 string
	if rf, ok := ret.Get(0).(func(*v1.Notifier) string); ok {
		r0 = rf(notifier)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1.Notifier) error); ok {
		r1 = rf(notifier)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNotifier provides a mock function with given fields: id
func (_m *Store) GetNotifier(id string) (*v1.Notifier, bool, error) {
	ret := _m.Called(id)

	var r0 *v1.Notifier
	if rf, ok := ret.Get(0).(func(string) *v1.Notifier); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Notifier)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetNotifiers provides a mock function with given fields: request
func (_m *Store) GetNotifiers(request *v1.GetNotifiersRequest) ([]*v1.Notifier, error) {
	ret := _m.Called(request)

	var r0 []*v1.Notifier
	if rf, ok := ret.Get(0).(func(*v1.GetNotifiersRequest) []*v1.Notifier); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.Notifier)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1.GetNotifiersRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveNotifier provides a mock function with given fields: id
func (_m *Store) RemoveNotifier(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateNotifier provides a mock function with given fields: notifier
func (_m *Store) UpdateNotifier(notifier *v1.Notifier) error {
	ret := _m.Called(notifier)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1.Notifier) error); ok {
		r0 = rf(notifier)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
