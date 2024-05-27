package model

type Rental struct {
	ID            int    `json:"id"`
	LodgerID      int    `json:"lodger_id"`
	RoomID        int    `json:"room_id"`
	Bank          string `json:"bank"`
	PaymentType   string `json:"payment_type"`
	AdditionalFee string `json:"additional_fee"`
	TotalFee      string `json:"total_fee"`
}
