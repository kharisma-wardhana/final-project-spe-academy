package entity

import "time"

type MerchantEntity struct {
	ID            uint64    `gorm:"primaryKey"`
	Name          string    `column:"name"`
	Phone         string    `gorm:"unique"`
	Email         string    `gorm:"unique"`
	AccountNumber string    `column:"account_number"`
	MID           string    `column:"mid"`
	NMID          string    `column:"nmid"`
	MPAN          string    `column:"mpan"`
	MCC           string    `column:"mcc"`
	PostalCode    string    `column:"postal_code"`
	Province      string    `column:"province"`
	District      string    `column:"district"`
	SubDistrict   string    `column:"sub_district"`
	City          string    `column:"city"`
	Status        string    `column:"status"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}
