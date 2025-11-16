package dto

import "time"

type CreateRentDTO struct {
    CarID     uint      `json:"car_id" binding:"required"`
    StartDate time.Time `json:"start_date" binding:"required"` 
    EndDate   time.Time `json:"end_date" binding:"required"`  
}

type GetBusyCarsRequest struct {
    StartDate time.Time `json:"start_date"`
    EndDate   time.Time `json:"end_date"`
}

type BusyCarResponse struct {
    CarID uint `json:"car_id"`
}