package postgres

import (
	"catalog/internal/entity"

	"gorm.io/gorm"
)

type postgresTypeRepo struct {
	db *gorm.DB
}

func (p *postgresTypeRepo) GetByID(id uint) (entity.CarType, error) {
	var t CarTypeModel
	if err := p.db.First(&t, id).Error; err != nil {
		return entity.CarType{}, err
	}
	return toEntityType(t), nil
}

func (p *postgresTypeRepo) List() ([]entity.CarType, error) {
	var types []CarTypeModel
	if err := p.db.Find(&types).Error; err != nil {
		return nil, err
	}
	return toEntityTypeList(types), nil
}

func (p *postgresTypeRepo) ListAdmin() ([]entity.CarType, error) {
	var types []CarTypeModel
	if err := p.db.Where("is_visible = ?", true).Find(&types).Error; err != nil {
		return nil, err
	}
	return toEntityTypeList(types), nil
}

func (p *postgresTypeRepo) Save(t entity.CarType) error {
	model := toModelType(t)
	if err := p.db.Create(&model).Error; err != nil {
		return err
	}
	t.ID = model.ID
	return nil
}

func (p *postgresTypeRepo) Update(id uint, t entity.CarType) error {
	var existing CarTypeModel
	if err := p.db.First(&existing, id).Error; err != nil {
		return err
	}

	existing.Name = t.Name

	if err := p.db.Save(&existing).Error; err != nil {
		return err
	}

	return nil
}

func (p *postgresTypeRepo) Delete(id uint) error {
	return p.db.Delete(&CarTypeModel{}, id).Error
}

func NewPostgresTypeRepo(db *gorm.DB) entity.CarTypeRepository {
	return &postgresTypeRepo{db: db}
}
