package postgres

import (
	"time"

	"gorm.io/gorm"
)

type CarModel struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	BrandID       uint
	Brand         CarBrandModel
	TypeID        uint
	Type          CarTypeModel
	ModelCar      string
	Year          int
	License_plate string
	Status        string
	Price         float32
	IsVisible bool `gorm:"default:false"`
}

type CarBrandModel struct {
	ID    uint   `gorm:"primaryKey"`
	Name string
	IsVisible bool `gorm:"default:false"`
}

type CarTypeModel struct {
	ID   uint   `gorm:"primaryKey"`
	Name string
	IsVisible bool `gorm:"default:false"`
}
