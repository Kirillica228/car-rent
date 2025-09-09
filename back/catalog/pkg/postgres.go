package pkg

import (
	postgresmodel "catalog/internal/repository/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}
	err = db.AutoMigrate(&postgresmodel.CarModel{}, &postgresmodel.CarTypeModel{}, &postgresmodel.CarBrandModel{})
	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}
	return db
}