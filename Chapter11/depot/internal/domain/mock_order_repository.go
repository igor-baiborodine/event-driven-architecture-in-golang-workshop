// Code generated by mockery v2.40.1. DO NOT EDIT.

package domain

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockOrderRepository is an autogenerated mock type for the OrderRepository type
type MockOrderRepository struct {
	mock.Mock
}

// Ready provides a mock function with given fields: ctx, orderID
func (_m *MockOrderRepository) Ready(ctx context.Context, orderID string) error {
	ret := _m.Called(ctx, orderID)

	if len(ret) == 0 {
		panic("no return value specified for Ready")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, orderID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockOrderRepository creates a new instance of MockOrderRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockOrderRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockOrderRepository {
	mock := &MockOrderRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
