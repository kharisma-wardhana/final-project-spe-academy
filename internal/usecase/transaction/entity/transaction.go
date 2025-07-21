package entity

type TransactionRequest struct {
	RefID         string  `json:"reference_id"`
	BillingID     string  `json:"billing_id"`
	MerchantID    uint64  `json:"merchant_id"`
	Amount        float64 `json:"amount"`
	FeeAmount     float64 `json:"fee_amount"`
	TotalAmount   float64 `json:"total_amount"`
	MDRPercent    float64 `json:"mdr_percent"`
	MDRAmount     float64 `json:"mdr_amount"`
	PaymentMethod string  `json:"payment_method"`
	Currency      string  `json:"currency"`
	Type          string  `json:"type"`
	CustomerMPAN  string  `json:"customer_mpan"`
	Status        string  `json:"status"`
}

type TransactionResponse struct {
	ID              uint64  `json:"id"`
	RefID           string  `json:"reference_id"`
	BillingID       string  `json:"billing_id"`
	MerchantID      uint64  `json:"merchant_id"`
	Amount          float64 `json:"amount"`
	FeeAmount       float64 `json:"fee_amount"`
	TotalAmount     float64 `json:"total_amount"`
	MDRPercent      float64 `json:"mdr_percent"`
	MDRAmount       float64 `json:"mdr_amount"`
	PaymentMethod   string  `json:"payment_method"`
	Currency        string  `json:"currency"`
	Type            string  `json:"type"`
	Issuer          string  `json:"issuer"`
	Aquirer         string  `json:"aquirer"`
	CustomerMPAN    string  `json:"customer_mpan"`
	TransactionDate string  `json:"transaction_date"`
	SettlementDate  string  `json:"settlement_date"`
	Status          string  `json:"status"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
}
