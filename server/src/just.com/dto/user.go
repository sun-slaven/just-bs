package dto

type UserUpdateRequest struct {
	Name    string
	Email   string
	IconUrl string `json:"icon_url"`
}