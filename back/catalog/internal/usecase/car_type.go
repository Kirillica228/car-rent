package usecase

import "catalog/internal/entity"

type CarTypeUC struct {
	repo entity.CarTypeRepository
}

func NewCarTypeUseCase(r entity.CarTypeRepository) entity.CarTypeUseCase {
	return &CarTypeUC{repo: r}
}

func (uc *CarTypeUC) GetByID(id uint) (entity.CarType, error) {
	return uc.repo.GetByID(id)
}

func (uc *CarTypeUC) List() ([]entity.CarType, error) {
	return uc.repo.List()
}

func (uc *CarTypeUC) ListAdmin() ([]entity.CarType, error) {
	return uc.repo.ListAdmin()
}

func (uc *CarTypeUC) Create(t entity.CarType) error {
	return uc.repo.Save(t)
}

func (uc *CarTypeUC) Update(id uint, t entity.CarType) error {
	return uc.repo.Update(id, t)
}

func (uc *CarTypeUC) Delete(id uint) error {
	return uc.repo.Delete(id)
}
