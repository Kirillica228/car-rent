package main

import (
	"catalog/config"
	httpHandlers "catalog/internal/interfaces/http"
	"catalog/internal/middleware"
	"catalog/internal/repository/postgres"
	"catalog/internal/usecase"
	"catalog/pkg"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	// подключение
	db := pkg.ConnectDB(dsn)

	// Repo
	carRepo := postgres.NewPostgresCarRepo(db)
	brandRepo := postgres.NewPostgresBrandRepo(db)
	typeRepo := postgres.NewPostgresTypeRepo(db)

	// Usecases
	carUC := usecase.NewCarUseCase(carRepo)
	brandUC := usecase.NewCarBrandUseCase(brandRepo)
	typeUC := usecase.NewCarTypeUseCase(typeRepo)

	// Handlers
	carHandler := httpHandlers.NewCarHandler(carUC)
	brandHandler := httpHandlers.NewBrandHandler(brandUC)
	typeHandler := httpHandlers.NewTypeHandler(typeUC)

	// Gin
	r := gin.Default()

	api := r.Group("/api")
	public := api.Group("/catalog")
	{
		carHandler.RegisterPublicRoutes(public.Group("/cars"))
		brandHandler.RegisterPublicRoutes(public.Group("/brands"))
		typeHandler.RegisterPublicRoutes(public.Group("/types"))
	}

	admin := api.Group("/catalog/admin")
	admin.Use(middleware.RequireAdmin())
	{
		carHandler.RegisterAdminRoutes(admin.Group("/cars"))
		brandHandler.RegisterAdminRoutes(admin.Group("/brands"))
		typeHandler.RegisterAdminRoutes(admin.Group("/types"))
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
