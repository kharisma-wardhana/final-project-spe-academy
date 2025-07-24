package usecase_qr

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	generalEntity "github.com/kharisma-wardhana/final-project-spe-academy/entity"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/helper"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/repository/redis"
	rEntity "github.com/kharisma-wardhana/final-project-spe-academy/internal/repository/redis/entity"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/usecase/qr/entity"
)

type QRUseCase struct {
	qrRepo *redis.QRRepository
}

func NewQRUseCase(qrRepo *redis.QRRepository) *QRUseCase {
	return &QRUseCase{qrRepo}
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
		helper.LogError("qrRepo.GenerateQR", funcName, err, captureFieldError, "")
		return nil, err
	}

	// Generate Billing ID
	billingID := fmt.Sprintf("ST-%d", generateRandomID())

	// Generate QR code payload
	qrCode := generateQRISPayload()

	err := u.qrRepo.Create(ctx, &rEntity.QREntity{
		MerchantID: request.MerchantID,
		BillingID:  billingID,
		Amount:     request.Amount,
		QRCode:     qrCode,
		Expiration: request.Expiration,
	})
	if err != nil {
		helper.LogError("qrRepo.SaveQR", funcName, err, captureFieldError, "")
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
		helper.LogError("qrRepo.GetByBillingID", "QRUseCase.ValidateQR", err, generalEntity.CaptureFields{"billingID": billingID}, "")
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

func generateQRISPayload() string {
	var payload strings.Builder

	// Format standar QRIS static (contoh merchant statis tanpa acquirer spesifik)
	payload.WriteString(formatTag("00", "01"))       // Payload Format Indicator
	payload.WriteString(formatTag("01", "11"))       // Point of Initiation Method (11 = Static)
	payload.WriteString(formatTag("52", "0000"))     // Merchant Category Code
	payload.WriteString(formatTag("53", "360"))      // Transaction Currency (360 = IDR)
	payload.WriteString(formatTag("54", "1000"))     // Transaction Amount (optional, here Rp10.00)
	payload.WriteString(formatTag("58", "ID"))       // Country Code
	payload.WriteString(formatTag("59", "TOKO ABC")) // Merchant Name
	payload.WriteString(formatTag("60", "JAKARTA"))  // Merchant City
	// Tag 62 = Additional Data, optional

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
