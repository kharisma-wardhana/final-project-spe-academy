package usecase_account

import (
	"context"
	"fmt"
	"time"

	generalEntity "github.com/kharisma-wardhana/final-project-spe-academy/entity"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/helper"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/repository/mysql"
	mEntity "github.com/kharisma-wardhana/final-project-spe-academy/internal/repository/mysql/entity"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/usecase"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/usecase/account/entity"
	errWrap "github.com/pkg/errors"
)

type AccountUseCase struct {
	// Add any dependencies needed for the use case here
	accountRepo mysql.IAccountRepository
}

func NewAccountUseCase(accountRepo mysql.IAccountRepository) *AccountUseCase {
	return &AccountUseCase{accountRepo}
}

type IAccountUseCase interface {
	GetAccountByID(ctx context.Context, id int64) (*entity.AccountResponse, error)
	GetAccountByMerchantID(ctx context.Context, merchantID int64) (*entity.AccountResponse, error)
	CreateAccount(ctx context.Context, req *entity.AccountRequest) (*entity.AccountResponse, error)
	UpdateAccount(ctx context.Context, id int64, req *entity.AccountRequest) (result *entity.AccountResponse, err error)
	DeleteAccount(ctx context.Context, id int64) error
}

func (u *AccountUseCase) GetAccountByID(ctx context.Context, id int64) (*entity.AccountResponse, error) {
	funcName := "AccountUseCase.GetAccountByID"
	captureFieldError := generalEntity.CaptureFields{"id": helper.ToString(id)}

	account, err := u.accountRepo.FindByID(ctx, id)
	if err != nil {
		helper.LogError("accountRepo.FindByID", funcName, err, captureFieldError, "")
		return nil, err
	}

	return &entity.AccountResponse{
		ID:           account.ID,
		MerchantID:   account.MerchantID,
		ClientID:     account.ClientID,
		ClientSecret: account.ClientSecret,
		PrivateKey:   account.PrivateKey,
		PublicKey:    account.PublicKey,
		Status:       account.Status,
		CreatedAt:    helper.ConvertToJakartaDate(account.CreatedAt),
		UpdatedAt:    helper.ConvertToJakartaDate(account.UpdatedAt),
	}, nil
}

func (u *AccountUseCase) GetAccountByMerchantID(ctx context.Context, merchantID int64) (*entity.AccountResponse, error) {
	funcName := "AccountUseCase.GetAccountByMerchantID"
	captureFieldError := generalEntity.CaptureFields{
		"merchantID": helper.ToString(merchantID),
	}

	account, err := u.accountRepo.FindByMerchantID(ctx, merchantID)
	if err != nil {
		helper.LogError("accountRepo.FindByMerchantID", funcName, err, captureFieldError, "")
		return nil, err
	}

	return &entity.AccountResponse{
		ID:           account.ID,
		MerchantID:   account.MerchantID,
		ClientID:     account.ClientID,
		ClientSecret: account.ClientSecret,
		PrivateKey:   account.PrivateKey,
		PublicKey:    account.PublicKey,
		Status:       account.Status,
		CreatedAt:    helper.ConvertToJakartaDate(account.CreatedAt),
		UpdatedAt:    helper.ConvertToJakartaDate(account.UpdatedAt),
	}, nil
}

func (u *AccountUseCase) CreateAccount(ctx context.Context, req *entity.AccountRequest) (*entity.AccountResponse, error) {
	funcName := "AccountUseCase.CreateAccount"
	captureFieldError := generalEntity.CaptureFields{
		"payload": helper.ToString(req),
	}
	if err := usecase.ValidateStruct(*req); err != "" {
		return nil, errWrap.Wrap(fmt.Errorf(generalEntity.INVALID_PAYLOAD_CODE), err)
	}
	var accountEntity = &mEntity.AccountEntity{
		MerchantID:   req.MerchantID,
		ClientID:     req.ClientID,
		ClientSecret: req.ClientSecret,
		PrivateKey:   req.PrivateKey,
		PublicKey:    req.PublicKey,
		Status:       req.Status,
	}

	err := u.accountRepo.Create(ctx, nil, accountEntity, true)
	if err != nil {
		helper.LogError("accountRepo.Create", funcName, err, captureFieldError, "")
		return nil, err
	}

	return &entity.AccountResponse{
		ID:           accountEntity.ID,
		MerchantID:   accountEntity.MerchantID,
		ClientID:     accountEntity.ClientID,
		ClientSecret: accountEntity.ClientSecret,
		PrivateKey:   accountEntity.PrivateKey,
		PublicKey:    accountEntity.PublicKey,
		Status:       accountEntity.Status,
		CreatedAt:    helper.ConvertToJakartaDate(accountEntity.CreatedAt),
		UpdatedAt:    helper.ConvertToJakartaDate(accountEntity.UpdatedAt),
	}, nil
}

func (u *AccountUseCase) UpdateAccount(ctx context.Context, id int64, req *entity.AccountRequest) (result *entity.AccountResponse, err error) {
	funcName := "AccountUseCase.UpdateAccount"
	captureFieldError := generalEntity.CaptureFields{
		"payload": helper.ToString(req),
	}
	if err := usecase.ValidateStruct(*req); err != "" {
		return nil, errWrap.Wrap(fmt.Errorf(generalEntity.INVALID_PAYLOAD_CODE), err)
	}

	if err := mysql.DBTransaction(u.accountRepo, func(dbTrx mysql.TrxObj) error {
		accountEntity, err := u.accountRepo.LockByID(ctx, dbTrx, id)
		if err != nil {
			helper.LogError("accountRepo.LockByID", funcName, err, captureFieldError, "")
			return err
		}
		if accountEntity == nil {
			return fmt.Errorf("account with id %d not found", id)
		}

		// Process the changes
		changes := &mEntity.AccountEntity{
			MerchantID:   req.MerchantID,
			ClientID:     req.ClientID,
			ClientSecret: req.ClientSecret,
			PrivateKey:   req.PrivateKey,
			PublicKey:    req.PublicKey,
			Status:       req.Status,
			UpdatedAt:    time.Now(),
		}
		if err := u.accountRepo.Update(ctx, dbTrx, accountEntity, changes); err != nil {
			helper.LogError("accountRepo.Update", funcName, err, captureFieldError, "")
			return err
		}
		result = &entity.AccountResponse{
			ID:           accountEntity.ID,
			MerchantID:   accountEntity.MerchantID,
			ClientID:     accountEntity.ClientID,
			ClientSecret: accountEntity.ClientSecret,
			Status:       accountEntity.Status,
			CreatedAt:    helper.ConvertToJakartaDate(accountEntity.CreatedAt),
			UpdatedAt:    helper.ConvertToJakartaDate(accountEntity.UpdatedAt),
		}
		return nil
	}); err != nil {
		helper.LogError("accountRepo.DBTransaction", funcName, err, captureFieldError, "")
		return nil, errWrap.Wrap(err, funcName)
	}

	return result, nil
}

func (u *AccountUseCase) DeleteAccount(ctx context.Context, id int64) error {
	funcName := "AccountUseCase.DeleteAccount"
	captureFieldError := generalEntity.CaptureFields{
		"id": helper.ToString(id),
	}

	if err := u.accountRepo.DeleteByID(ctx, nil, id); err != nil {
		helper.LogError("accountRepo.DeleteByID", funcName, err, captureFieldError, "")
		return err
	}
	return nil
}
