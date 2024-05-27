package model

type Reservation struct {
	ID        int    `json:"id"`
	LodgerID  int    `json:"lodger_id"`
	RoomID    int    `json:"room_id"`
	DateStart string `json:"date_start"`
}
