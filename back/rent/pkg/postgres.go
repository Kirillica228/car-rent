package pkg

import (
	postgresmodel "rent/internal/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}
	err = db.AutoMigrate(
		&postgresmodel.Rent{},
	)
	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}
	return db
}