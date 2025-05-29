package models

type AuthRequestDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
