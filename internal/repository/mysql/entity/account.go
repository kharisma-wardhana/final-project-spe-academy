package entity

import "time"

type AccountEntity struct {
	ID           uint64 `gorm:"primaryKey"`
	MerchantID   uint64
	ClientID     string
	ClientSecret string
	PrivateKey   string
	PublicKey    string
	Status       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
