package models

type User struct {
	ID          string
	Email       string
	FirstName   string
	LastName    string
	Address     string
	PhoneNumber string
	Password    string
}

type UserJson struct {
	Email       string `json:"email"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
