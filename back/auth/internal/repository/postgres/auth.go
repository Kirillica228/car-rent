package postgres

import (
	"auth/internal/entity"

	"gorm.io/gorm"
)

type postgresAuthRepo struct {
	db *gorm.DB
}

func NewPostgresAuthRepo(db *gorm.DB) entity.AuthRepository {
	return &postgresAuthRepo{ db:db }
}

func (r *postgresAuthRepo) Create(authUser *entity.AuthUser) error {
	dbModel := User{
		Email:    authUser.Email,
		Password: authUser.Password,
	}

	return r.db.Create(&dbModel).Error
}

func (r *postgresAuthRepo) GetByEmail(email string) (*entity.AuthUser,error) {
	var dbModel User
	if err := r.db.Preload("Role").Where("email = ?", email).First(&dbModel).Error; err != nil {
		return nil, err
	}
	return &entity.AuthUser{
		ID: dbModel.ID,
		Email:    dbModel.Email,
		Password: dbModel.Password,
		Role:     entity.Role{
			ID:   dbModel.Role.ID,
			Name: dbModel.Role.Name,
		},
	}, nil
}