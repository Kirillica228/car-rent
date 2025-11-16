package repository

import (
	"catalog/internal/domain/entity"
	"catalog/internal/interfaces/http/dto"

	"gorm.io/gorm"
)

type postgresCarRepo struct {
	db *gorm.DB
}

func NewPostgresCarRepo(db *gorm.DB) entity.CarRepository {
	return &postgresCarRepo{db: db}
}

func (p *postgresCarRepo) Delete(ids []uint) error {
	return p.db.Unscoped().Delete(&entity.Car{}, ids).Error
}

func (p *postgresCarRepo) GetByID(id uint) (entity.Car, error) {
	var car entity.Car
	if err := p.db.Preload("Brand").Preload("Type").Preload("Images").First(&car, id).Error; err != nil {
		return entity.Car{}, err
	}
	return car, nil
}

func (p *postgresCarRepo) List(params dto.ListCarsParams, public bool) ([]entity.Car, error) {
	var cars []entity.Car

	db := p.db.Model(&entity.Car{}).
		Preload("Brand").
		Preload("Type").
		Preload("Images")

	if public {
		db = db.Where("is_visible = ?", true)
	}
	if params.BrandID != nil && len(*params.BrandID) > 0 {
    db = db.Where("brand_id IN ?", *params.BrandID)
	}

	if params.TypeID != nil && len(*params.TypeID) > 0 {
		db = db.Where("type_id IN ?", *params.TypeID)
	}
	if params.YearFrom != nil {
		db = db.Where("year >= ?", *params.YearFrom)
	}
	if params.YearTo != nil {
		db = db.Where("year <= ?", *params.YearTo)
	}
	if params.PriceFrom != nil {
		db = db.Where("price >= ?", *params.PriceFrom)
	}
	if params.PriceTo != nil {
		db = db.Where("price <= ?", *params.PriceTo)
	}

	// Выполняем запрос
	if err := db.Find(&cars).Error; err != nil {
		return nil, err
	}

	return cars, nil
}

func (p *postgresCarRepo) Save(car *entity.Car) error {
	if err := p.db.Omit("Brand", "Type", "Images").Create(car).Error; err != nil {
		return err
	}
	return nil
}

func (p *postgresCarRepo) Update(id uint, car entity.Car) error {
	updated := car
	updated.ID = id
	
	if err := p.db.Omit("Brand", "Type", "Images").Save(&updated).Error; err != nil {
		return err
	}
	return nil
}

func (r *postgresCarRepo) AddPhoto(carID uint, url string) error {
	photo := entity.Image{
		CarID: carID,
		URL:   url,
	}
	return r.db.Create(&photo).Error
}

func (r *postgresCarRepo) DeletePhoto(photoURL string) error {
	return r.db.Where("url = ?", photoURL).Delete(&entity.Image{}).Error
}
