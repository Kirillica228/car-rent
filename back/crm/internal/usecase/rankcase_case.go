package usecase

import "crm/internal/entity"

type RankCaseUseCase interface {
	Save(rankCase *entity.RankCase) error
	GetAll() ([]*entity.RankCase, error)
	GetByID(id uint) (*entity.RankCase, error)
}

type RankCaseUC struct {
	repo entity.RankCaseRepository
}

func NewRankCaseUseCase(repo entity.RankCaseRepository) RankCaseUseCase {
	return &RankCaseUC{repo: repo}
}

func (r *RankCaseUC) Save(rankCase *entity.RankCase) error {
	return r.repo.Save(rankCase)
}

func (r *RankCaseUC) GetAll() ([]*entity.RankCase, error) {
	return r.repo.GetAll()
}

func (r *RankCaseUC) GetByID(id uint) (*entity.RankCase, error) {
	return r.repo.GetByID(id)
}
