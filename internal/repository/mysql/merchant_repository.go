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

type IMerchantRepository interface {
	TrxSupportRepo
	FindByID(ctx context.Context, id int64) (*entity.MerchantEntity, error)
	FindByMID(ctx context.Context, mid string) (*entity.MerchantEntity, error)
	LockByID(ctx context.Context, dbTrx TrxObj, id int64) (*entity.MerchantEntity, error)
	Create(ctx context.Context, dbTrx TrxObj, params *entity.MerchantEntity, nonZeroVal bool) error
	Update(ctx context.Context, dbTrx TrxObj, params *entity.MerchantEntity, changes *entity.MerchantEntity) error
	DeleteByID(ctx context.Context, dbTrx TrxObj, id int64) error
}

type MerchantRepository struct {
	GormTrxSupport
}

func NewMerchantRepository(mysql *config.Mysql) *MerchantRepository {
	return &MerchantRepository{GormTrxSupport{db: mysql.DB}}
}

func (r *MerchantRepository) FindByID(ctx context.Context, id int64) (*entity.MerchantEntity, error) {
	funcName := "MerchantRepository.FindByID"
	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errwrap.Wrap(err, funcName)
	}

	var merchant entity.MerchantEntity
	if err := r.db.
		Raw("SELECT * FROM merchants WHERE id = ?", id).
		Scan(&merchant).
		Error; err != nil {
		return nil, err
	}
	return &merchant, nil
}

func (r *MerchantRepository) FindByMID(ctx context.Context, mid string) (*entity.MerchantEntity, error) {
	funcName := "MerchantRepository.FindByMID"
	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errwrap.Wrap(err, funcName)
	}

	var merchant entity.MerchantEntity
	if err := r.db.
		Raw("SELECT * FROM merchants WHERE mid = ?", mid).
		Scan(&merchant).
		Error; err != nil {
		return nil, err
	}
	return &merchant, nil
}

func (r *MerchantRepository) LockByID(ctx context.Context, dbTrx TrxObj, id int64) (*entity.MerchantEntity, error) {
	funcName := "MerchantRepository.LockByID"
	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errwrap.Wrap(err, funcName)
	}

	var merchant entity.MerchantEntity
	err := r.Trx(dbTrx).
		Raw("SELECT * FROM merchants WHERE id = ? FOR UPDATE", id).
		Scan(&merchant).Error

	if errwrap.Is(err, gorm.ErrRecordNotFound) {
		return nil, appErr.ErrRecordNotFound()
	} else if err != nil {
		return nil, errwrap.Wrap(err, funcName)
	}

	return &merchant, nil
}

func (r *MerchantRepository) Create(ctx context.Context, dbTrx TrxObj, params *entity.MerchantEntity, nonZeroVal bool) error {
	funcName := "MerchantRepository.Create"
	if err := helper.CheckDeadline(ctx); err != nil {
		return errwrap.Wrap(err, funcName)
	}

	cols := helper.NonZeroCols(params, nonZeroVal)
	return r.Trx(dbTrx).Select(cols).Create(&params).Error
}

func (r *MerchantRepository) Update(ctx context.Context, dbTrx TrxObj, params *entity.MerchantEntity, changes *entity.MerchantEntity) error {
	funcName := "MerchantRepository.Update"
	if err := helper.CheckDeadline(ctx); err != nil {
		return errwrap.Wrap(err, funcName)
	}

	if err := r.db.WithContext(ctx).Model(params).Updates(changes).Error; err != nil {
		return err
	}
	return nil
}

func (r *MerchantRepository) DeleteByID(ctx context.Context, dbTrx TrxObj, id int64) error {
	funcName := "MerchantRepository.DeleteByID"
	if err := helper.CheckDeadline(ctx); err != nil {
		return errwrap.Wrap(err, funcName)
	}

	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.MerchantEntity{}).Error; err != nil {
		return err
	}
	return nil
}
