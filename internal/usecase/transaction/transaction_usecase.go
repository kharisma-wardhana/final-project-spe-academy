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
	usecase_log "github.com/kharisma-wardhana/final-project-spe-academy/internal/usecase/log"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/usecase/transaction/entity"
	errWrap "github.com/pkg/errors"
)

type TransactionUseCase struct {
	logUseCase      usecase_log.ILogUseCase
	transactionRepo mysql.ITransactionRepository
	qrRepo          redis.IQRRepository
}

func NewTransactionUseCase(logUseCase usecase_log.ILogUseCase, transactionRepo mysql.ITransactionRepository, qrRepo redis.IQRRepository) *TransactionUseCase {
	return &TransactionUseCase{
		logUseCase:      logUseCase,
		transactionRepo: transactionRepo,
		qrRepo:          qrRepo,
	}
}

type ITransactionUseCase interface {
	CreateTransaction(ctx context.Context, req *entity.TransactionRequest) (*entity.TransactionResponse, error)
	GetTransactionsByMerchantID(ctx context.Context, merchantID uint64) ([]*entity.TransactionResponse, error)
	GetTransactionsByRefID(ctx context.Context, refID string) (*entity.TransactionResponse, error)
}

func (u *TransactionUseCase) CreateTransaction(ctx context.Context, req *entity.TransactionRequest) (*entity.TransactionResponse, error) {
	funcName := "TransactionUseCase.CreateTransaction"
	captureFieldError := generalEntity.CaptureFields{
		"payload": helper.ToString(req),
	}

	if err := usecase.ValidateStruct(*req); err != "" {
		u.logUseCase.Error("usecase.ValidateStruct", funcName, fmt.Errorf("%s", err), captureFieldError)
		return nil, errWrap.Wrap(fmt.Errorf(generalEntity.INVALID_PAYLOAD_CODE), err)
	}

	if req.Amount <= 0 || req.FeeAmount < 0 || req.TotalAmount <= 0 {
		err := fmt.Errorf("invalid request parameters: %v", captureFieldError)
		u.logUseCase.Error("transactionRepo.CreateTransaction", funcName, err, captureFieldError)
		return nil, err
	}

	qr, err := u.qrRepo.GetByBillingID(ctx, req.BillingID)
	if err != nil {
		u.logUseCase.Error("qrRepo.GetByBillingID", funcName, err, captureFieldError)
		return nil, err
	}

	if qr == nil {
		err := fmt.Errorf("QR code not found for billing ID: %s", req.BillingID)
		u.logUseCase.Error("qrRepo.GetByBillingID", funcName, err, captureFieldError)
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

func (u *TransactionUseCase) GetTransactionsByMerchantID(ctx context.Context, merchantID uint64) ([]*entity.TransactionResponse, error) {
	funcName := "TransactionUseCase.GetTransactionsByMerchantID"
	captureFieldError := generalEntity.CaptureFields{"merchantID": helper.ToString(merchantID)}

	transactions, err := u.transactionRepo.FindByMerchantID(ctx, merchantID)
	if err != nil {
		u.logUseCase.Error("transactionRepo.FindByMerchantID", funcName, err, captureFieldError)
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
		u.logUseCase.Error("transactionRepo.FindByRefID", funcName, err, captureFieldError)
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
