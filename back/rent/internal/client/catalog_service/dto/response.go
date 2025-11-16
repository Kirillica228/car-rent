// internal/rent/client/cars/dto/car.go
package dto

type Car struct {
	ID           uint   `json:"id"`
	Model        string `json:"model"`
	Status       string `json:"status"`
	Price        float64 `json:"price"`
	Brand        CarBrand `json:"brand"`
	Type         CarType  `json:"type"`
	Image        []Image `json:"image"`
}

type CarType struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type CarBrand struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Image struct {
	ID    uint		
	URL   string
}
