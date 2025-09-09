package dto

type CreateCarRequest struct {
    BrandID       uint    `json:"brand_id"`
    TypeID        uint    `json:"type_id"`
    ModelCar      string  `json:"model"`
    Year          int     `json:"year"`
    LicensePlate  string  `json:"license_plate"`
    Status        string  `json:"status"`
    Price         float32 `json:"price"`
	IsVisible     bool    `json:"is_visible"`
}