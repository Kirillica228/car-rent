package pkg

import (
	"catalog/internal/domain/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}
	err = db.AutoMigrate(
		&entity.Car{}, 
		&entity.Type{}, 
		&entity.Brand{}, 
		&entity.Image{})
	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}
	return db
}