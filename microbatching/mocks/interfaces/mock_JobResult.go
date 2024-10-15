// Code generated by mockery v2.46.3. DO NOT EDIT.

package interfaces

import mock "github.com/stretchr/testify/mock"

// MockJobResult is an autogenerated mock type for the JobResult type
type MockJobResult[R any] struct {
	mock.Mock
}

type MockJobResult_Expecter[R any] struct {
	mock *mock.Mock
}

func (_m *MockJobResult[R]) EXPECT() *MockJobResult_Expecter[R] {
	return &MockJobResult_Expecter[R]{mock: &_m.Mock}
}

// GetError provides a mock function with given fields:
func (_m *MockJobResult[R]) GetError() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetError")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockJobResult_GetError_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetError'
type MockJobResult_GetError_Call[R any] struct {
	*mock.Call
}

// GetError is a helper method to define mock.On call
func (_e *MockJobResult_Expecter[R]) GetError() *MockJobResult_GetError_Call[R] {
	return &MockJobResult_GetError_Call[R]{Call: _e.mock.On("GetError")}
}

func (_c *MockJobResult_GetError_Call[R]) Run(run func()) *MockJobResult_GetError_Call[R] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockJobResult_GetError_Call[R]) Return(_a0 error) *MockJobResult_GetError_Call[R] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockJobResult_GetError_Call[R]) RunAndReturn(run func() error) *MockJobResult_GetError_Call[R] {
	_c.Call.Return(run)
	return _c
}

// GetJobID provides a mock function with given fields:
func (_m *MockJobResult[R]) GetJobID() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetJobID")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// MockJobResult_GetJobID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetJobID'
type MockJobResult_GetJobID_Call[R any] struct {
	*mock.Call
}

// GetJobID is a helper method to define mock.On call
func (_e *MockJobResult_Expecter[R]) GetJobID() *MockJobResult_GetJobID_Call[R] {
	return &MockJobResult_GetJobID_Call[R]{Call: _e.mock.On("GetJobID")}
}

func (_c *MockJobResult_GetJobID_Call[R]) Run(run func()) *MockJobResult_GetJobID_Call[R] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockJobResult_GetJobID_Call[R]) Return(_a0 int) *MockJobResult_GetJobID_Call[R] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockJobResult_GetJobID_Call[R]) RunAndReturn(run func() int) *MockJobResult_GetJobID_Call[R] {
	_c.Call.Return(run)
	return _c
}

// GetJobName provides a mock function with given fields:
func (_m *MockJobResult[R]) GetJobName() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetJobName")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockJobResult_GetJobName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetJobName'
type MockJobResult_GetJobName_Call[R any] struct {
	*mock.Call
}

// GetJobName is a helper method to define mock.On call
func (_e *MockJobResult_Expecter[R]) GetJobName() *MockJobResult_GetJobName_Call[R] {
	return &MockJobResult_GetJobName_Call[R]{Call: _e.mock.On("GetJobName")}
}

func (_c *MockJobResult_GetJobName_Call[R]) Run(run func()) *MockJobResult_GetJobName_Call[R] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockJobResult_GetJobName_Call[R]) Return(_a0 string) *MockJobResult_GetJobName_Call[R] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockJobResult_GetJobName_Call[R]) RunAndReturn(run func() string) *MockJobResult_GetJobName_Call[R] {
	_c.Call.Return(run)
	return _c
}

// GetResult provides a mock function with given fields:
func (_m *MockJobResult[R]) GetResult() R {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetResult")
	}

	var r0 R
	if rf, ok := ret.Get(0).(func() R); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(R)
	}

	return r0
}

// MockJobResult_GetResult_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetResult'
type MockJobResult_GetResult_Call[R any] struct {
	*mock.Call
}

// GetResult is a helper method to define mock.On call
func (_e *MockJobResult_Expecter[R]) GetResult() *MockJobResult_GetResult_Call[R] {
	return &MockJobResult_GetResult_Call[R]{Call: _e.mock.On("GetResult")}
}

func (_c *MockJobResult_GetResult_Call[R]) Run(run func()) *MockJobResult_GetResult_Call[R] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockJobResult_GetResult_Call[R]) Return(_a0 R) *MockJobResult_GetResult_Call[R] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockJobResult_GetResult_Call[R]) RunAndReturn(run func() R) *MockJobResult_GetResult_Call[R] {
	_c.Call.Return(run)
	return _c
}

// IsSuccess provides a mock function with given fields:
func (_m *MockJobResult[R]) IsSuccess() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsSuccess")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockJobResult_IsSuccess_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsSuccess'
type MockJobResult_IsSuccess_Call[R any] struct {
	*mock.Call
}

// IsSuccess is a helper method to define mock.On call
func (_e *MockJobResult_Expecter[R]) IsSuccess() *MockJobResult_IsSuccess_Call[R] {
	return &MockJobResult_IsSuccess_Call[R]{Call: _e.mock.On("IsSuccess")}
}

func (_c *MockJobResult_IsSuccess_Call[R]) Run(run func()) *MockJobResult_IsSuccess_Call[R] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockJobResult_IsSuccess_Call[R]) Return(_a0 bool) *MockJobResult_IsSuccess_Call[R] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockJobResult_IsSuccess_Call[R]) RunAndReturn(run func() bool) *MockJobResult_IsSuccess_Call[R] {
	_c.Call.Return(run)
	return _c
}

// NewMockJobResult creates a new instance of MockJobResult. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockJobResult[R any](t interface {
	mock.TestingT
	Cleanup(func())
}) *MockJobResult[R] {
	mock := &MockJobResult[R]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
