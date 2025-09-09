package repository

import (
	"rent/internal/entity"
	"gorm.io/gorm"
)

type RentRepository struct {
	db *gorm.DB
}

func NewRentRepository(db *gorm.DB) *RentRepository {
	return &RentRepository{db: db}
}

func (r *RentRepository) Save(rent entity.Rent) (entity.Rent, error) {
	dbRent := ToRentDB(rent)
	if err := r.db.Create(&dbRent).Error; err != nil {
		return entity.Rent{}, err
	}
	return ToRentEntity(dbRent), nil
}

func (r *RentRepository) GetByID(id uint) (entity.Rent, error) {
	var dbRent RentModel
	if err := r.db.First(&dbRent, id).Error; err != nil {
		return entity.Rent{}, err
	}
	return ToRentEntity(dbRent), nil
}

func (r *RentRepository) GetAll() ([]entity.Rent, error) {
	var dbRents []RentModel
	if err := r.db.Find(&dbRents).Error; err != nil {
		return nil, err
	}

	rents := make([]entity.Rent, len(dbRents))
	for i, rnt := range dbRents {
		rents[i] = ToRentEntity(rnt)
	}
	return rents, nil
}

func (r *RentRepository) Update(rent entity.Rent) (entity.Rent, error) {
	dbRent := ToRentDB(rent)
	if err := r.db.Save(&dbRent).Error; err != nil {
		return entity.Rent{}, err
	}
	return ToRentEntity(dbRent), nil
}

func (r *RentRepository) Delete(id uint) error {
	return r.db.Delete(&RentModel{}, id).Error
}
