// Code generated by mockery v2.14.0. DO NOT EDIT.

package domain

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockOrderRepository is an autogenerated mock type for the OrderRepository type
type MockOrderRepository struct {
	mock.Mock
}

// Load provides a mock function with given fields: ctx, orderID
func (_m *MockOrderRepository) Load(ctx context.Context, orderID string) (*Order, error) {
	ret := _m.Called(ctx, orderID)

	var r0 *Order
	if rf, ok := ret.Get(0).(func(context.Context, string) *Order); ok {
		r0 = rf(ctx, orderID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, orderID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, order
func (_m *MockOrderRepository) Save(ctx context.Context, order *Order) error {
	ret := _m.Called(ctx, order)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *Order) error); ok {
		r0 = rf(ctx, order)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockOrderRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockOrderRepository creates a new instance of MockOrderRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockOrderRepository(t mockConstructorTestingTNewMockOrderRepository) *MockOrderRepository {
	mock := &MockOrderRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
