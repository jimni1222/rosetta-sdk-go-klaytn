// Code generated by mockery v1.0.0. DO NOT EDIT.

package reconciler

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	database "github.com/klaytn/rosetta-sdk-go-klaytn/storage/database"
	types "github.com/klaytn/rosetta-sdk-go-klaytn/types"
)

// Helper is an autogenerated mock type for the Helper type
type Helper struct {
	mock.Mock
}

// CanonicalBlock provides a mock function with given fields: ctx, dbTx, block
func (_m *Helper) CanonicalBlock(ctx context.Context, dbTx database.Transaction, block *types.BlockIdentifier) (bool, error) {
	ret := _m.Called(ctx, dbTx, block)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, database.Transaction, *types.BlockIdentifier) bool); ok {
		r0 = rf(ctx, dbTx, block)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, database.Transaction, *types.BlockIdentifier) error); ok {
		r1 = rf(ctx, dbTx, block)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ComputedBalance provides a mock function with given fields: ctx, dbTx, account, currency, index
func (_m *Helper) ComputedBalance(ctx context.Context, dbTx database.Transaction, account *types.AccountIdentifier, currency *types.Currency, index int64) (*types.Amount, error) {
	ret := _m.Called(ctx, dbTx, account, currency, index)

	var r0 *types.Amount
	if rf, ok := ret.Get(0).(func(context.Context, database.Transaction, *types.AccountIdentifier, *types.Currency, int64) *types.Amount); ok {
		r0 = rf(ctx, dbTx, account, currency, index)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Amount)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, database.Transaction, *types.AccountIdentifier, *types.Currency, int64) error); ok {
		r1 = rf(ctx, dbTx, account, currency, index)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CurrentBlock provides a mock function with given fields: ctx, dbTx
func (_m *Helper) CurrentBlock(ctx context.Context, dbTx database.Transaction) (*types.BlockIdentifier, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 *types.BlockIdentifier
	if rf, ok := ret.Get(0).(func(context.Context, database.Transaction) *types.BlockIdentifier); ok {
		r0 = rf(ctx, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.BlockIdentifier)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, database.Transaction) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DatabaseTransaction provides a mock function with given fields: ctx
func (_m *Helper) DatabaseTransaction(ctx context.Context) database.Transaction {
	ret := _m.Called(ctx)

	var r0 database.Transaction
	if rf, ok := ret.Get(0).(func(context.Context) database.Transaction); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(database.Transaction)
		}
	}

	return r0
}

// ForceInactiveReconciliation provides a mock function with given fields: ctx, account, currency, lastCheck
func (_m *Helper) ForceInactiveReconciliation(ctx context.Context, account *types.AccountIdentifier, currency *types.Currency, lastCheck *types.BlockIdentifier) bool {
	ret := _m.Called(ctx, account, currency, lastCheck)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, *types.AccountIdentifier, *types.Currency, *types.BlockIdentifier) bool); ok {
		r0 = rf(ctx, account, currency, lastCheck)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IndexAtTip provides a mock function with given fields: ctx, index
func (_m *Helper) IndexAtTip(ctx context.Context, index int64) (bool, error) {
	ret := _m.Called(ctx, index)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, int64) bool); ok {
		r0 = rf(ctx, index)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, index)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LiveBalance provides a mock function with given fields: ctx, account, currency, index
func (_m *Helper) LiveBalance(ctx context.Context, account *types.AccountIdentifier, currency *types.Currency, index int64) (*types.Amount, *types.BlockIdentifier, error) {
	ret := _m.Called(ctx, account, currency, index)

	var r0 *types.Amount
	if rf, ok := ret.Get(0).(func(context.Context, *types.AccountIdentifier, *types.Currency, int64) *types.Amount); ok {
		r0 = rf(ctx, account, currency, index)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Amount)
		}
	}

	var r1 *types.BlockIdentifier
	if rf, ok := ret.Get(1).(func(context.Context, *types.AccountIdentifier, *types.Currency, int64) *types.BlockIdentifier); ok {
		r1 = rf(ctx, account, currency, index)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*types.BlockIdentifier)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *types.AccountIdentifier, *types.Currency, int64) error); ok {
		r2 = rf(ctx, account, currency, index)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PruneBalances provides a mock function with given fields: ctx, account, currency, index
func (_m *Helper) PruneBalances(ctx context.Context, account *types.AccountIdentifier, currency *types.Currency, index int64) error {
	ret := _m.Called(ctx, account, currency, index)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.AccountIdentifier, *types.Currency, int64) error); ok {
		r0 = rf(ctx, account, currency, index)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
