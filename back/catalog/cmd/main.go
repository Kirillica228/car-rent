package main

import (
	"catalog/config"
	"catalog/internal/domain/entity"
	httpHandlers "catalog/internal/interfaces/http"
	"catalog/internal/middleware"
	repo "catalog/internal/repository"
	file "catalog/internal/repository/image"
	"catalog/internal/usecase"
	"catalog/pkg"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
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

	storagePath := "./storage/photos"
	baseURL := "/files"

	// Repo
	carRepo := repo.NewPostgresCarRepo(db)
	brandRepo := repo.NewPostgresRepository[entity.Brand](db)
	typeRepo := repo.NewPostgresRepository[entity.Type](db)
	fileRepo := file.NewFileRepository(storagePath, baseURL)

	// Usecases
	carUC := usecase.NewCarUseCase(carRepo, fileRepo)
	brandUC := usecase.NewUseCase(brandRepo)
	typeUC := usecase.NewUseCase(typeRepo)

	// Handlers
	carHandler := httpHandlers.NewCarHandler(carUC)
	brandHandler := httpHandlers.NewHandler[entity.Brand](brandUC)
	typeHandler := httpHandlers.NewHandler[entity.Type](typeUC)

	// Gin
	gin.SetMode(gin.DebugMode) // 🔥 Включаем debug-режим GIN
	r := gin.New()

	// Добавляем встроенный логгер и recovery
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // твой фронт
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.RedirectTrailingSlash = false
	// Регистрируем все маршруты

	api := r.Group("/api")
	{
		api.Static("/files", "./storage/photos")
		carHandler.RegisterPublicRoutes(api.Group("/cars"))
		brandHandler.RegisterRoutes(api.Group("/brands"))
		typeHandler.RegisterRoutes(api.Group("/types"))

		admin := api.Group("/admin")
		admin.Use(middleware.RequireUserHeaders(true))
		{
			carHandler.RegisterAdminRoutes(admin.Group("/cars"))
			brandHandler.RegisterRoutes(admin.Group("/brands"))
			typeHandler.RegisterRoutes(admin.Group("/types"))
		}
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
