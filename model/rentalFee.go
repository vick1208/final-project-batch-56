package model

type RentalFee struct {
	ID            int    `json:"id"`
	ReservationID int    `json:"reservation_id"`
	Bank          string `json:"bank"`
	PaymentType   string `json:"payment_type"`
	AdditionalFee string `json:"additional_fee"`
	TotalFee      string `json:"total_fee"`
}
