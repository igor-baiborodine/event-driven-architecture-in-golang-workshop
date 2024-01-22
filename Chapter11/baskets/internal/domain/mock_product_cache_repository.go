// Code generated by mockery v2.40.1. DO NOT EDIT.

package domain

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockProductCacheRepository is an autogenerated mock type for the ProductCacheRepository type
type MockProductCacheRepository struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, productID, storeID, name, price
func (_m *MockProductCacheRepository) Add(ctx context.Context, productID string, storeID string, name string, price float64) error {
	ret := _m.Called(ctx, productID, storeID, name, price)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, float64) error); ok {
		r0 = rf(ctx, productID, storeID, name, price)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: ctx, productID
func (_m *MockProductCacheRepository) Find(ctx context.Context, productID string) (*Product, error) {
	ret := _m.Called(ctx, productID)

	if len(ret) == 0 {
		panic("no return value specified for Find")
	}

	var r0 *Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*Product, error)); ok {
		return rf(ctx, productID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *Product); ok {
		r0 = rf(ctx, productID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Product)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, productID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Rebrand provides a mock function with given fields: ctx, productID, name
func (_m *MockProductCacheRepository) Rebrand(ctx context.Context, productID string, name string) error {
	ret := _m.Called(ctx, productID, name)

	if len(ret) == 0 {
		panic("no return value specified for Rebrand")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, productID, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Remove provides a mock function with given fields: ctx, productID
func (_m *MockProductCacheRepository) Remove(ctx context.Context, productID string) error {
	ret := _m.Called(ctx, productID)

	if len(ret) == 0 {
		panic("no return value specified for Remove")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, productID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdatePrice provides a mock function with given fields: ctx, productID, delta
func (_m *MockProductCacheRepository) UpdatePrice(ctx context.Context, productID string, delta float64) error {
	ret := _m.Called(ctx, productID, delta)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePrice")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, float64) error); ok {
		r0 = rf(ctx, productID, delta)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockProductCacheRepository creates a new instance of MockProductCacheRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockProductCacheRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockProductCacheRepository {
	mock := &MockProductCacheRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}