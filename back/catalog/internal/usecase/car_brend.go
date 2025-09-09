package usecase

import "catalog/internal/entity"

type CarBrandUseCase struct {
	repo entity.CarBrandRepository
}

func NewCarBrandUseCase(r entity.CarBrandRepository) entity.CarBrandUseCase {
	return &CarBrandUseCase{repo: r}
}

func (uc *CarBrandUseCase) GetByID(id uint) (entity.CarBrand, error) {
	return uc.repo.GetByID(id)
}

func (uc *CarBrandUseCase) List() ([]entity.CarBrand, error) {
	return uc.repo.List()
}

func (uc *CarBrandUseCase) ListAdmin() ([]entity.CarBrand, error) {
	return uc.repo.ListAdmin()
}

func (uc *CarBrandUseCase) Create(brand entity.CarBrand) error {
	return uc.repo.Save(brand)
}

func (uc *CarBrandUseCase) Update(id uint, brand entity.CarBrand) error {
	return uc.repo.Update(id, brand)
}

func (uc *CarBrandUseCase) Delete(id uint) error {
	return uc.repo.Delete(id)
}
