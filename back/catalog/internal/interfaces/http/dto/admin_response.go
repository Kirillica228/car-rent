package dto

type CarResponseAdmin struct {
	ID            uint   `json:"id"`
	Brand         string `json:"brand"`
	Type          string `json:"type"`
	Model         string `json:"model"`
	Year          int    `json:"year"`
	License_plate string `json:"license_plate"`
	Status        string `json:"status"`
	Price         float32 `json:"price"`
	IsVisible     bool   `json:"is_visible"`
}