package rank

import (
	"crm/internal/entity"
	"crm/internal/repository"
)

func toEntity(rank *repository.RankModel) *entity.Rank {
    cases := make([]string, len(rank.Cases))
    for i, c := range rank.Cases {
        cases[i] = c.Name
    }

    return &entity.Rank{
        ID:          rank.ID,
        Name:        rank.Name,
        Description: rank.Description,
        Cases:       cases,
    }
}

func toModel(rank *entity.Rank) *repository.RankModel {
    cases := make([]repository.CaseModel, len(rank.Cases))
    for i, name := range rank.Cases {
        cases[i] = repository.CaseModel{Name: name}
    }

    return &repository.RankModel{
        ID:          rank.ID,
        Name:        rank.Name,
        Description: rank.Description,
        Cases:       cases,
    }
}
