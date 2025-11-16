package dto

import "time"

type RentReponse struct {
	ID         uint      `json:"id"`
	Car		   CarReponse `json:"car"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	TotalPrice float64   `json:"total_price"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}

type CarReponse struct {
	ID           uint `json:"id"`
	Brand        string `json:"brand"`
	Type         string  `json:"type"`
	Model        string `json:"model"`
	Image        string `json:"image"`
}