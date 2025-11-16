package dto

type CarResponseAdmin struct {
    ID           uint     `json:"id"`
    Brand        string   `json:"brand"`
    Type         string   `json:"type"`
    Model        string   `json:"model"`
    Year         int      `json:"year"`
    LicensePlate string   `json:"license_plate"`
    Status       string   `json:"status"`       
    Price        float64  `json:"price"`
    IsVisible    bool     `json:"is_visible"`
    Images       []string `json:"images"`
    Color        string   `json:"color"`        
    Description  string   `json:"description"` 
    Transmission string   `json:"transmission"` 
}
