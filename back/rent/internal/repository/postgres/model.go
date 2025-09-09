package repository

import (
	"time"
)

type RentModel struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	UserID    uint      `gorm:"not null;index"`
	CarID     uint      `gorm:"not null;index"`
	StartTime time.Time `gorm:"not null"`
	EndTime   time.Time
	Status    string    `gorm:"type:varchar(20);default:'ACTIVE'"` // ACTIVE, COMPLETED, CANCELED
	CreatedAt time.Time
	UpdatedAt time.Time
}
