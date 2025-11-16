package entity

import (
	"catalog/internal/interfaces/http/dto"
	"mime/multipart"
	"time"

	"gorm.io/gorm"
)

type CarStatus string
type Transmission string

// Коробка передач
const (
	Automatic Transmission = "АКПП"
	Manual    Transmission = "МКПП"
)

// Статусы автомобиля (готовность к аренде)
const (
	InCheck   CarStatus = "На диагностике"
	Ready     CarStatus = "Готова к аренде"
	Unavailable CarStatus = "Недоступна"
)


type Car struct {
	ID           uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Images       []Image        `gorm:"constraint:OnDelete:CASCADE" json:"images"`
	BrandID      uint           `json:"brand_id"`
	Brand        Brand          `json:"brand"`
	TypeID       uint           `json:"type_id"`
	Type         Type           `json:"type"`
	Model        string         `gorm:"type:varchar(100);not null" json:"model"`
	Year         int            `json:"year"`
	Price        float64        `gorm:"type:numeric(10,2)" json:"price"`
	Color        string         `gorm:"type:varchar(50)" json:"color"`
	Description  string         `gorm:"type:text" json:"description"`
	Transmission Transmission   `gorm:"type:varchar(10)" json:"transmission"`
	LicensePlate string         `gorm:"type:varchar(20);unique;not null" json:"license_plate"`
	Status       CarStatus      `gorm:"type:varchar(20);default:'Готова к аренде'" json:"status"`
	IsVisible    bool           `gorm:"default:false" json:"is_visible"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}


type CarRepository interface {
	List(params dto.ListCarsParams, public bool) ([]Car, error)
	GetByID(id uint) (Car, error)
	Save(car *Car) error
	Delete(ids []uint) error
	Update(id uint, car Car) error
	AddPhoto(carID uint, url string) error
	DeletePhoto(photoURL string) error
}

type CarUseCase interface {
	GetByID(id uint) (Car, error)
	List(params dto.ListCarsParams, public bool) ([]Car, error)
	Create(car *Car) error
	Update(id uint, car Car) error 
	Delete(ids []uint) error
	UploadPhoto(carID uint, file multipart.File, header *multipart.FileHeader) (string, error)
	DeletePhotos(photoURL []string) error
}

type FileRepository interface {
	Save(file multipart.File, header *multipart.FileHeader, carID uint) (string, error)
	Delete(fileName string) error
}
