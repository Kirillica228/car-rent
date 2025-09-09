package postgres

import (
	"catalog/internal/entity"

	"gorm.io/gorm"
)

type postgresCarRepo struct {
	db *gorm.DB
}

func (p *postgresCarRepo) Delete(id uint) error {
	return p.db.Delete(&CarModel{}, id).Error
}

func (p *postgresCarRepo) GetByID(id uint) (entity.Car, error) {
	var car CarModel
	if err := p.db.Preload("Brand").Preload("Type").First(&car, id).Error; err != nil {
		return entity.Car{}, err
	}
	return toEntity(car), nil
}

func (p *postgresCarRepo) List() ([]entity.Car, error) {
	var cars []CarModel

	if err := p.db.
    Preload("Brand").
    Preload("Type").
    Where("is_visible = ?", true).
    Find(&cars).Error; err != nil {
    return nil, err
	}

	return toEntityList(cars), nil
}

func (p *postgresCarRepo) Save(car entity.Car) error {
    model := toModel(car)

    // ⚡ отключаем пересоздание брендов и типов
    if err := p.db.Omit("Brand").Omit("Type").Create(&model).Error; err != nil {
        return err
    }
    return nil
}

func (p *postgresCarRepo) Update(id uint, car entity.Car) error {
    updated := toModel(car)
    updated.ID = id

    if err := p.db.Omit("Brand").Omit("Type").Save(&updated).Error; err != nil {
        return err
    }
    return nil
}

func (p *postgresCarRepo) ListAdmin() ([]entity.Car, error) {
	var cars []CarModel

	if err := p.db.
    Preload("Brand").
    Preload("Type").
    Find(&cars).Error; err != nil {
    return nil, err
	}

	return toEntityList(cars), nil
}


func NewPostgresCarRepo(db *gorm.DB) entity.CarRepository {
	return &postgresCarRepo{db: db}
}
