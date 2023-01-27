// Code generated by mockery v2.10.0. DO NOT EDIT.

package repository

import mock "github.com/stretchr/testify/mock"

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

// Add provides a mock function with given fields: article
func (_m *MockRepository) Add(article Article) {
	_m.Called(article)
}

// Delete provides a mock function with given fields: id
func (_m *MockRepository) Delete(id string) {
	_m.Called(id)
}

// Get provides a mock function with given fields: id
func (_m *MockRepository) Get(id string) Article {
	ret := _m.Called(id)

	var r0 Article
	if rf, ok := ret.Get(0).(func(string) Article); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(Article)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *MockRepository) GetAll() []Article {
	ret := _m.Called()

	var r0 []Article
	if rf, ok := ret.Get(0).(func() []Article); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Article)
		}
	}

	return r0
}
