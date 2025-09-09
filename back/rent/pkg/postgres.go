package pkg

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"rent/internal/repository/postgres"
)

func ConnectDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open("rent.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	// Миграция таблицы Rent
	if err := db.AutoMigrate(&repository.RentModel{}); err != nil {
		log.Fatal("failed to migrate database")
	}
	return db
}