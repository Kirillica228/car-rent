package usecase

import "crm/internal/entity"

type RankUseCase interface {
	Save(rank *entity.Rank) error
	GetAll() ([]*entity.Rank, error)
	GetByID(id uint) (*entity.Rank, error)
}

type RankUC struct {
	repo entity.RankRepository
}

func NewRankUseCase(repo entity.RankRepository) RankUseCase {
	return &RankUC{repo: repo}
}

func (r *RankUC) Save(rank *entity.Rank) error {
	return r.repo.Save(rank)
}

func (r *RankUC) GetAll() ([]*entity.Rank, error) {
	return r.repo.GetAll()
}

func (r *RankUC) GetByID(id uint) (*entity.Rank, error) {
	return r.repo.GetByID(id)
}
