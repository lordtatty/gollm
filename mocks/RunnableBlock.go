// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	gollm "github.com/lordtatty/gollm"
	mock "github.com/stretchr/testify/mock"
)

// RunnableBlock is an autogenerated mock type for the RunnableBlock type
type RunnableBlock struct {
	mock.Mock
}

type RunnableBlock_Expecter struct {
	mock *mock.Mock
}

func (_m *RunnableBlock) EXPECT() *RunnableBlock_Expecter {
	return &RunnableBlock_Expecter{mock: &_m.Mock}
}

// Run provides a mock function with given fields: _a0
func (_m *RunnableBlock) Run(_a0 map[string]string) (*gollm.BlockResult, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Run")
	}

	var r0 *gollm.BlockResult
	var r1 error
	if rf, ok := ret.Get(0).(func(map[string]string) (*gollm.BlockResult, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(map[string]string) *gollm.BlockResult); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gollm.BlockResult)
		}
	}

	if rf, ok := ret.Get(1).(func(map[string]string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RunnableBlock_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type RunnableBlock_Run_Call struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
//   - _a0 map[string]string
func (_e *RunnableBlock_Expecter) Run(_a0 interface{}) *RunnableBlock_Run_Call {
	return &RunnableBlock_Run_Call{Call: _e.mock.On("Run", _a0)}
}

func (_c *RunnableBlock_Run_Call) Run(run func(_a0 map[string]string)) *RunnableBlock_Run_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(map[string]string))
	})
	return _c
}

func (_c *RunnableBlock_Run_Call) Return(_a0 *gollm.BlockResult, _a1 error) *RunnableBlock_Run_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RunnableBlock_Run_Call) RunAndReturn(run func(map[string]string) (*gollm.BlockResult, error)) *RunnableBlock_Run_Call {
	_c.Call.Return(run)
	return _c
}

// UniqName provides a mock function with given fields:
func (_m *RunnableBlock) UniqName() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for UniqName")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// RunnableBlock_UniqName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UniqName'
type RunnableBlock_UniqName_Call struct {
	*mock.Call
}

// UniqName is a helper method to define mock.On call
func (_e *RunnableBlock_Expecter) UniqName() *RunnableBlock_UniqName_Call {
	return &RunnableBlock_UniqName_Call{Call: _e.mock.On("UniqName")}
}

func (_c *RunnableBlock_UniqName_Call) Run(run func()) *RunnableBlock_UniqName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RunnableBlock_UniqName_Call) Return(_a0 string) *RunnableBlock_UniqName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RunnableBlock_UniqName_Call) RunAndReturn(run func() string) *RunnableBlock_UniqName_Call {
	_c.Call.Return(run)
	return _c
}

// NewRunnableBlock creates a new instance of RunnableBlock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRunnableBlock(t interface {
	mock.TestingT
	Cleanup(func())
}) *RunnableBlock {
	mock := &RunnableBlock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
