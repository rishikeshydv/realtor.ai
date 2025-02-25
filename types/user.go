package types

type LoginType struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpType struct {
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}
