package models

type User struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	PhoneNumber string  `json:"phoneNumber"`
	Email       string  `json:"email"`
	Address     string  `json:"address"`
	Password    string  `json:"password"`
	CreatedAt   int     `json:"createdAt"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}
