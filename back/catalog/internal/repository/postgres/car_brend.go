package postgres

import (
	"catalog/internal/entity"

	"gorm.io/gorm"
)

type postgresBrandRepo struct {
	db *gorm.DB
}

func (p *postgresBrandRepo) GetByID(id uint) (entity.CarBrand, error) {
	var brand CarBrandModel
	if err := p.db.First(&brand, id).Error; err != nil {
		return entity.CarBrand{}, err
	}
	return toEntityBrand(brand), nil
}

func (p *postgresBrandRepo) List() ([]entity.CarBrand, error) {
	var brands []CarBrandModel
	if err := p.db.Find(&brands).Error; err != nil {
		return nil, err
	}
	return toEntityBrandList(brands), nil
}

func (p *postgresBrandRepo) ListAdmin() ([]entity.CarBrand, error) {
	var brands []CarBrandModel
	if err := p.db.Where("is_visible = ?", true).Find(&brands).Error; err != nil {
		return nil, err
	}
	return toEntityBrandList(brands), nil
}


func (p *postgresBrandRepo) Save(brand entity.CarBrand) error {
	model := toModelBrand(brand)
	if err := p.db.Create(&model).Error; err != nil {
		return err
	}
	brand.ID = model.ID
	return nil
}

func (p *postgresBrandRepo) Update(id uint, brand entity.CarBrand) error {
	var existing CarBrandModel
	if err := p.db.First(&existing, id).Error; err != nil {
		return err
	}

	existing.Name = brand.Name

	if err := p.db.Save(&existing).Error; err != nil {
		return err
	}

	return nil
}

func (p *postgresBrandRepo) Delete(id uint) error {
	return p.db.Delete(&CarBrandModel{}, id).Error
}

func NewPostgresBrandRepo(db *gorm.DB) entity.CarBrandRepository {
	return &postgresBrandRepo{db: db}
}
