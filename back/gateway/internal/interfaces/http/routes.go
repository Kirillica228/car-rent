package http

import (
	"gateway/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")

	api.GET("authorize/me", middleware.AuthMiddleware(0), func(c *gin.Context) {
			userID, _ := c.Get("user_id")
			role, _ := c.Get("role")

			c.JSON(http.StatusOK, gin.H{
				"user_id": userID,
				"role":    role,
			})
		})

	auth := api.Group("/auth")
	{

		auth.Any("/*proxyPath", middleware.ReverseProxy("http://auth:8080/api/auth"))
		
	}

	admin := api.Group("/admin")
	admin.Use(middleware.AuthMiddleware(1))
	{
		cars := admin.Group("/cars")
		{
			cars.Any("/*proxyPath", middleware.ReverseProxy("http://catalog:8080/api/catalog/admin/cars"))
		}
		brands := admin.Group("/brands")
		{
			brands.Any("/*proxyPath", middleware.ReverseProxy("http://catalog:8080/api/catalog/admin/brands"))
		}
		types := admin.Group("/types")
		{
			types.Any("/*proxyPath", middleware.ReverseProxy("http://catalog:8080/api/catalog/admin/types"))
		}
	}

}
