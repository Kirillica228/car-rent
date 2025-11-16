package http

import (
	"gateway/internal/middleware"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")

	api.GET("authorize/me", middleware.AuthMiddleware(), func(c *gin.Context) {
		userID, _ := c.Get("user_id")
		role, _ := c.Get("role")
		log.Println("userID:", userID, "role:", role)
		c.JSON(http.StatusOK, gin.H{
			"user_id": userID,
			"role":    role,
		})
	})
	

	auth := api.Group("/auth")
	{
		auth.Any("/*proxyPath", middleware.ReverseProxy("http://auth:8080/api/auth"))
	}

	catalog := api.Group("/catalog")
	{
		catalog.Any("/*proxyPath", middleware.ReverseProxy("http://catalog:8080/api"))
	}

	rent := api.Group("/rents")
	rent.Use(middleware.AuthMiddleware())
	{
		rent.Any("/*proxyPath", middleware.ReverseProxy("http://rent:8080/api/rents"))
	}

	admin := api.Group("/admin")
	admin.Use(middleware.AuthMiddleware(middleware.RoleAdmin))
	{
		cars := admin.Group("/cars")
		{
			cars.Any("/*proxyPath", middleware.ReverseProxy("http://catalog:8080/api/admin/cars"))
		}
		brands := admin.Group("/brands")
		{
			brands.Any("/*proxyPath", middleware.ReverseProxy("http://catalog:8080/api/admin/brands"))
		}
		types := admin.Group("/types")
		{
			types.Any("/*proxyPath", middleware.ReverseProxy("http://catalog:8080/api/admin/types"))
		}
	}

}
