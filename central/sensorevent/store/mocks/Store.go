// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import mock "github.com/stretchr/testify/mock"

import v1 "bitbucket.org/stack-rox/apollo/generated/api/v1"

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// AddSensorEvent provides a mock function with given fields: deployment
func (_m *Store) AddSensorEvent(deployment *v1.SensorEvent) (uint64, error) {
	ret := _m.Called(deployment)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(*v1.SensorEvent) uint64); ok {
		r0 = rf(deployment)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1.SensorEvent) error); ok {
		r1 = rf(deployment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSensorEvent provides a mock function with given fields: id
func (_m *Store) GetSensorEvent(id uint64) (*v1.SensorEvent, bool, error) {
	ret := _m.Called(id)

	var r0 *v1.SensorEvent
	if rf, ok := ret.Get(0).(func(uint64) *v1.SensorEvent); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.SensorEvent)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(uint64) bool); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(uint64) error); ok {
		r2 = rf(id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetSensorEventIds provides a mock function with given fields: clusterID
func (_m *Store) GetSensorEventIds(clusterID string) ([]uint64, map[string]uint64, error) {
	ret := _m.Called(clusterID)

	var r0 []uint64
	if rf, ok := ret.Get(0).(func(string) []uint64); ok {
		r0 = rf(clusterID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint64)
		}
	}

	var r1 map[string]uint64
	if rf, ok := ret.Get(1).(func(string) map[string]uint64); ok {
		r1 = rf(clusterID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]uint64)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(clusterID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RemoveSensorEvent provides a mock function with given fields: id
func (_m *Store) RemoveSensorEvent(id uint64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateSensorEvent provides a mock function with given fields: id, event
func (_m *Store) UpdateSensorEvent(id uint64, event *v1.SensorEvent) error {
	ret := _m.Called(id, event)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64, *v1.SensorEvent) error); ok {
		r0 = rf(id, event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
