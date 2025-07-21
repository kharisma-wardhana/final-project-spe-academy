package entity

type AccountRequest struct {
	MerchantID   uint64 `json:"merchant_id"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	PrivateKey   string `json:"private_key"`
	PublicKey    string `json:"public_key"`
	Status       string `json:"status"`
}

type AccountResponse struct {
	ID           uint64 `json:"id"`
	MerchantID   uint64 `json:"merchant_id"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	PrivateKey   string `json:"private_key"`
	PublicKey    string `json:"public_key"`
	Status       string `json:"status"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
