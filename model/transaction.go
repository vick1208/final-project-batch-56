package model

type Transaction struct {
	ID            int    `json:"id"`
	ReservationID int    `json:"reservation_id"`
	Bank          string `json:"bank"`
	PaymentType   string `json:"payment_type"`
	AdditionalFee int    `json:"additional_fee"`
	TotalFee      int    `json:"total_fee"`
}
