package entity

import "time"

type TransactionEntity struct {
	ID              uint64 `gorm:"primaryKey"`
	RefID           string
	BillingID       string
	MerchantID      uint64
	Amount          float64
	FeeAmount       float64
	TotalAmount     float64
	MDRPercent      float64
	MDRAmount       float64
	PaymentMethod   string
	Currency        string
	Type            string
	Issuer          string
	Aquirer         string
	CustomerMPAN    string
	TransactionDate time.Time
	SettlementDate  time.Time
	Status          string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
