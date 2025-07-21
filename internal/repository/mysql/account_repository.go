package mysql

import (
	"context"

	"github.com/kharisma-wardhana/spe-academy-learn-golang/final-project/config"
	"github.com/kharisma-wardhana/spe-academy-learn-golang/final-project/internal/repository/mysql/entity"
)

type IAccountRepository interface {
	TrxSupportRepo
	GetAccountByID(ctx context.Context, id uint64) (*entity.AccountEntity, error)
	GetAccountByMerchantID(ctx context.Context, merchantID uint64) (*entity.AccountEntity, error)
}

type AccountRepository struct {
	GormTrxObj
}

func NewAccountRepository(mysql *config.Mysql) *AccountRepository {
	return &AccountRepository{GormTrxObj{db: mysql.DB}}
}

func (r *AccountRepository) GetAccountByID(ctx context.Context, id uint64) (*entity.AccountEntity, error) {
	var account entity.AccountEntity
	if err := r.db.First(&account, id).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *AccountRepository) GetAccountByMerchantID(ctx context.Context, merchantID uint64) (*entity.AccountEntity, error) {
	var account entity.AccountEntity
	if err := r.db.Where("merchant_id = ?", merchantID).First(&account).Error; err != nil {
		return nil, err
	}
	return &account, nil
}
