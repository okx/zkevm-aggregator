// Code generated by mockery v2.39.0. DO NOT EDIT.

package ethtxmanager

import (
	context "context"
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	time "time"

	types "github.com/ethereum/go-ethereum/core/types"
)

// ethermanMock is an autogenerated mock type for the ethermanInterface type
type ethermanMock struct {
	mock.Mock
}

// CheckTxWasMined provides a mock function with given fields: ctx, txHash
func (_m *ethermanMock) CheckTxWasMined(ctx context.Context, txHash common.Hash) (bool, *types.Receipt, error) {
	ret := _m.Called(ctx, txHash)

	if len(ret) == 0 {
		panic("no return value specified for CheckTxWasMined")
	}

	var r0 bool
	var r1 *types.Receipt
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) (bool, *types.Receipt, error)); ok {
		return rf(ctx, txHash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) bool); ok {
		r0 = rf(ctx, txHash)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) *types.Receipt); ok {
		r1 = rf(ctx, txHash)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*types.Receipt)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, common.Hash) error); ok {
		r2 = rf(ctx, txHash)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CurrentNonce provides a mock function with given fields: ctx, account
func (_m *ethermanMock) CurrentNonce(ctx context.Context, account common.Address) (uint64, error) {
	ret := _m.Called(ctx, account)

	if len(ret) == 0 {
		panic("no return value specified for CurrentNonce")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) (uint64, error)); ok {
		return rf(ctx, account)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) uint64); ok {
		r0 = rf(ctx, account)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address) error); ok {
		r1 = rf(ctx, account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EstimateGas provides a mock function with given fields: ctx, from, to, value, data
func (_m *ethermanMock) EstimateGas(ctx context.Context, from common.Address, to *common.Address, value *big.Int, data []byte) (uint64, error) {
	ret := _m.Called(ctx, from, to, value, data)

	if len(ret) == 0 {
		panic("no return value specified for EstimateGas")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *common.Address, *big.Int, []byte) (uint64, error)); ok {
		return rf(ctx, from, to, value, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *common.Address, *big.Int, []byte) uint64); ok {
		r0 = rf(ctx, from, to, value, data)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address, *common.Address, *big.Int, []byte) error); ok {
		r1 = rf(ctx, from, to, value, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRevertMessage provides a mock function with given fields: ctx, tx
func (_m *ethermanMock) GetRevertMessage(ctx context.Context, tx *types.Transaction) (string, error) {
	ret := _m.Called(ctx, tx)

	if len(ret) == 0 {
		panic("no return value specified for GetRevertMessage")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.Transaction) (string, error)); ok {
		return rf(ctx, tx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.Transaction) string); ok {
		r0 = rf(ctx, tx)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.Transaction) error); ok {
		r1 = rf(ctx, tx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTx provides a mock function with given fields: ctx, txHash
func (_m *ethermanMock) GetTx(ctx context.Context, txHash common.Hash) (*types.Transaction, bool, error) {
	ret := _m.Called(ctx, txHash)

	if len(ret) == 0 {
		panic("no return value specified for GetTx")
	}

	var r0 *types.Transaction
	var r1 bool
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) (*types.Transaction, bool, error)); ok {
		return rf(ctx, txHash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *types.Transaction); ok {
		r0 = rf(ctx, txHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) bool); ok {
		r1 = rf(ctx, txHash)
	} else {
		r1 = ret.Get(1).(bool)
	}

	if rf, ok := ret.Get(2).(func(context.Context, common.Hash) error); ok {
		r2 = rf(ctx, txHash)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetTxReceipt provides a mock function with given fields: ctx, txHash
func (_m *ethermanMock) GetTxReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	ret := _m.Called(ctx, txHash)

	if len(ret) == 0 {
		panic("no return value specified for GetTxReceipt")
	}

	var r0 *types.Receipt
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) (*types.Receipt, error)); ok {
		return rf(ctx, txHash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *types.Receipt); ok {
		r0 = rf(ctx, txHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Receipt)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, txHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendTx provides a mock function with given fields: ctx, tx
func (_m *ethermanMock) SendTx(ctx context.Context, tx *types.Transaction) error {
	ret := _m.Called(ctx, tx)

	if len(ret) == 0 {
		panic("no return value specified for SendTx")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.Transaction) error); ok {
		r0 = rf(ctx, tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SignTx provides a mock function with given fields: ctx, sender, tx
func (_m *ethermanMock) SignTx(ctx context.Context, sender common.Address, tx *types.Transaction) (*types.Transaction, error) {
	ret := _m.Called(ctx, sender, tx)

	if len(ret) == 0 {
		panic("no return value specified for SignTx")
	}

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *types.Transaction) (*types.Transaction, error)); ok {
		return rf(ctx, sender, tx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *types.Transaction) *types.Transaction); ok {
		r0 = rf(ctx, sender, tx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address, *types.Transaction) error); ok {
		r1 = rf(ctx, sender, tx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SuggestedGasPrice provides a mock function with given fields: ctx
func (_m *ethermanMock) SuggestedGasPrice(ctx context.Context) (*big.Int, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for SuggestedGasPrice")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*big.Int, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *big.Int); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WaitTxToBeMined provides a mock function with given fields: ctx, tx, timeout
func (_m *ethermanMock) WaitTxToBeMined(ctx context.Context, tx *types.Transaction, timeout time.Duration) (bool, error) {
	ret := _m.Called(ctx, tx, timeout)

	if len(ret) == 0 {
		panic("no return value specified for WaitTxToBeMined")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.Transaction, time.Duration) (bool, error)); ok {
		return rf(ctx, tx, timeout)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.Transaction, time.Duration) bool); ok {
		r0 = rf(ctx, tx, timeout)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.Transaction, time.Duration) error); ok {
		r1 = rf(ctx, tx, timeout)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// newEthermanMock creates a new instance of ethermanMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newEthermanMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *ethermanMock {
	mock := &ethermanMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
