package main

import (
	"log"

	"gateway/internal/interfaces/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // твой фронт
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.RedirectTrailingSlash = false
	// Регистрируем все маршруты
	http.RegisterRoutes(r)

	log.Println("🚀 Gateway running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
