package models

type Contact struct {
	ID       uint    `json:"id"`
	Email    string  `json:"email"`
	Phone    string  `json:"phone"`
	Street   *string `json:"street"`
	Suburb   *string `json:"suburb"`
	Postcode *uint8  `json:"postcode"`
	State    *string `json:"state"`
	Country  *string `json:"country"`
}
