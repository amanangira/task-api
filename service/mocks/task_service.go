// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	context "context"
	models "task/persistence/models"

	mock "github.com/stretchr/testify/mock"
)

// ITaskService is an autogenerated mock type for the ITaskService type
type ITaskService struct {
	mock.Mock
}

type ITaskService_Expecter struct {
	mock *mock.Mock
}

func (_m *ITaskService) EXPECT() *ITaskService_Expecter {
	return &ITaskService_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, input
func (_m *ITaskService) Create(ctx context.Context, input models.Task) (string, error) {
	ret := _m.Called(ctx, input)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Task) (string, error)); ok {
		return rf(ctx, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.Task) string); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.Task) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ITaskService_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type ITaskService_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - input models.Task
func (_e *ITaskService_Expecter) Create(ctx interface{}, input interface{}) *ITaskService_Create_Call {
	return &ITaskService_Create_Call{Call: _e.mock.On("Create", ctx, input)}
}

func (_c *ITaskService_Create_Call) Run(run func(ctx context.Context, input models.Task)) *ITaskService_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.Task))
	})
	return _c
}

func (_c *ITaskService_Create_Call) Return(_a0 string, _a1 error) *ITaskService_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ITaskService_Create_Call) RunAndReturn(run func(context.Context, models.Task) (string, error)) *ITaskService_Create_Call {
	_c.Call.Return(run)
	return _c
}

// NewITaskService creates a new instance of ITaskService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewITaskService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ITaskService {
	mock := &ITaskService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}