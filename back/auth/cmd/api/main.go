package main

import (
	"auth/config"
	httpHandlers "auth/internal/interfaces/http"
	"auth/internal/repository/postgres"
	"auth/internal/usecase"
	"auth/pkg"
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
	// Repos
	AuthRepo := postgres.NewPostgresAuthRepo(db)

	// Usecases
	AuthUC := usecase.NewAuthUseCase(AuthRepo)

	// Handlers
	AuthHandler := httpHandlers.NewAuthHandler(AuthUC)

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
		AuthHandler.RegisterRoutes(api.Group("/auth"))
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
