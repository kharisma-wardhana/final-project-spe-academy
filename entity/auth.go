package entity

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	jwt.RegisteredClaims
	AccountID  uint64 `json:"account_id"`
	MerchantID uint64 `json:"merchant_id"`
}
