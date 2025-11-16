package rank

import (
	"crm/internal/entity"
	"crm/internal/repository"

	"gorm.io/gorm"
)

type RankRepo struct {
	db *gorm.DB
}

func NewRankRepo(db *gorm.DB) entity.RankRepository {
	return &RankRepo{db: db}
}


func (r *RankRepo) Save(rank *entity.Rank) error {
	model := toModel(rank)

	if err := r.db.Create(model).Error; err != nil {
		return err
	}
	return nil
}

func (r *RankRepo) GetByID(id uint) (*entity.Rank, error) {
	var model repository.RankModel
	if err := r.db.Preload("Cases").First(&model, id).Error; err != nil {
		return nil, err
	}
	return toEntity(&model), nil
}

func (r *RankRepo) GetAll() ([]*entity.Rank, error) {
	var models []repository.RankModel
	if err := r.db.Preload("Cases").Find(&models).Error; err != nil {
		return nil, err
	}

	result := make([]*entity.Rank, len(models))
	for i, m := range models {
		result[i] = toEntity(&m)
	}
	return result, nil
}

func (r *RankRepo) Update(id uint, rank *entity.Rank) error {
	var model repository.RankModel
	if err := r.db.Preload("Cases").First(&model, id).Error; err != nil {
		return err
	}

	
	model.Name = rank.Name
	model.Description = rank.Description

	newCases := toModel(rank).Cases
	if err := r.db.Model(&model).Association("Cases").Replace(newCases); err != nil {
		return err
	}

	return r.db.Save(&model).Error
}

func (r *RankRepo) Delete(id uint) error {
	return r.db.Delete(&repository.RankModel{}, id).Error
}
