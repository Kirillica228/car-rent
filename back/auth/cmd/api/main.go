package main

import (
	"auth/config"
	httpHandlers "auth/internal/interfaces/http"
	"auth/internal/repository/postgres"
	"auth/internal/usecase"
	"auth/pkg"
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
	// Repos
	AuthRepo := postgres.NewPostgresAuthRepo(db)

	// Usecases
	AuthUC := usecase.NewAuthUseCase(AuthRepo)

	// Handlers
	AuthHandler := httpHandlers.NewAuthHandler(AuthUC)

	// Gin
	r := gin.Default()

	api := r.Group("/api")
	{
		AuthHandler.RegisterRoutes(api.Group("/auth"))
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
