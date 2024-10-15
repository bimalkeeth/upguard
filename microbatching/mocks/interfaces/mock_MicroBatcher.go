// Code generated by mockery v2.46.3. DO NOT EDIT.

package interfaces

import (
	interfaces "github.com/bimalkeeth/upguard/microbatching/interfaces"
	mock "github.com/stretchr/testify/mock"
)

// MockMicroBatcher is an autogenerated mock type for the MicroBatcher type
type MockMicroBatcher[T any, R any] struct {
	mock.Mock
}

type MockMicroBatcher_Expecter[T any, R any] struct {
	mock *mock.Mock
}

func (_m *MockMicroBatcher[T, R]) EXPECT() *MockMicroBatcher_Expecter[T, R] {
	return &MockMicroBatcher_Expecter[T, R]{mock: &_m.Mock}
}

// ReadResult provides a mock function with given fields:
func (_m *MockMicroBatcher[T, R]) ReadResult() <-chan interfaces.JobResult[R] {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ReadResult")
	}

	var r0 <-chan interfaces.JobResult[R]
	if rf, ok := ret.Get(0).(func() <-chan interfaces.JobResult[R]); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan interfaces.JobResult[R])
		}
	}

	return r0
}

// MockMicroBatcher_ReadResult_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadResult'
type MockMicroBatcher_ReadResult_Call[T any, R any] struct {
	*mock.Call
}

// ReadResult is a helper method to define mock.On call
func (_e *MockMicroBatcher_Expecter[T, R]) ReadResult() *MockMicroBatcher_ReadResult_Call[T, R] {
	return &MockMicroBatcher_ReadResult_Call[T, R]{Call: _e.mock.On("ReadResult")}
}

func (_c *MockMicroBatcher_ReadResult_Call[T, R]) Run(run func()) *MockMicroBatcher_ReadResult_Call[T, R] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockMicroBatcher_ReadResult_Call[T, R]) Return(_a0 <-chan interfaces.JobResult[R]) *MockMicroBatcher_ReadResult_Call[T, R] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMicroBatcher_ReadResult_Call[T, R]) RunAndReturn(run func() <-chan interfaces.JobResult[R]) *MockMicroBatcher_ReadResult_Call[T, R] {
	_c.Call.Return(run)
	return _c
}

// Shutdown provides a mock function with given fields:
func (_m *MockMicroBatcher[T, R]) Shutdown() {
	_m.Called()
}

// MockMicroBatcher_Shutdown_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Shutdown'
type MockMicroBatcher_Shutdown_Call[T any, R any] struct {
	*mock.Call
}

// Shutdown is a helper method to define mock.On call
func (_e *MockMicroBatcher_Expecter[T, R]) Shutdown() *MockMicroBatcher_Shutdown_Call[T, R] {
	return &MockMicroBatcher_Shutdown_Call[T, R]{Call: _e.mock.On("Shutdown")}
}

func (_c *MockMicroBatcher_Shutdown_Call[T, R]) Run(run func()) *MockMicroBatcher_Shutdown_Call[T, R] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockMicroBatcher_Shutdown_Call[T, R]) Return() *MockMicroBatcher_Shutdown_Call[T, R] {
	_c.Call.Return()
	return _c
}

func (_c *MockMicroBatcher_Shutdown_Call[T, R]) RunAndReturn(run func()) *MockMicroBatcher_Shutdown_Call[T, R] {
	_c.Call.Return(run)
	return _c
}

// Submit provides a mock function with given fields: job
func (_m *MockMicroBatcher[T, R]) Submit(job interfaces.Job[T, R]) error {
	ret := _m.Called(job)

	if len(ret) == 0 {
		panic("no return value specified for Submit")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interfaces.Job[T, R]) error); ok {
		r0 = rf(job)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockMicroBatcher_Submit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Submit'
type MockMicroBatcher_Submit_Call[T any, R any] struct {
	*mock.Call
}

// Submit is a helper method to define mock.On call
//   - job interfaces.Job[T,R]
func (_e *MockMicroBatcher_Expecter[T, R]) Submit(job interface{}) *MockMicroBatcher_Submit_Call[T, R] {
	return &MockMicroBatcher_Submit_Call[T, R]{Call: _e.mock.On("Submit", job)}
}

func (_c *MockMicroBatcher_Submit_Call[T, R]) Run(run func(job interfaces.Job[T, R])) *MockMicroBatcher_Submit_Call[T, R] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interfaces.Job[T, R]))
	})
	return _c
}

func (_c *MockMicroBatcher_Submit_Call[T, R]) Return(_a0 error) *MockMicroBatcher_Submit_Call[T, R] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMicroBatcher_Submit_Call[T, R]) RunAndReturn(run func(interfaces.Job[T, R]) error) *MockMicroBatcher_Submit_Call[T, R] {
	_c.Call.Return(run)
	return _c
}

// NewMockMicroBatcher creates a new instance of MockMicroBatcher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMicroBatcher[T any, R any](t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMicroBatcher[T, R] {
	mock := &MockMicroBatcher[T, R]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
