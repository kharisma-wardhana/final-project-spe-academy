package entity

type QREntity struct {
	MerchantID uint64
	BillingID  string
	Amount     float64
	QRCode     string
	Expiration int64
}
