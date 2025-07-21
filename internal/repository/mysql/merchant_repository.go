package mysql

import (
	"github.com/kharisma-wardhana/spe-academy-learn-golang/final-project/config"
	"github.com/kharisma-wardhana/spe-academy-learn-golang/final-project/internal/repository/mysql/entity"
)

type IMerchantRepository interface {
	TrxSupportRepo
	GetMerchantByID(id uint64) (*entity.MerchantEntity, error)
	GetMerchantByMID(mid string) (*entity.MerchantEntity, error)
}

type MerchantRepository struct {
	GormTrxObj
}

func NewMerchantRepository(mysql *config.Mysql) *MerchantRepository {
	return &MerchantRepository{GormTrxObj{db: mysql.DB}}
}

func (r *MerchantRepository) GetMerchantByID(id uint64) (*entity.MerchantEntity, error) {
	var merchant entity.MerchantEntity
	if err := r.db.First(&merchant, id).Error; err != nil {
		return nil, err
	}
	return &merchant, nil
}

func (r *MerchantRepository) GetMerchantByMID(mid string) (*entity.MerchantEntity, error) {
	var merchant entity.MerchantEntity
	if err := r.db.Where("mid = ?", mid).First(&merchant).Error; err != nil {
		return nil, err
	}
	return &merchant, nil
}
