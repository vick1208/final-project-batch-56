package model

type Room struct {
	ID          int    `json:"id"`
	RoomNumber  string `json:"room_number"`
	BedQty      int    `json:"bed_qty"`
	Description string `json:"description"`
	Price       string `json:"price"`
	RoomStatus  string `json:"room_status"`
}
