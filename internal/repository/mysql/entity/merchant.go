package entity

import "time"

type MerchantEntity struct {
	ID            uint64    `gorm:"primaryKey"`
	Name          string    `gorm:"column:name"`
	Phone         string    `gorm:"unique"`
	Email         string    `gorm:"unique"`
	AccountNumber string    `gorm:"column:account_number"`
	MID           string    `gorm:"column:mid"`
	NMID          string    `gorm:"column:nmid"`
	MPAN          string    `gorm:"column:mpan"`
	MCC           string    `gorm:"column:mcc"`
	PostalCode    string    `gorm:"column:postal_code"`
	Province      string    `gorm:"column:province"`
	District      string    `gorm:"column:district"`
	SubDistrict   string    `gorm:"column:subdistrict"`
	City          string    `gorm:"column:city"`
	Status        string    `gorm:"column:status"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}

func (MerchantEntity) TableName() string {
	return "merchants"
}
