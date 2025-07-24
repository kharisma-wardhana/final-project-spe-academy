package usecase_transaction

import (
	"context"
	"fmt"
	"time"

	generalEntity "github.com/kharisma-wardhana/final-project-spe-academy/entity"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/helper"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/repository/mysql"
	mEntity "github.com/kharisma-wardhana/final-project-spe-academy/internal/repository/mysql/entity"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/repository/redis"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/usecase"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/usecase/transaction/entity"
	errWrap "github.com/pkg/errors"
)

type TransactionUseCase struct {
	transactionRepo mysql.ITransactionRepository
	qrRepo          redis.IQRRepository
}

func NewTransactionUseCase(transactionRepo mysql.ITransactionRepository, qrRepo redis.IQRRepository) *TransactionUseCase {
	return &TransactionUseCase{transactionRepo, qrRepo}
}

type ITransactionUseCase interface {
	CreateTransaction(ctx context.Context, req *entity.TransactionRequest) (*entity.TransactionResponse, error)
	GetTransactionsByMerchantID(ctx context.Context, merchantID int64) ([]*entity.TransactionResponse, error)
	GetTransactionsByRefID(ctx context.Context, refID string) (*entity.TransactionResponse, error)
}

func (u *TransactionUseCase) CreateTransaction(ctx context.Context, req *entity.TransactionRequest) (*entity.TransactionResponse, error) {
	funcName := "TransactionUseCase.CreateTransaction"
	captureFieldError := generalEntity.CaptureFields{
		"payload": helper.ToString(req),
	}

	if err := usecase.ValidateStruct(*req); err != "" {
		return nil, errWrap.Wrap(fmt.Errorf(generalEntity.INVALID_PAYLOAD_CODE), err)
	}

	if req.Amount <= 0 || req.FeeAmount < 0 || req.TotalAmount <= 0 {
		err := fmt.Errorf("invalid request parameters: %v", captureFieldError)
		helper.LogError("transactionRepo.CreateTransaction", funcName, err, captureFieldError, "")
		return nil, err
	}

	qr, err := u.qrRepo.GetByBillingID(ctx, req.BillingID)
	if err != nil {
		helper.LogError("qrRepo.GetByBillingID", funcName, err, captureFieldError, "")
		return nil, err
	}

	if qr == nil {
		err := fmt.Errorf("QR code not found for billing ID: %s", req.BillingID)
		helper.LogError("qrRepo.GetByBillingID", funcName, err, captureFieldError, "")
		return nil, err
	}

	transaction := &mEntity.TransactionEntity{
		RefID:         req.RefID,
		BillingID:     req.BillingID,
		MerchantID:    req.MerchantID,
		Amount:        req.Amount,
		FeeAmount:     req.FeeAmount,
		TotalAmount:   req.TotalAmount,
		MDRAmount:     req.MDRAmount,
		MDRPercent:    req.MDRPercent,
		PaymentMethod: req.PaymentMethod,
		Currency:      req.Currency,
		Type:          req.Type,
		CustomerMPAN:  req.CustomerMPAN,
		// Issuer:          req.Issuer,
		// Acquirer:        req.Acquirer,
		TransactionDate: time.Now(),
		Status:          req.Status,
	}

	if err := u.transactionRepo.Create(ctx, nil, transaction, true); err != nil {
		helper.LogError("transactionRepo.Create", funcName, err, captureFieldError, "")
		return nil, err
	}

	return &entity.TransactionResponse{
		ID:              transaction.ID,
		MerchantID:      transaction.MerchantID,
		RefID:           transaction.RefID,
		BillingID:       transaction.BillingID,
		Type:            transaction.Type,
		Amount:          transaction.Amount,
		PaymentMethod:   transaction.PaymentMethod,
		TotalAmount:     transaction.TotalAmount,
		TransactionDate: helper.ConvertToJakartaDate(transaction.TransactionDate),
		SettlementDate:  helper.ConvertToJakartaDate(transaction.SettlementDate),
		Status:          transaction.Status,
		CreatedAt:       helper.ConvertToJakartaDate(transaction.CreatedAt),
		UpdatedAt:       helper.ConvertToJakartaDate(transaction.UpdatedAt),
	}, nil
}

func (u *TransactionUseCase) GetTransactionsByMerchantID(ctx context.Context, merchantID int64) ([]*entity.TransactionResponse, error) {
	funcName := "TransactionUseCase.GetTransactionsByMerchantID"
	captureFieldError := generalEntity.CaptureFields{"merchantID": helper.ToString(merchantID)}

	transactions, err := u.transactionRepo.FindByMerchantID(ctx, merchantID)
	if err != nil {
		helper.LogError("transactionRepo.FindByMerchantID", funcName, err, captureFieldError, "")
		return nil, err
	}

	var response []*entity.TransactionResponse
	for _, transaction := range transactions {
		response = append(response, &entity.TransactionResponse{
			ID:              transaction.ID,
			MerchantID:      transaction.MerchantID,
			RefID:           transaction.RefID,
			BillingID:       transaction.BillingID,
			Type:            transaction.Type,
			Amount:          transaction.Amount,
			PaymentMethod:   transaction.PaymentMethod,
			TotalAmount:     transaction.TotalAmount,
			TransactionDate: helper.ConvertToJakartaDate(transaction.TransactionDate),
			SettlementDate:  helper.ConvertToJakartaDate(transaction.SettlementDate),
			Status:          transaction.Status,
			CreatedAt:       helper.ConvertToJakartaDate(transaction.CreatedAt),
			UpdatedAt:       helper.ConvertToJakartaDate(transaction.UpdatedAt),
		})
	}

	return response, nil
}

func (u *TransactionUseCase) GetTransactionsByRefID(ctx context.Context, refID string) (*entity.TransactionResponse, error) {
	funcName := "TransactionUseCase.GetTransactionsByRefID"
	captureFieldError := generalEntity.CaptureFields{"refID": refID}

	transaction, err := u.transactionRepo.FindByRefID(ctx, refID)
	if err != nil {
		helper.LogError("transactionRepo.FindByRefID", funcName, err, captureFieldError, "")
		return nil, err
	}

	return &entity.TransactionResponse{
		ID:              transaction.ID,
		MerchantID:      transaction.MerchantID,
		RefID:           transaction.RefID,
		BillingID:       transaction.BillingID,
		Type:            transaction.Type,
		Amount:          transaction.Amount,
		PaymentMethod:   transaction.PaymentMethod,
		TotalAmount:     transaction.TotalAmount,
		TransactionDate: helper.ConvertToJakartaDate(transaction.TransactionDate),
		SettlementDate:  helper.ConvertToJakartaDate(transaction.SettlementDate),
		Status:          transaction.Status,
		CreatedAt:       helper.ConvertToJakartaDate(transaction.CreatedAt),
		UpdatedAt:       helper.ConvertToJakartaDate(transaction.UpdatedAt),
	}, nil
}
