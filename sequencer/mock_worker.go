// Code generated by mockery v2.39.0. DO NOT EDIT.

package sequencer

import (
	context "context"
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	state "github.com/0xPolygonHermez/zkevm-aggregator/state"

	types "github.com/ethereum/go-ethereum/core/types"
)

// WorkerMock is an autogenerated mock type for the workerInterface type
type WorkerMock struct {
	mock.Mock
}

// AddForcedTx provides a mock function with given fields: txHash, addr
func (_m *WorkerMock) AddForcedTx(txHash common.Hash, addr common.Address) {
	_m.Called(txHash, addr)
}

// AddPendingTxToStore provides a mock function with given fields: txHash, addr
func (_m *WorkerMock) AddPendingTxToStore(txHash common.Hash, addr common.Address) {
	_m.Called(txHash, addr)
}

// AddTxTracker provides a mock function with given fields: ctx, txTracker
func (_m *WorkerMock) AddTxTracker(ctx context.Context, txTracker *TxTracker) (*TxTracker, error) {
	ret := _m.Called(ctx, txTracker)

	if len(ret) == 0 {
		panic("no return value specified for AddTxTracker")
	}

	var r0 *TxTracker
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *TxTracker) (*TxTracker, error)); ok {
		return rf(ctx, txTracker)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *TxTracker) *TxTracker); ok {
		r0 = rf(ctx, txTracker)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*TxTracker)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *TxTracker) error); ok {
		r1 = rf(ctx, txTracker)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteForcedTx provides a mock function with given fields: txHash, addr
func (_m *WorkerMock) DeleteForcedTx(txHash common.Hash, addr common.Address) {
	_m.Called(txHash, addr)
}

// DeletePendingTxToStore provides a mock function with given fields: txHash, addr
func (_m *WorkerMock) DeletePendingTxToStore(txHash common.Hash, addr common.Address) {
	_m.Called(txHash, addr)
}

// DeleteTx provides a mock function with given fields: txHash, from
func (_m *WorkerMock) DeleteTx(txHash common.Hash, from common.Address) {
	_m.Called(txHash, from)
}

// GetBestFittingTx provides a mock function with given fields: resources
func (_m *WorkerMock) GetBestFittingTx(resources state.BatchResources) (*TxTracker, error) {
	ret := _m.Called(resources)

	if len(ret) == 0 {
		panic("no return value specified for GetBestFittingTx")
	}

	var r0 *TxTracker
	var r1 error
	if rf, ok := ret.Get(0).(func(state.BatchResources) (*TxTracker, error)); ok {
		return rf(resources)
	}
	if rf, ok := ret.Get(0).(func(state.BatchResources) *TxTracker); ok {
		r0 = rf(resources)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*TxTracker)
		}
	}

	if rf, ok := ret.Get(1).(func(state.BatchResources) error); ok {
		r1 = rf(resources)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MoveTxToNotReady provides a mock function with given fields: txHash, from, actualNonce, actualBalance
func (_m *WorkerMock) MoveTxToNotReady(txHash common.Hash, from common.Address, actualNonce *uint64, actualBalance *big.Int) []*TxTracker {
	ret := _m.Called(txHash, from, actualNonce, actualBalance)

	if len(ret) == 0 {
		panic("no return value specified for MoveTxToNotReady")
	}

	var r0 []*TxTracker
	if rf, ok := ret.Get(0).(func(common.Hash, common.Address, *uint64, *big.Int) []*TxTracker); ok {
		r0 = rf(txHash, from, actualNonce, actualBalance)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*TxTracker)
		}
	}

	return r0
}

// NewTxTracker provides a mock function with given fields: tx, counters, ip
func (_m *WorkerMock) NewTxTracker(tx types.Transaction, counters state.ZKCounters, ip string) (*TxTracker, error) {
	ret := _m.Called(tx, counters, ip)

	if len(ret) == 0 {
		panic("no return value specified for NewTxTracker")
	}

	var r0 *TxTracker
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Transaction, state.ZKCounters, string) (*TxTracker, error)); ok {
		return rf(tx, counters, ip)
	}
	if rf, ok := ret.Get(0).(func(types.Transaction, state.ZKCounters, string) *TxTracker); ok {
		r0 = rf(tx, counters, ip)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*TxTracker)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Transaction, state.ZKCounters, string) error); ok {
		r1 = rf(tx, counters, ip)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAfterSingleSuccessfulTxExecution provides a mock function with given fields: from, touchedAddresses
func (_m *WorkerMock) UpdateAfterSingleSuccessfulTxExecution(from common.Address, touchedAddresses map[common.Address]*state.InfoReadWrite) []*TxTracker {
	ret := _m.Called(from, touchedAddresses)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAfterSingleSuccessfulTxExecution")
	}

	var r0 []*TxTracker
	if rf, ok := ret.Get(0).(func(common.Address, map[common.Address]*state.InfoReadWrite) []*TxTracker); ok {
		r0 = rf(from, touchedAddresses)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*TxTracker)
		}
	}

	return r0
}

// UpdateTxZKCounters provides a mock function with given fields: txHash, from, ZKCounters
func (_m *WorkerMock) UpdateTxZKCounters(txHash common.Hash, from common.Address, ZKCounters state.ZKCounters) {
	_m.Called(txHash, from, ZKCounters)
}

// NewWorkerMock creates a new instance of WorkerMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWorkerMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *WorkerMock {
	mock := &WorkerMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
