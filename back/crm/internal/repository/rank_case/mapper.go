package rankcase

import (
	"crm/internal/entity"
	"crm/internal/repository"
)

func toEntity(rankCase *repository.CaseModel) *entity.RankCase {
	return &entity.RankCase{
		ID:     rankCase.ID,
		Name:   rankCase.Name,
	}
}

func toModel(rankCase *entity.RankCase) *repository.CaseModel {
	return &repository.CaseModel{
		ID:   rankCase.ID,
		Name: rankCase.Name,
	}
}