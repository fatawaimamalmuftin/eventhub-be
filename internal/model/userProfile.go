package model

type UserProfile struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
	Location string `json:"location"`
	Profile  string `json:"profile"`
	Job      string `json:"job"`
}
