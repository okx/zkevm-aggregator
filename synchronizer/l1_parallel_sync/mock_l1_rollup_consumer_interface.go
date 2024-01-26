// Code generated by mockery. DO NOT EDIT.

package l1_parallel_sync

import (
	context "context"

	state "github.com/0xPolygonHermez/zkevm-aggregator/state"
	mock "github.com/stretchr/testify/mock"
)

// l1RollupConsumerInterfaceMock is an autogenerated mock type for the l1RollupConsumerInterface type
type l1RollupConsumerInterfaceMock struct {
	mock.Mock
}

type l1RollupConsumerInterfaceMock_Expecter struct {
	mock *mock.Mock
}

func (_m *l1RollupConsumerInterfaceMock) EXPECT() *l1RollupConsumerInterfaceMock_Expecter {
	return &l1RollupConsumerInterfaceMock_Expecter{mock: &_m.Mock}
}

// GetLastEthBlockSynced provides a mock function with given fields:
func (_m *l1RollupConsumerInterfaceMock) GetLastEthBlockSynced() (state.Block, bool) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetLastEthBlockSynced")
	}

	var r0 state.Block
	var r1 bool
	if rf, ok := ret.Get(0).(func() (state.Block, bool)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() state.Block); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(state.Block)
	}

	if rf, ok := ret.Get(1).(func() bool); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// l1RollupConsumerInterfaceMock_GetLastEthBlockSynced_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLastEthBlockSynced'
type l1RollupConsumerInterfaceMock_GetLastEthBlockSynced_Call struct {
	*mock.Call
}

// GetLastEthBlockSynced is a helper method to define mock.On call
func (_e *l1RollupConsumerInterfaceMock_Expecter) GetLastEthBlockSynced() *l1RollupConsumerInterfaceMock_GetLastEthBlockSynced_Call {
	return &l1RollupConsumerInterfaceMock_GetLastEthBlockSynced_Call{Call: _e.mock.On("GetLastEthBlockSynced")}
}

func (_c *l1RollupConsumerInterfaceMock_GetLastEthBlockSynced_Call) Run(run func()) *l1RollupConsumerInterfaceMock_GetLastEthBlockSynced_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *l1RollupConsumerInterfaceMock_GetLastEthBlockSynced_Call) Return(_a0 state.Block, _a1 bool) *l1RollupConsumerInterfaceMock_GetLastEthBlockSynced_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *l1RollupConsumerInterfaceMock_GetLastEthBlockSynced_Call) RunAndReturn(run func() (state.Block, bool)) *l1RollupConsumerInterfaceMock_GetLastEthBlockSynced_Call {
	_c.Call.Return(run)
	return _c
}

// Reset provides a mock function with given fields: startingBlockNumber
func (_m *l1RollupConsumerInterfaceMock) Reset(startingBlockNumber uint64) {
	_m.Called(startingBlockNumber)
}

// l1RollupConsumerInterfaceMock_Reset_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Reset'
type l1RollupConsumerInterfaceMock_Reset_Call struct {
	*mock.Call
}

// Reset is a helper method to define mock.On call
//   - startingBlockNumber uint64
func (_e *l1RollupConsumerInterfaceMock_Expecter) Reset(startingBlockNumber interface{}) *l1RollupConsumerInterfaceMock_Reset_Call {
	return &l1RollupConsumerInterfaceMock_Reset_Call{Call: _e.mock.On("Reset", startingBlockNumber)}
}

func (_c *l1RollupConsumerInterfaceMock_Reset_Call) Run(run func(startingBlockNumber uint64)) *l1RollupConsumerInterfaceMock_Reset_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint64))
	})
	return _c
}

func (_c *l1RollupConsumerInterfaceMock_Reset_Call) Return() *l1RollupConsumerInterfaceMock_Reset_Call {
	_c.Call.Return()
	return _c
}

func (_c *l1RollupConsumerInterfaceMock_Reset_Call) RunAndReturn(run func(uint64)) *l1RollupConsumerInterfaceMock_Reset_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields: ctx, lastEthBlockSynced
func (_m *l1RollupConsumerInterfaceMock) Start(ctx context.Context, lastEthBlockSynced *state.Block) error {
	ret := _m.Called(ctx, lastEthBlockSynced)

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *state.Block) error); ok {
		r0 = rf(ctx, lastEthBlockSynced)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// l1RollupConsumerInterfaceMock_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type l1RollupConsumerInterfaceMock_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - ctx context.Context
//   - lastEthBlockSynced *state.Block
func (_e *l1RollupConsumerInterfaceMock_Expecter) Start(ctx interface{}, lastEthBlockSynced interface{}) *l1RollupConsumerInterfaceMock_Start_Call {
	return &l1RollupConsumerInterfaceMock_Start_Call{Call: _e.mock.On("Start", ctx, lastEthBlockSynced)}
}

func (_c *l1RollupConsumerInterfaceMock_Start_Call) Run(run func(ctx context.Context, lastEthBlockSynced *state.Block)) *l1RollupConsumerInterfaceMock_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*state.Block))
	})
	return _c
}

func (_c *l1RollupConsumerInterfaceMock_Start_Call) Return(_a0 error) *l1RollupConsumerInterfaceMock_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *l1RollupConsumerInterfaceMock_Start_Call) RunAndReturn(run func(context.Context, *state.Block) error) *l1RollupConsumerInterfaceMock_Start_Call {
	_c.Call.Return(run)
	return _c
}

// StopAfterProcessChannelQueue provides a mock function with given fields:
func (_m *l1RollupConsumerInterfaceMock) StopAfterProcessChannelQueue() {
	_m.Called()
}

// l1RollupConsumerInterfaceMock_StopAfterProcessChannelQueue_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StopAfterProcessChannelQueue'
type l1RollupConsumerInterfaceMock_StopAfterProcessChannelQueue_Call struct {
	*mock.Call
}

// StopAfterProcessChannelQueue is a helper method to define mock.On call
func (_e *l1RollupConsumerInterfaceMock_Expecter) StopAfterProcessChannelQueue() *l1RollupConsumerInterfaceMock_StopAfterProcessChannelQueue_Call {
	return &l1RollupConsumerInterfaceMock_StopAfterProcessChannelQueue_Call{Call: _e.mock.On("StopAfterProcessChannelQueue")}
}

func (_c *l1RollupConsumerInterfaceMock_StopAfterProcessChannelQueue_Call) Run(run func()) *l1RollupConsumerInterfaceMock_StopAfterProcessChannelQueue_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *l1RollupConsumerInterfaceMock_StopAfterProcessChannelQueue_Call) Return() *l1RollupConsumerInterfaceMock_StopAfterProcessChannelQueue_Call {
	_c.Call.Return()
	return _c
}

func (_c *l1RollupConsumerInterfaceMock_StopAfterProcessChannelQueue_Call) RunAndReturn(run func()) *l1RollupConsumerInterfaceMock_StopAfterProcessChannelQueue_Call {
	_c.Call.Return(run)
	return _c
}

// newL1RollupConsumerInterfaceMock creates a new instance of l1RollupConsumerInterfaceMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newL1RollupConsumerInterfaceMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *l1RollupConsumerInterfaceMock {
	mock := &l1RollupConsumerInterfaceMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}