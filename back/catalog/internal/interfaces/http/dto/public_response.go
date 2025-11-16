package dto

type PublicCarResponse struct {
    ID           uint    `json:"id"`
    Brand        string  `json:"brand"`
    Type         string  `json:"type"`
    Model        string  `json:"model"`
    Year         int     `json:"year"`
    Price        float64 `json:"price"`
    Color        string  `json:"color"`       
    Transmission string  `json:"transmission"` 
    Status       string  `json:"status"`       
}
type ListCarsParams struct {
	BrandID   *[]uint   `form:"brand_id"`   
	TypeID    *[]uint   `form:"type_id"`    
	YearFrom  *int    `form:"year_from"` 
	YearTo    *int    `form:"year_to"`    
	PriceFrom *float64 `form:"price_from"`
	PriceTo   *float64 `form:"price_to"`
}

type DeleteRequest struct {
	ID []uint `json:"id"`
}

type SaveRequest struct {
	Name string `json:"name"`
	IsVisible bool `json:"is_visible"`
}