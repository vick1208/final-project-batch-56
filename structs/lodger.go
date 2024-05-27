package model

type Lodger struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	City      string `json:"city"`
	Phone     string `json:"phone"`
}
