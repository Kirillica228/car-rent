package pkg

import (
	postgresmodel "auth/internal/repository/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}

	// Автоматические миграции
	err = db.AutoMigrate(&postgresmodel.UserModel{})
	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	return db
}
