package mysql

import (
	"github.com/kharisma-wardhana/spe-academy-learn-golang/final-project/config"
	"github.com/kharisma-wardhana/spe-academy-learn-golang/final-project/internal/repository/mysql/entity"
)

type ITransactionRepository interface {
	TrxSupportRepo
	GetTransactionByID(id uint64) (*entity.TransactionEntity, error)
	GetTransactionsByMerchantID(merchantID uint64) ([]entity.TransactionEntity, error)
	GetTransactionsByRefID(refID uint64) (*entity.TransactionEntity, error)
}

type TransactionRepository struct {
	GormTrxObj
}

func NewTransactionRepository(mysql *config.Mysql) *TransactionRepository {
	return &TransactionRepository{GormTrxObj{db: mysql.DB}}
}

func (r *TransactionRepository) GetTransactionByID(id uint64) (*entity.TransactionEntity, error) {
	var transaction entity.TransactionEntity
	if err := r.db.First(&transaction, id).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *TransactionRepository) GetTransactionsByMerchantID(merchantID uint64) ([]entity.TransactionEntity, error) {
	var transactions []entity.TransactionEntity
	if err := r.db.Where("merchant_id = ?", merchantID).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *TransactionRepository) GetTransactionsByRefID(refID string) (*entity.TransactionEntity, error) {
	var transaction entity.TransactionEntity
	if err := r.db.Where("ref_id = ?", refID).First(&transaction).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}
