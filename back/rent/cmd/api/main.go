package main

import (
	"fmt"
	"log"
	"rent/config"
	"rent/internal/client/catalog_service"
	httpHandlers "rent/internal/interfaces/http"
	"rent/internal/middleware"
	"rent/internal/repository"
	"rent/internal/usecase"
	"rent/pkg"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.Load()

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	db := pkg.ConnectDB(dsn)

	catalogService := catalog_service.NewCarService("http://catalog:8080/api")

	rentRepo := repository.NewRentRepository(db)
	rentUC := usecase.NewRentUsecase(rentRepo, catalogService)
	rentHandler := httpHandlers.NewRentHandler(rentUC)


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
	public := api.Group("/rents")
	public.Use(middleware.RequireUserHeaders(false))
	{
		rentHandler.RegisterPublicRoutes(public)
	}

	admin := api.Group("/rents/admin")
	admin.Use(middleware.RequireUserHeaders(true))
	{
		rentHandler.RegisterAdminRoutes(admin)
	}
	
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
