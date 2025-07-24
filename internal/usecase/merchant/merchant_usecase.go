package usecase_merchant

import (
	"context"
	"fmt"
	"time"

	generalEntity "github.com/kharisma-wardhana/final-project-spe-academy/entity"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/helper"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/repository/mysql"
	mEntity "github.com/kharisma-wardhana/final-project-spe-academy/internal/repository/mysql/entity"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/usecase"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/usecase/merchant/entity"
	errWrap "github.com/pkg/errors"
)

type MerchantUseCase struct {
	merchantRepo mysql.IMerchantRepository
}

func NewMerchantUseCase(merchantRepo mysql.IMerchantRepository) *MerchantUseCase {
	return &MerchantUseCase{
		merchantRepo,
	}
}

type IMerchantUseCase interface {
	CreateMerchant(ctx context.Context, req *entity.MerchantRequest) (*entity.MerchantResponse, error)
	UpdateMerchant(ctx context.Context, id int64, req *entity.MerchantRequest) (*entity.MerchantResponse, error)
	GetMerchantByMID(ctx context.Context, mid string) (*entity.MerchantResponse, error)
	DeleteMerchantByID(ctx context.Context, id int64) error
}

func (u *MerchantUseCase) CreateMerchant(ctx context.Context, req *entity.MerchantRequest) (*entity.MerchantResponse, error) {
	funcName := "MerchantUseCase.CreateMerchant"
	captureFieldError := generalEntity.CaptureFields{
		"payload": helper.ToString(req),
	}
	if err := usecase.ValidateStruct(*req); err != "" {
		return nil, errWrap.Wrap(fmt.Errorf(generalEntity.INVALID_PAYLOAD_CODE), err)
	}
	merchant := &mEntity.MerchantEntity{
		Name:          req.Name,
		Phone:         req.Phone,
		Email:         req.Email,
		MID:           req.MID,
		NMID:          req.NMID,
		MPAN:          req.MPAN,
		MCC:           req.MCC,
		AccountNumber: req.AccountNumber,
		PostalCode:    req.PostalCode,
		Province:      req.Province,
		District:      req.District,
		SubDistrict:   req.SubDistrict,
		City:          req.City,
		Status:        req.Status,
	}
	if err := u.merchantRepo.Create(ctx, nil, merchant, true); err != nil {
		helper.LogError("merchantRepo.Create", funcName, err, captureFieldError, "")
		return nil, err
	}

	return &entity.MerchantResponse{
		ID:            merchant.ID,
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

func (u *MerchantUseCase) GetMerchantByMID(ctx context.Context, mid string) (*entity.MerchantResponse, error) {
	funcName := "MerchantUseCase.GetMerchantByMID"
	captureFieldError := generalEntity.CaptureFields{"mid": mid}

	merchant, err := u.merchantRepo.FindByMID(ctx, mid)
	if err != nil {
		helper.LogError("merchantRepo.FindByMID", funcName, err, captureFieldError, "")
		return nil, err
	}

	return &entity.MerchantResponse{
		ID:            merchant.ID,
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

func (u *MerchantUseCase) UpdateMerchant(ctx context.Context, id int64, req *entity.MerchantRequest) (result *entity.MerchantResponse, err error) {
	funcName := "MerchantUseCase.UpdateMerchant"
	captureFieldError := generalEntity.CaptureFields{
		"id":      helper.ToString(id),
		"payload": helper.ToString(req),
	}
	if err := usecase.ValidateStruct(*req); err != "" {
		return nil, errWrap.Wrap(fmt.Errorf(generalEntity.INVALID_PAYLOAD_CODE), err)
	}

	if err := mysql.DBTransaction(u.merchantRepo, func(dbTrx mysql.TrxObj) error {
		merchantEntity, err := u.merchantRepo.LockByID(ctx, dbTrx, id)
		if err != nil {
			helper.LogError("merchantRepo.LockByID", funcName, err, captureFieldError, "")
			return err
		}
		changes := &mEntity.MerchantEntity{
			Name:          req.Name,
			Phone:         req.Phone,
			Email:         req.Email,
			MID:           req.MID,
			NMID:          req.NMID,
			MPAN:          req.MPAN,
			MCC:           req.MCC,
			AccountNumber: req.AccountNumber,
			PostalCode:    req.PostalCode,
			Province:      req.Province,
			District:      req.District,
			SubDistrict:   req.SubDistrict,
			City:          req.City,
			Status:        req.Status,
			UpdatedAt:     time.Now(),
		}
		err = u.merchantRepo.Update(ctx, dbTrx, merchantEntity, changes)
		if err != nil {
			helper.LogError("merchantRepo.Update", funcName, err, captureFieldError, "")
			return err
		}
		result = &entity.MerchantResponse{
			ID:            merchantEntity.ID,
			Name:          merchantEntity.Name,
			Phone:         merchantEntity.Phone,
			Email:         merchantEntity.Email,
			MID:           merchantEntity.MID,
			NMID:          merchantEntity.NMID,
			MPAN:          merchantEntity.MPAN,
			MCC:           merchantEntity.MCC,
			AccountNumber: merchantEntity.AccountNumber,
			PostalCode:    merchantEntity.PostalCode,
			Province:      merchantEntity.Province,
			District:      merchantEntity.District,
			SubDistrict:   merchantEntity.SubDistrict,
			City:          merchantEntity.City,
			Status:        merchantEntity.Status,
			CreatedAt:     helper.ConvertToJakartaDate(merchantEntity.CreatedAt),
			UpdatedAt:     helper.ConvertToJakartaDate(merchantEntity.UpdatedAt),
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}

func (u *MerchantUseCase) DeleteMerchantByID(ctx context.Context, id int64) error {
	funcName := "MerchantUseCase.DeleteMerchantByID"
	captureFieldError := generalEntity.CaptureFields{"id": helper.ToString(id)}

	if err := u.merchantRepo.DeleteByID(ctx, nil, id); err != nil {
		helper.LogError("merchantRepo.DeleteByID", funcName, err, captureFieldError, "")
		return err
	}

	return nil
}
