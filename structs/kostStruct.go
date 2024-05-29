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
