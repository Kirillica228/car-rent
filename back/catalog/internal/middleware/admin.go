package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetHeader("X-User-Role")

		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "доступ запрещен: нужны права администратора",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
