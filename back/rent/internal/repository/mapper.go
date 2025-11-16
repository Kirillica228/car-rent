package repository

import (
	"rent/internal/entity"
)

// DB → бизнес
func ToRentEntity(db Rent) entity.Rent {
	return entity.Rent{
		ID:         db.ID,
		UserID:     db.UserID,
		CarID:      db.CarID,
		Name:       db.Name,
		Image:      db.Image,
		StartDate:  db.StartDate,
		EndDate:    db.EndDate,
		Status:     db.Status,
		TotalPrice: db.TotalPrice,
		CreatedAt:  db.CreatedAt,
	}
}

// бизнес → DB
func ToRentDB(ent entity.Rent) Rent {
	return Rent{
		ID:         ent.ID,
		UserID:     ent.UserID,
		CarID:      ent.CarID,
		Name:       ent.Name,
		Image:      ent.Image,
		StartDate:  ent.StartDate,
		EndDate:    ent.EndDate,
		Status:     ent.Status,
		TotalPrice: ent.TotalPrice,
		CreatedAt:  ent.CreatedAt,
	}
}
