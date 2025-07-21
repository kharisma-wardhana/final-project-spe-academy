package entity

type MerchantRequest struct {
	Name          string `json:"name"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	AccountNumber string `json:"account_number"`
	MID           string `json:"mid"`
	NMID          string `json:"nmid"`
	MPAN          string `json:"mpan"`
	MCC           string `json:"mcc"`
	PostalCode    string `json:"postal_code"`
	Province      string `json:"province"`
	District      string `json:"district"`
	SubDistrict   string `json:"sub_district"`
	City          string `json:"city"`
	Status        string `json:"status"`
}

type MerchantResponse struct {
	ID            uint64 `json:"id"`
	Name          string `json:"name"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	AccountNumber string `json:"account_number"`
	MID           string `json:"mid"`
	NMID          string `json:"nmid"`
	MPAN          string `json:"mpan"`
	MCC           string `json:"mcc"`
	PostalCode    string `json:"postal_code"`
	Province      string `json:"province"`
	District      string `json:"district"`
	SubDistrict   string `json:"sub_district"`
	City          string `json:"city"`
	Status        string `json:"status"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}
