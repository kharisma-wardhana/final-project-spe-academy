package entity

type QRRequest struct {
	MerchantID uint64  `json:"merchant_id"`
	Amount     float64 `json:"amount"`
	Currency   string  `json:"currency"`
	Expiration int64   `json:"expiration"` // in seconds
}

type QRResponse struct {
	QRCode     string  `json:"qr_code"`
	BillingID  string  `json:"billing_id"`
	Amount     float64 `json:"amount"`
	Expiration int64   `json:"expiration"` // in seconds
}
