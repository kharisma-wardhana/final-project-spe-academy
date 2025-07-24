package redis

import (
	"context"
	"encoding/json"

	generalEntity "github.com/kharisma-wardhana/final-project-spe-academy/entity"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/helper"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/repository/redis/entity"
	"github.com/redis/go-redis/v9"
)

type IQRRepository interface {
	Create(ctx context.Context, qr *entity.QREntity) error
	GetByBillingID(ctx context.Context, billingID string) (*entity.QREntity, error)
}

type QRRepository struct {
	redisClient *redis.Client
}

func NewQRRepository(redisClient *redis.Client) *QRRepository {
	return &QRRepository{redisClient}
}

func (r *QRRepository) Create(ctx context.Context, qr *entity.QREntity) error {
	funcName := "QRRepository.Create"
	captureFieldError := generalEntity.CaptureFields{
		"payload": helper.ToString(qr),
	}

	qrJSON, err := json.Marshal(qr)

	if err != nil {
		helper.LogError("json.Marshal", funcName, err, captureFieldError, "")
		return err
	}

	if err := r.redisClient.Set(ctx, qr.BillingID, qrJSON, 0).Err(); err != nil {
		helper.LogError("redisClient.Set", funcName, err, captureFieldError, "")
		return err
	}

	return nil
}

func (r *QRRepository) GetByBillingID(ctx context.Context, billingID string) (*entity.QREntity, error) {
	funcName := "QRRepository.GetByBillingID"
	captureFieldError := generalEntity.CaptureFields{
		"billingID": billingID,
	}

	data, err := r.redisClient.Get(ctx, billingID).Result()
	if err != nil {
		helper.LogError("redisClient.Get", funcName, err, captureFieldError, "")
		return nil, err
	}

	var qr entity.QREntity
	if err := json.Unmarshal([]byte(data), &qr); err != nil {
		helper.LogError("json.Unmarshal", funcName, err, captureFieldError, "")
		return nil, err
	}

	return &qr, nil
}
