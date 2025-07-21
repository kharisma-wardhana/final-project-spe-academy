package usecase_merchant

import (
	"context"

	"github.com/kharisma-wardhana/spe-academy-learn-golang/final-project/internal/helper"
	"github.com/kharisma-wardhana/spe-academy-learn-golang/final-project/internal/repository/mysql"
	"github.com/kharisma-wardhana/spe-academy-learn-golang/final-project/internal/usecase/merchant/entity"
)

type IMerchantUsecase interface {
	CreateMerchant(ctx context.Context, req *entity.MerchantRequest) (*entity.MerchantResponse, error)
	GetAllMerchant() ([]*entity.MerchantResponse, error)
	GetMerchantByMID(mid string) (*entity.MerchantResponse, error)
}

type MerchantUsecase struct {
	merchantRepo mysql.IMerchantRepository
}

func NewMerchantUsecase(merchantRepo mysql.IMerchantRepository) *MerchantUsecase {
	return &MerchantUsecase{
		merchantRepo,
	}
}

func (u *MerchantUsecase) CreateMerchant(ctx context.Context, req *entity.MerchantRequest) (*entity.MerchantResponse, error) {
	merchant, err := u.merchantRepo.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return merchant, nil
}

func (u *MerchantUsecase) GetAllMerchant() ([]*entity.MerchantResponse, error) {
	merchants, err := u.merchantRepo.GetAllMerchant()
	if err != nil {
		return nil, err
	}

	return merchants, nil
}

func (u *MerchantUsecase) GetMerchantByMID(mid string) (*entity.MerchantResponse, error) {
	merchant, err := u.merchantRepo.GetMerchantByMID(mid)
	if err != nil {
		return nil, err
	}

	return &entity.MerchantResponse{
		Name:          merchant.Name,
		Phone:         merchant.Phone,
		Email:         merchant.Email,
		MID:           merchant.MID,
		NMID:          merchant.NMID,
		MPAN:          merchant.MPAN,
		MCC:           merchant.MCC,
		AccountNumber: merchant.AccountNumber,
		PostalCode:    merchant.PostalCode,
		Province:      merchant.Province,
		District:      merchant.District,
		SubDistrict:   merchant.SubDistrict,
		City:          merchant.City,
		Status:        merchant.Status,
		CreatedAt:     helper.ConvertToJakartaDate(merchant.CreatedAt),
		UpdatedAt:     helper.ConvertToJakartaDate(merchant.UpdatedAt),
	}, nil
}
