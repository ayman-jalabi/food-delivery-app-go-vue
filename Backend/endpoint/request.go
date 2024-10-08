package endpoint

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegistrationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
