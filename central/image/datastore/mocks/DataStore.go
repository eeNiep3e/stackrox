// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import mock "github.com/stretchr/testify/mock"
import v1 "bitbucket.org/stack-rox/apollo/generated/api/v1"

// DataStore is an autogenerated mock type for the DataStore type
type DataStore struct {
	mock.Mock
}

// CountImages provides a mock function with given fields:
func (_m *DataStore) CountImages() (int, error) {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetImage provides a mock function with given fields: sha
func (_m *DataStore) GetImage(sha string) (*v1.Image, bool, error) {
	ret := _m.Called(sha)

	var r0 *v1.Image
	if rf, ok := ret.Get(0).(func(string) *v1.Image); ok {
		r0 = rf(sha)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Image)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(sha)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(sha)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetImages provides a mock function with given fields:
func (_m *DataStore) GetImages() ([]*v1.Image, error) {
	ret := _m.Called()

	var r0 []*v1.Image
	if rf, ok := ret.Get(0).(func() []*v1.Image); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.Image)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetImagesBatch provides a mock function with given fields: shas
func (_m *DataStore) GetImagesBatch(shas []string) ([]*v1.Image, error) {
	ret := _m.Called(shas)

	var r0 []*v1.Image
	if rf, ok := ret.Get(0).(func([]string) []*v1.Image); ok {
		r0 = rf(shas)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.Image)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(shas)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListImage provides a mock function with given fields: sha
func (_m *DataStore) ListImage(sha string) (*v1.ListImage, bool, error) {
	ret := _m.Called(sha)

	var r0 *v1.ListImage
	if rf, ok := ret.Get(0).(func(string) *v1.ListImage); ok {
		r0 = rf(sha)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ListImage)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(sha)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(sha)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListImages provides a mock function with given fields:
func (_m *DataStore) ListImages() ([]*v1.ListImage, error) {
	ret := _m.Called()

	var r0 []*v1.ListImage
	if rf, ok := ret.Get(0).(func() []*v1.ListImage); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.ListImage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchImages provides a mock function with given fields: request
func (_m *DataStore) SearchImages(request *v1.ParsedSearchRequest) ([]*v1.SearchResult, error) {
	ret := _m.Called(request)

	var r0 []*v1.SearchResult
	if rf, ok := ret.Get(0).(func(*v1.ParsedSearchRequest) []*v1.SearchResult); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.SearchResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1.ParsedSearchRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchListImages provides a mock function with given fields: request
func (_m *DataStore) SearchListImages(request *v1.ParsedSearchRequest) ([]*v1.ListImage, error) {
	ret := _m.Called(request)

	var r0 []*v1.ListImage
	if rf, ok := ret.Get(0).(func(*v1.ParsedSearchRequest) []*v1.ListImage); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.ListImage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1.ParsedSearchRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchRawImages provides a mock function with given fields: request
func (_m *DataStore) SearchRawImages(request *v1.ParsedSearchRequest) ([]*v1.Image, error) {
	ret := _m.Called(request)

	var r0 []*v1.Image
	if rf, ok := ret.Get(0).(func(*v1.ParsedSearchRequest) []*v1.Image); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.Image)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1.ParsedSearchRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpsertDedupeImage provides a mock function with given fields: image
func (_m *DataStore) UpsertDedupeImage(image *v1.Image) error {
	ret := _m.Called(image)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1.Image) error); ok {
		r0 = rf(image)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
