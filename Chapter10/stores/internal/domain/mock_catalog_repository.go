// Code generated by mockery v2.40.1. DO NOT EDIT.

package domain

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockCatalogRepository is an autogenerated mock type for the CatalogRepository type
type MockCatalogRepository struct {
	mock.Mock
}

// AddProduct provides a mock function with given fields: ctx, productID, storeID, name, description, sku, price
func (_m *MockCatalogRepository) AddProduct(ctx context.Context, productID string, storeID string, name string, description string, sku string, price float64) error {
	ret := _m.Called(ctx, productID, storeID, name, description, sku, price)

	if len(ret) == 0 {
		panic("no return value specified for AddProduct")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, string, float64) error); ok {
		r0 = rf(ctx, productID, storeID, name, description, sku, price)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: ctx, productID
func (_m *MockCatalogRepository) Find(ctx context.Context, productID string) (*CatalogProduct, error) {
	ret := _m.Called(ctx, productID)

	if len(ret) == 0 {
		panic("no return value specified for Find")
	}

	var r0 *CatalogProduct
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*CatalogProduct, error)); ok {
		return rf(ctx, productID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *CatalogProduct); ok {
		r0 = rf(ctx, productID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*CatalogProduct)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, productID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCatalog provides a mock function with given fields: ctx, storeID
func (_m *MockCatalogRepository) GetCatalog(ctx context.Context, storeID string) ([]*CatalogProduct, error) {
	ret := _m.Called(ctx, storeID)

	if len(ret) == 0 {
		panic("no return value specified for GetCatalog")
	}

	var r0 []*CatalogProduct
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*CatalogProduct, error)); ok {
		return rf(ctx, storeID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*CatalogProduct); ok {
		r0 = rf(ctx, storeID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*CatalogProduct)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, storeID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Rebrand provides a mock function with given fields: ctx, productID, name, description
func (_m *MockCatalogRepository) Rebrand(ctx context.Context, productID string, name string, description string) error {
	ret := _m.Called(ctx, productID, name, description)

	if len(ret) == 0 {
		panic("no return value specified for Rebrand")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, productID, name, description)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveProduct provides a mock function with given fields: ctx, productID
func (_m *MockCatalogRepository) RemoveProduct(ctx context.Context, productID string) error {
	ret := _m.Called(ctx, productID)

	if len(ret) == 0 {
		panic("no return value specified for RemoveProduct")
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
func (_m *MockCatalogRepository) UpdatePrice(ctx context.Context, productID string, delta float64) error {
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

// NewMockCatalogRepository creates a new instance of MockCatalogRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCatalogRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCatalogRepository {
	mock := &MockCatalogRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
