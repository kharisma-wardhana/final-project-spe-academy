package usecase_qr

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	generalEntity "github.com/kharisma-wardhana/final-project-spe-academy/entity"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/helper"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/repository/mysql"
	mEntity "github.com/kharisma-wardhana/final-project-spe-academy/internal/repository/mysql/entity"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/repository/redis"
	rEntity "github.com/kharisma-wardhana/final-project-spe-academy/internal/repository/redis/entity"
	usecase_log "github.com/kharisma-wardhana/final-project-spe-academy/internal/usecase/log"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/usecase/qr/entity"
)

type QRUseCase struct {
	logUseCase   usecase_log.ILogUseCase
	qrRepo       redis.IQRRepository
	merchantRepo mysql.IMerchantRepository
}

func NewQRUseCase(logUseCase usecase_log.ILogUseCase, qrRepo redis.IQRRepository, merchantRepo mysql.IMerchantRepository) *QRUseCase {
	return &QRUseCase{
		logUseCase:   logUseCase,
		qrRepo:       qrRepo,
		merchantRepo: merchantRepo,
	}
}

type IQRUseCase interface {
	GenerateQR(ctx context.Context, request entity.QRRequest) (*entity.QRResponse, error)
	ValidateQR(ctx context.Context, billingID string) (bool, error)
}

func (u *QRUseCase) GenerateQR(ctx context.Context, request entity.QRRequest) (*entity.QRResponse, error) {
	// Implement QR code generation logic here
	funcName := "QRUseCase.GenerateQR"
	captureFieldError := generalEntity.CaptureFields{
		"payload": helper.ToString(request),
	}

	if request.Amount <= 0 || request.Expiration <= 0 {
		err := fmt.Errorf("invalid request parameters: %v", captureFieldError)
		u.logUseCase.Error("QRUseCase.GenerateQR", funcName, err, captureFieldError)
		return nil, err
	}

	// Generate Billing ID
	billingID := fmt.Sprintf("ST-%d", generateRandomID())

	// Generate QR code payload
	merchant, err := u.merchantRepo.FindByID(ctx, request.MerchantID)
	if err != nil {
		u.logUseCase.Error("merchantRepo.FindByID", funcName, err, captureFieldError)
		return nil, err
	}
	qrCode := generateQRISPayload(merchant, request)

	err = u.qrRepo.Create(ctx, &rEntity.QREntity{
		MerchantID: request.MerchantID,
		BillingID:  billingID,
		Amount:     request.Amount,
		QRCode:     qrCode,
		Expiration: request.Expiration,
	})
	if err != nil {
		u.logUseCase.Error("qrRepo.Create", funcName, err, captureFieldError)
		return nil, err
	}

	// This is a placeholder implementation
	return &entity.QRResponse{
		QRCode:     qrCode,
		BillingID:  billingID,
		Amount:     request.Amount,
		Expiration: request.Expiration,
	}, nil
}

func (u *QRUseCase) ValidateQR(ctx context.Context, billingID string) (bool, error) {
	// Implement QR code validation logic here
	// This is a placeholder implementation
	if billingID == "" {
		return false, nil
	}

	_, err := u.qrRepo.GetByBillingID(ctx, billingID)
	if err != nil {
		u.logUseCase.Error("qrRepo.GetByBillingID", "QRUseCase.ValidateQR", err, generalEntity.CaptureFields{"billingID": billingID})
		return false, err
	}

	return true, nil
}

func generateRandomID() int64 {
	// Generate a random ID for the billing ID
	return time.Now().UnixNano() + rand.Int63n(1000000)
}

func formatTag(tag, value string) string {
	return fmt.Sprintf("%s%02d%s", tag, len(value), value)
}

func generateQRISPayload(merchant *mEntity.MerchantEntity, request entity.QRRequest) string {
	var payload strings.Builder

	// Format standar QRIS static (contoh merchant statis tanpa acquirer spesifik)
	payload.WriteString(formatTag("00", "01"))                                // Payload Format Indicator
	payload.WriteString(formatTag("01", "11"))                                // Point of Initiation Method (11 = Static)
	payload.WriteString(formatTag("52", merchant.MCC))                        // Merchant Category Code
	payload.WriteString(formatTag("53", request.Currency))                    // Transaction Currency (360 = IDR)
	payload.WriteString(formatTag("54", fmt.Sprintf("%.2f", request.Amount))) // Transaction Amount (optional, here Rp10.00)
	payload.WriteString(formatTag("58", "ID"))                                // Country Code
	payload.WriteString(formatTag("59", merchant.Name))                       // Merchant Name
	payload.WriteString(formatTag("60", merchant.City))                       // Merchant City
	payload.WriteString(formatTag("61", "01"))                                // Transaction Type (01 = Payment)
	// Tag 62 = Additional Data, optional
	payload.WriteString(formatTag("62", fmt.Sprintf("01%02d%s", len(merchant.MID), merchant.MID))) // Merchant ID

	// Append CRC placeholder
	data := payload.String()
	crc := calculateCRC(data + "6304") // CRC untuk seluruh data + "6304" (Tag 63 + Length 4)
	final := data + "6304" + strings.ToUpper(crc)

	return final
}

// CRC-16/CCITT-FALSE calculation
func calculateCRC(input string) string {
	crc := uint16(0xFFFF)
	for _, b := range []byte(input) {
		crc ^= uint16(b) << 8
		for i := 0; i < 8; i++ {
			if crc&0x8000 != 0 {
				crc = (crc << 1) ^ 0x1021
			} else {
				crc <<= 1
			}
		}
	}
	return fmt.Sprintf("%04X", crc)
}
