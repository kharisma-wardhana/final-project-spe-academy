package mocks

import (
	context "context"

	"github.com/kharisma-wardhana/final-project-spe-academy/internal/repository/mysql"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/repository/mysql/entity"
	"github.com/stretchr/testify/mock"
)

type AccountRepository struct {
	mock.Mock
}

func (_m *AccountRepository) Begin() (mysql.TrxObj, error) {
	ret := _m.Called()

	var r0 mysql.TrxObj
	if rf, ok := ret.Get(0).(func() mysql.TrxObj); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mysql.TrxObj)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *AccountRepository) Create(ctx context.Context, dbTrx mysql.TrxObj, account *entity.AccountEntity) error {
	ret := _m.Called(ctx, dbTrx, account)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, mysql.TrxObj, *entity.AccountEntity) error); ok {
		r0 = rf(ctx, dbTrx, account)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *AccountRepository) FindByClientID(ctx context.Context, dbTrx mysql.TrxObj, clientID string) (*entity.AccountEntity, error) {
	ret := _m.Called(ctx, dbTrx, clientID)

	var r0 *entity.AccountEntity
	if rf, ok := ret.Get(0).(func(context.Context, mysql.TrxObj, string) *entity.AccountEntity); ok {
		r0 = rf(ctx, dbTrx, clientID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.AccountEntity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, mysql.TrxObj, string) error); ok {
		r1 = rf(ctx, dbTrx, clientID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *AccountRepository) FindByID(ctx context.Context, dbTrx mysql.TrxObj, id uint64) (*entity.AccountEntity, error) {
	ret := _m.Called(ctx, dbTrx, id)

	var r0 *entity.AccountEntity
	if rf, ok := ret.Get(0).(func(context.Context, mysql.TrxObj, uint64) *entity.AccountEntity); ok {
		r0 = rf(ctx, dbTrx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.AccountEntity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, mysql.TrxObj, uint64) error); ok {
		r1 = rf(ctx, dbTrx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAccountRepository interface {
	mock.TestingT
	Cleanup(func())
}

func NewAccountRepository(t mockConstructorTestingTNewAccountRepository) *AccountRepository {
	mock := &AccountRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
