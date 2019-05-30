package objects

type V1TaxesObjectRequest struct {
	Name    string  `json:"name"`
	TaxCode int64   `json:"tax_code"`
	Price   float64 `json:"price"`
}

type V1TaxesObjectResponse struct {
	Name       string  `json:"name"`
	TaxCode    int64   `json:"tax_code"`
	Type       string  `json:"type"`
	Refundable string  `json:"refundable"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	Amount     float64 `json:"amount"`
}
