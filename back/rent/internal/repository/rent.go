package repository

import (
	"errors"
	"fmt"
	"net/http"
	"rent/internal/entity"
	"time"

	"gorm.io/gorm"
)

type RentRepository struct {
	db *gorm.DB
}

func NewRentRepository(db *gorm.DB) entity.RentRepository {
	return &RentRepository{db: db}
}

func (r *RentRepository) Save(rent entity.Rent) error {
	dbRent := ToRentDB(rent)
	if err := r.db.Create(&dbRent).Error; err != nil {
		return err
	}
	return nil
}

func (r *RentRepository) GetByID(id uint) (entity.Rent, error) {
	var dbRent Rent
	if err := r.db.First(&dbRent, id).Error; err != nil {
		return entity.Rent{}, err
	}
	return ToRentEntity(dbRent), nil
}

func (r *RentRepository) GetAll() ([]entity.Rent, error) {
	var dbRents []Rent
	if err := r.db.Find(&dbRents).Error; err != nil {
		return nil, err
	}

	rents := make([]entity.Rent, len(dbRents))
	for i, rnt := range dbRents {
		rents[i] = ToRentEntity(rnt)
	}
	return rents, nil
}

func (r *RentRepository) Update(rent entity.Rent) error {
	dbRent := ToRentDB(rent)
	if err := r.db.Save(&dbRent).Error; err != nil {
		return err
	}
	return nil
}

func (r *RentRepository) Delete(id uint) error {
	return r.db.Delete(&Rent{}, id).Error
}

func (r *RentRepository) GetAllByUserID(userID uint) ([]entity.Rent, error) {
	var dbRents []Rent
	if err := r.db.Where("user_id = ?", userID).Find(&dbRents).Error; err != nil {
		return nil, err
	}

	rents := make([]entity.Rent, len(dbRents))
	for i, rnt := range dbRents {
		rents[i] = ToRentEntity(rnt)
	}
	return rents, nil
}

func (r *RentRepository) ExistsOverlap(carID uint, startDate, endDate time.Time) (bool, error) {
	var count int64
	err := r.db.Model(&Rent{}).
		Where("car_id = ?", carID).
		Where("NOT (? < start_date OR ? > end_date)", endDate, startDate).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *RentRepository) CancelRent(rentID uint) (int, error) {
	var rent Rent

	if err := r.db.First(&rent, "id = ?", rentID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusNotFound, fmt.Errorf("аренда не найдена")
		}
		return http.StatusInternalServerError, err
	}

	if rent.Status == entity.StatusCanceled {
		return http.StatusConflict, fmt.Errorf("аренда уже отменена")
	}

	if time.Now().After(rent.StartDate.Add(-24 * time.Hour)) {
		return http.StatusForbidden, fmt.Errorf("нельзя отменить менее чем за 1 день до начала")
	}

	if err := r.db.Model(&rent).Update("status", entity.StatusCanceled).Error; err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}


func (r *RentRepository) ListByCarID(carID uint) ([]entity.Rent, error) {
    var rents []entity.Rent
    err := r.db.Where("car_id = ? AND end_date >= CURRENT_DATE", carID).Find(&rents).Error
    return rents, err
}

func (r *RentRepository) GetBusyCars(start, end time.Time) ([]uint, error) {
	var rents []entity.Rent
	if err := r.db.Model(&entity.Rent{}).
		Where("status = ?", entity.StatusActive).
		Where("NOT (end_date < ? OR start_date > ?)", start, end).
		Find(&rents).Error; err != nil {
		return nil, err
	}

	carIDs := make([]uint, 0, len(rents))
	for _, rent := range rents {
		carIDs = append(carIDs, rent.CarID)
	}

	return carIDs, nil
}