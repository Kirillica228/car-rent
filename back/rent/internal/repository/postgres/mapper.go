package repository

import (
	"rent/internal/entity"
)

// DB → бизнес
func ToRentEntity(db RentModel) entity.Rent {
	return entity.Rent{
		ID:        db.ID,
		UserID:    db.UserID,
		CarID:     db.CarID,
		StartTime: db.StartTime,
		EndTime:   db.EndTime,
		Status:    db.Status,
	}
}

// бизнес → DB
func ToRentDB(ent entity.Rent) RentModel{
	return RentModel{
		ID:        ent.ID,
		UserID:    ent.UserID,
		CarID:     ent.CarID,
		StartTime: ent.StartTime,
		EndTime:   ent.EndTime,
		Status:    ent.Status,
	}
}
