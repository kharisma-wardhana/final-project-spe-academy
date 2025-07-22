package mysql

import (
	"context"

	"github.com/kharisma-wardhana/final-project-spe-academy/config"
	appErr "github.com/kharisma-wardhana/final-project-spe-academy/error"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/helper"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/repository/mysql/entity"
	errwrap "github.com/pkg/errors"
	"gorm.io/gorm"
)

type IAccountRepository interface {
	TrxSupportRepo
	LockByID(ctx context.Context, dbTrx TrxObj, id int64) (*entity.AccountEntity, error)
	FindByID(ctx context.Context, id int64) (*entity.AccountEntity, error)
	FindByMerchantID(ctx context.Context, merchantID int64) (*entity.AccountEntity, error)
	Create(ctx context.Context, dbTrx TrxObj, params *entity.AccountEntity, nonZeroVal bool) error
	Update(ctx context.Context, dbTrx TrxObj, params *entity.AccountEntity, changes *entity.AccountEntity) (err error)
	DeleteByID(ctx context.Context, dbTrx TrxObj, id int64) error
}

type AccountRepository struct {
	GormTrxSupport
}

func NewAccountRepository(mysql *config.Mysql) *AccountRepository {
	return &AccountRepository{GormTrxSupport{db: mysql.DB}}
}

func (r *AccountRepository) FindByID(ctx context.Context, id int64) (*entity.AccountEntity, error) {
	funcName := "AccountRepository.FindByID"
	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errwrap.Wrap(err, funcName)
	}

	var account entity.AccountEntity
	if err := r.db.
		Raw("SELECT * FROM accounts WHERE id = ?", id).
		First(&account).
		Error; err != nil {
		if errwrap.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErr.ErrRecordNotFound()
		}
		return nil, err
	}
	return &account, nil
}

func (r *AccountRepository) FindByMerchantID(ctx context.Context, id int64) (*entity.AccountEntity, error) {
	funcName := "AccountRepository.FindByMerchantID"
	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errwrap.Wrap(err, funcName)
	}

	var account entity.AccountEntity
	if err := r.db.
		Raw("SELECT * FROM merchants WHERE account_id = ?", id).
		Scan(&account).
		Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *AccountRepository) Create(ctx context.Context, dbTrx TrxObj, params *entity.AccountEntity, nonZeroVal bool) error {
	funcName := "AccountRepository.Create"
	if err := helper.CheckDeadline(ctx); err != nil {
		return errwrap.Wrap(err, funcName)
	}

	cols := helper.NonZeroCols(params, nonZeroVal)
	return r.Trx(dbTrx).Select(cols).Create(&params).Error
}

func (r *AccountRepository) Update(ctx context.Context, dbTrx TrxObj, params *entity.AccountEntity, changes *entity.AccountEntity) (err error) {
	funcName := "AccountRepository.Update"
	if err := helper.CheckDeadline(ctx); err != nil {
		return errwrap.Wrap(err, funcName)
	}

	db := r.Trx(dbTrx).Model(params)
	if changes != nil {
		err = db.Updates(*changes).Error
	} else {
		err = db.Updates(helper.StructToMap(params, false)).Error
	}

	if err != nil {
		return errwrap.Wrap(err, funcName)
	}

	return nil
}

func (r *AccountRepository) DeleteByID(ctx context.Context, dbTrx TrxObj, id int64) error {
	funcName := "AccountRepository.DeleteByID"
	if err := helper.CheckDeadline(ctx); err != nil {
		return errwrap.Wrap(err, funcName)
	}

	err := r.Trx(dbTrx).Where("id = ?", id).Delete(&entity.AccountEntity{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return appErr.ErrRecordNotFound()
		}
	}
	return nil
}

func (r *AccountRepository) LockByID(ctx context.Context, dbTrx TrxObj, id int64) (*entity.AccountEntity, error) {
	funcName := "AccountRepository.LockByID"
	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errwrap.Wrap(err, funcName)
	}

	var account entity.AccountEntity
	err := r.Trx(dbTrx).
		Raw("SELECT * FROM accounts WHERE id = ? FOR UPDATE", id).
		First(&account).Error

	if errwrap.Is(err, gorm.ErrRecordNotFound) {
		return nil, appErr.ErrRecordNotFound()
	} else if err != nil {
		return nil, errwrap.Wrap(err, funcName)
	}

	return &account, nil
}
