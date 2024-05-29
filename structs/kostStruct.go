package structs

type Lodger struct {
	ID    int    `json:"id"`
	Name  string `json:"name" binding:"required"`
	City  string `json:"city" binding:"required"`
	Phone string `json:"phone" binding:"required"`
}

type Room struct {
	ID         int    `json:"id"`
	RoomNumber string `json:"room_number" binding:"required"`
	Price      int    `json:"price" binding:"required"`
	RoomStatus string `json:"room_status" binding:"required"`
}

type Reservation struct {
	ID        int    `json:"id"`
	LodgerID  int    `json:"lodger_id"`
	RoomID    int    `json:"room_id"`
	DateStart string `json:"date_start"`
}

type Transaction struct {
	ID            int    `json:"id"`
	RentalID      int    `json:"rental_id"`
	Bank          string `json:"bank"`
	PaymentDate   string `json:"payment_date"`
	DueDate       string `json:"due_date"`
	PaymentType   string `json:"payment_type"`
	MainFee       int    `json:"main_fee"`
	AdditionalFee int    `json:"additional_fee"`
}
