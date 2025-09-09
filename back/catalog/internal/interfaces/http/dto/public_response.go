package dto

type PublicCarResponse struct {
	ID    uint   `json:"id"`
	Brand string `json:"brand"`
	Type  string `json:"type"`
	Model string `json:"model"`
	Year  int    `json:"year"`
	Price float32 `json:"price"`
}
