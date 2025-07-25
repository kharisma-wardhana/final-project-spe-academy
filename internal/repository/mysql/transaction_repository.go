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

type ITransactionRepository interface {
	TrxSupportRepo
	FindByID(ctx context.Context, id uint64) (*entity.TransactionEntity, error)
	FindByRefID(ctx context.Context, refID string) (*entity.TransactionEntity, error)
	FindByMerchantID(ctx context.Context, merchantID uint64) ([]entity.TransactionEntity, error)
	LockByID(ctx context.Context, dbTrx TrxObj, id uint64) (*entity.TransactionEntity, error)
	Create(ctx context.Context, dbTrx TrxObj, params *entity.TransactionEntity, nonZeroVal bool) error
}

type TransactionRepository struct {
	GormTrxSupport
}

func NewTransactionRepository(mysql *config.Mysql) *TransactionRepository {
	return &TransactionRepository{GormTrxSupport{db: mysql.DB}}
}

func (r *TransactionRepository) FindByID(ctx context.Context, id uint64) (*entity.TransactionEntity, error) {
	funcName := "TransactionRepository.FindByID"
	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errwrap.Wrap(err, funcName)
	}
	var transaction entity.TransactionEntity
	if err := r.db.
		Raw("SELECT * FROM transactions WHERE id = ?", id).
		Scan(&transaction).
		Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *TransactionRepository) FindByRefID(ctx context.Context, refID string) (*entity.TransactionEntity, error) {
	funcName := "TransactionRepository.FindByRefID"
	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errwrap.Wrap(err, funcName)
	}
	var transaction entity.TransactionEntity
	if err := r.db.
		Raw("SELECT * FROM transactions WHERE reference_id = ?", refID).
		Scan(&transaction).
		Error; err != nil {
		return nil, err
	}
	if transaction.ID == 0 {
		return nil, appErr.ErrRecordNotFound()
	}
	return &transaction, nil
}

func (r *TransactionRepository) FindByMerchantID(ctx context.Context, merchantID uint64) ([]entity.TransactionEntity, error) {
	funcName := "TransactionRepository.FindByMerchantID"
	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errwrap.Wrap(err, funcName)
	}
	var transactions []entity.TransactionEntity
	if err := r.db.
		Raw("SELECT * FROM transactions WHERE merchant_id = ?", merchantID).
		Scan(&transactions).
		Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *TransactionRepository) LockByID(ctx context.Context, dbTrx TrxObj, id uint64) (*entity.TransactionEntity, error) {
	funcName := "TransactionRepository.LockByID"
	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errwrap.Wrap(err, funcName)
	}
	var transaction entity.TransactionEntity
	err := r.Trx(dbTrx).
		Raw("SELECT * FROM transactions WHERE id = ? FOR UPDATE", id).
		Scan(&transaction).Error

	if errwrap.Is(err, gorm.ErrRecordNotFound) {
		return nil, appErr.ErrRecordNotFound()
	} else if err != nil {
		return nil, errwrap.Wrap(err, funcName)
	}

	return &transaction, nil
}

func (r *TransactionRepository) Create(ctx context.Context, dbTrx TrxObj, params *entity.TransactionEntity, nonZeroVal bool) error {
	funcName := "TransactionRepository.Create"
	if err := helper.CheckDeadline(ctx); err != nil {
		return errwrap.Wrap(err, funcName)
	}

	cols := helper.NonZeroCols(params, nonZeroVal)
	return r.Trx(dbTrx).Select(cols).Create(&params).Error
}
