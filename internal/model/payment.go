package model

// PaymentPayload is a struct to hold payment row data
type PaymentPayload struct {
	ID       string  `json: "id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Date     string  `json:"date"`
	Type     string  `json:"type"`
	Comment  string  `json:"comment"`
	Category string  `json:"category"`
}

//PaymentDB ...
type PaymentDB struct {
	ID       int64  `db: "id"`
	Category string `db:"Category"`
	Payment  string `db:"Payment"`
}

//NewPaymentDB ...
func NewPaymentDB() *PaymentDB {
	return &PaymentDB{}
}
