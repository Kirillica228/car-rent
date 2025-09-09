package usecase

import "catalog/internal/entity"

type CarUseCase struct {
	repo entity.CarRepository
}

func NewCarUseCase(r entity.CarRepository) entity.CarUseCase {
	return &CarUseCase{repo: r}
}

func (uc *CarUseCase) GetByID(id uint) (entity.Car, error) {
	return uc.repo.GetByID(id)
}

func (uc *CarUseCase) List() ([]entity.Car, error) {
	return uc.repo.List()
}

func (uc *CarUseCase) ListAdmin() ([]entity.Car, error) {
	return uc.repo.ListAdmin()
}

func (uc *CarUseCase) Create(car entity.Car) error {
	return uc.repo.Save(car)
}

func (uc *CarUseCase) Update(id uint, car entity.Car) error {
	return uc.repo.Update(id, car)
}

func (uc *CarUseCase) Delete(id uint) error {
	return uc.repo.Delete(id)
}
