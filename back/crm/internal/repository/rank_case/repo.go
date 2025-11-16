package rankcase

import (
	"crm/internal/entity"
	"crm/internal/repository"

	"gorm.io/gorm"
)

type RankCaseRepo struct {
	db *gorm.DB
}

func NewRankCaseRepo(db *gorm.DB) entity.RankCaseRepository {
	return &RankCaseRepo{db: db}
}

func (r *RankCaseRepo) Save(rankCase *entity.RankCase) error {
	model := toModel(rankCase)
	if err := r.db.Create(model).Error; err != nil {
		return err
	}
	rankCase.ID = model.ID 
	return nil
}

func (r *RankCaseRepo) GetByID(id uint) (*entity.RankCase, error) {
	var model repository.CaseModel
	if err := r.db.First(&model, id).Error; err != nil {
		return nil, err
	}
	return toEntity(&model), nil
}

func (r *RankCaseRepo) GetAll() ([]*entity.RankCase, error) {
	var models []repository.CaseModel
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}

	result := make([]*entity.RankCase, len(models))
	for i, m := range models {
		result[i] = toEntity(&m)
	}
	return result, nil
}

func (r *RankCaseRepo) Update(id uint, rankCase *entity.RankCase) error {
	var model repository.CaseModel
	if err := r.db.First(&model, id).Error; err != nil {
		return err
	}

	model.Name = rankCase.Name
	return r.db.Save(&model).Error
}

func (r *RankCaseRepo) Delete(id uint) error {
	return r.db.Delete(&repository.CaseModel{}, id).Error
}
