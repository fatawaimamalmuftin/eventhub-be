package dto

type Account struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ChangeUserProfile struct {
	Bio      string `json:"bio"`
	Location string `json:"location"`
	Job      string `json:"job"`
	Profile  string `json:"profile"`
}
