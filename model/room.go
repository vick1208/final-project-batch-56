package model

type Room struct {
	ID           int    `json:"id"`
	RoomNumber   string `json:"room_number"`
	BedQty       string `json:"bed_qty"`
	Description  string `json:"description"`
	Price        string `json:"price"`
	RoomStatusID int    `json:"room_status_id"`
}