package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func RequireUserHeaders(requireAdmin bool) gin.HandlerFunc {
	return func(c *gin.Context) {  
		for k, v := range c.Request.Header {
			if len(v) > 0 && strings.HasPrefix(k, "X-User-") {
				val := v[0]      

				// если это ID — конвертим в uint
				if k == "X-User-Id" {
					if id, err := strconv.Atoi(val); err == nil {
						c.Set(k, uint(id))
						continue
					}
				}

				// иначе оставляем строкой
				c.Set(k, val)
			}
		}

		// Проверка на админа
		if requireAdmin {
			role, _ := c.Get("X-User-Role")
			if roleStr, ok := role.(string); !ok || roleStr != "admin" {
				c.JSON(http.StatusForbidden, gin.H{
					"error": "доступ запрещен: нужны права администратора",
				})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
