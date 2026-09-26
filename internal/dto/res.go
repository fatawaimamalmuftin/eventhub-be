package dto

type Res struct {
	Status  bool
	Message string
	Data    any `json:",omitempty"`
}
