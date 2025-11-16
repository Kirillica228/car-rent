package entity

import (

	"time"
)

type RentStatus string

const (
	StatusActive    RentStatus = "Новый"
	StatusCompleted RentStatus = "Завершён"
	StatusCanceled  RentStatus = "Отменён"
)


type Rent struct {
	ID         uint
	UserID     uint
	CarID      uint
	Name       string
	Image      string
	StartDate  time.Time
	EndDate    time.Time
	TotalPrice float64
	Status     RentStatus
	CreatedAt  time.Time
}

type BookingInterval struct {
    StartDate time.Time `json:"start_date"`
    EndDate   time.Time `json:"end_date"`
}

type RentRepository interface {
	GetAllByUserID(id uint) ([]Rent, error)
	GetByID(id uint) (Rent, error)
	GetAll() ([]Rent, error)
	Save(rent Rent) error
	Update(rent Rent) error
	Delete(id uint) error
	ExistsOverlap(carID uint, startDate, endDate time.Time) (bool, error)
	CancelRent(rentID uint) (int, error)
	ListByCarID(carID uint) ([]Rent, error)
	GetBusyCars(start, end time.Time) ([]uint, error)
}
