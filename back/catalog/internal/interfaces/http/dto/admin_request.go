package dto

type CreateCarRequest struct {
    BrandID      uint      `json:"brand_id"`
    TypeID       uint      `json:"type_id"`
    Model        string    `json:"model"`
    Year         int       `json:"year"`
    LicensePlate string    `json:"license_plate"`
    Status       string    `json:"status"`      
    Price        float64   `json:"price"`
    IsVisible    bool      `json:"is_visible"`
    Images       []string  `json:"images"`
    Color        string    `json:"color"`       
    Description  string    `json:"description"`  
    Transmission string    `json:"transmission"` 
}