package structs

type Lodger struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	City  string `json:"city"`
	Phone string `json:"phone"`
}

type Room struct {
	ID         int    `json:"id"`
	RoomNumber string `json:"room_number"`
	Price      int    `json:"price"`
	RoomStatus string `json:"room_status"`
}

type Reservation struct {
	ID        int    `json:"id"`
	LodgerID  int    `json:"lodger_id"`
	RoomID    int    `json:"room_id"`
	DateStart string `json:"date_start"`
}

type Transaction struct {
	ID            int    `json:"id"`
	ReservationID int    `json:"reservation_id"`
	Bank          string `json:"bank"`
	PaymentDate   string `json:"payment_date"`
	DueDate       string `json:"due_date"`
	PaymentType   string `json:"payment_type"`
	MainFee       int    `json:"main_fee"`
	AdditionalFee int    `json:"additional_fee"`
}
