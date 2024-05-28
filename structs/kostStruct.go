package structs

type Lodger struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	City      string `json:"city"`
	Phone     string `json:"phone"`
}
