package dto

type MeResponse struct {
	Id       uint
	Username string `json:"username"`
	Email    string `json:"email"`
	Points   int    `json:"points"`
}
