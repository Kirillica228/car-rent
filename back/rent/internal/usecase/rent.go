package usecase

import (
	"fmt"
	"rent/internal/entity"
	"time"

	"rent/internal/client/catalog_service"
)

type RentUseCase interface {
	GetAllByUserID(id uint) ([]entity.Rent, error)
	GetByID(id uint) (entity.Rent, error)
	GetAll() ([]entity.Rent, error)
	Save(rent entity.Rent) error
	Update(rent entity.Rent) error
	Delete(id uint) error
	CancelRent(rentID uint) (int, error)
	GetBookingIntervals(carID uint) ([]entity.BookingInterval, error) 
	GetBusyCars(start, end time.Time) ([]uint, error)
}

type RentUC struct {
	repo entity.RentRepository
	cars catalog_service.CarsService
}


func NewRentUsecase(repo entity.RentRepository, cars catalog_service.CarsService) RentUseCase {
	return &RentUC{repo: repo, cars: cars}
}

func (r *RentUC) Delete(id uint) error {
	return r.repo.Delete(id)
}

func (r *RentUC) GetAll() ([]entity.Rent, error) {
	return r.repo.GetAll()
}

func (r *RentUC) GetAllByUserID(id uint) ([]entity.Rent, error) {
	return r.repo.GetAllByUserID(id)
}

func (r *RentUC) GetByID(id uint) (entity.Rent, error) {
	return r.repo.GetByID(id)
}

func (r *RentUC) Save(rent entity.Rent) error {
	// 1. Получаем машину из catalog_service
	car, err := r.cars.GetCarByID(rent.CarID)
	if err != nil {
		return fmt.Errorf("не удалось получить машину из catalog_service: %w", err)
	}

	// 2. Проверяем бизнес-логику
	if car.ID == 0 {
		return fmt.Errorf("машина с ID %d не найдена", rent.CarID)
	}
	if car.Price <= 0 {
		return fmt.Errorf("у машины %d некорректная цена", car.ID)
	}

	// 3. Проверяем пересечение дат аренды
	overlap, err := r.repo.ExistsOverlap(rent.CarID, rent.StartDate, rent.EndDate)
	if err != nil {
		return fmt.Errorf("ошибка при проверке доступности машины: %w", err)
	}
	if overlap {
		return fmt.Errorf("машина %d уже забронирована на выбранный период", rent.CarID)
	}

	// 4. Считаем стоимость
	days := int(rent.EndDate.Sub(rent.StartDate).Hours()/24) + 1
	if days <= 0 {
		return fmt.Errorf("некорректный период аренды")
	}

	rent.CarID = car.ID
	rent.TotalPrice = float64(days) * car.Price
	rent.Name = car.Brand + " " + car.Model + " " + car.Type
	rent.Image = car.Image

	return r.repo.Save(rent)
}

func (r *RentUC) Update(rent entity.Rent) error {
	return r.repo.Update(rent)
}

func (r *RentUC) CancelRent(rentID uint) (int, error) {
	return r.repo.CancelRent(rentID)
}

func (r *RentUC) GetBusyCars(start, end time.Time) ([]uint, error) {
	return r.repo.GetBusyCars(start, end)
}

func (s *RentUC) GetBookingIntervals(carID uint) ([]entity.BookingInterval, error) {
    rents, err := s.repo.ListByCarID(carID)
    if err != nil {
        return nil, err
    }

    intervals := make([]entity.BookingInterval, len(rents))
    for i, r := range rents {
        intervals[i] = entity.BookingInterval{
            StartDate: r.StartDate,
            EndDate:   r.EndDate,
        }
    }
    return intervals, nil
}