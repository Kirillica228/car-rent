package repository

import (
	"rent/internal/entity"
	"time"
)


type Rent struct {
	ID         uint       `gorm:"primaryKey;autoIncrement"`
	UserID     uint       `gorm:"not null;index"`
	CarID      uint       `gorm:"not null;index"`
	Name       string     `gorm:"not null"`
	Image      string     `gorm:"not null"`
	StartDate  time.Time  `gorm:"not null"`
	EndDate    time.Time  `gorm:"not null"`
	Status     entity.RentStatus `gorm:"type:varchar(20);default:'Новый';check:status IN ('Новый','Завершён','Отменён')"`
	TotalPrice float64    `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
