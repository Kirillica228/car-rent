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
	dbModel := UserModel{
		Email:    authUser.Email,
		Password: authUser.Password,
	}

	return r.db.Create(&dbModel).Error
}

func (r *postgresAuthRepo) GetByEmail(email string) (*entity.AuthUser,error) {
	var dbModel UserModel
	if err := r.db.Where("email = ?", email).First(&dbModel).Error; err != nil {
		return nil, err
	}
	return &entity.AuthUser{
		ID: dbModel.ID,
		Email:    dbModel.Email,
		Password: dbModel.Password,
		Role:     dbModel.Role,
	}, nil
}