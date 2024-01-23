// Code generated by mockery v2.40.1. DO NOT EDIT.

package am

import (
	context "context"
	ddd "eda-in-golang/internal/ddd"

	mock "github.com/stretchr/testify/mock"
)

// MockCommandPublisher is an autogenerated mock type for the CommandPublisher type
type MockCommandPublisher struct {
	mock.Mock
}

// Publish provides a mock function with given fields: ctx, topicName, cmd
func (_m *MockCommandPublisher) Publish(ctx context.Context, topicName string, cmd ddd.Command) error {
	ret := _m.Called(ctx, topicName, cmd)

	if len(ret) == 0 {
		panic("no return value specified for Publish")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ddd.Command) error); ok {
		r0 = rf(ctx, topicName, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockCommandPublisher creates a new instance of MockCommandPublisher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCommandPublisher(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCommandPublisher {
	mock := &MockCommandPublisher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}