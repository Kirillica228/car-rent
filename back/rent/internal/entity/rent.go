package entity

import "time"

type Rent struct {
	ID        uint
	UserID    uint
	CarID     uint
	StartTime time.Time
	EndTime   time.Time
	Price     float64
	Status    string
	CreatedAt time.Time
}


type RentRepository interface {
	GetAllByUserID(id uint) ([]Rent, error)
	GetByID(id uint) (*Rent, error)
	GetAll() ([]Rent, error)
	Save(rent *Rent) error
	Update(rent *Rent) error
	Delete(rent *Rent) error
}