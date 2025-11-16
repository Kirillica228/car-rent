package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("supersecretkey") // вынеси в конфиг

// роли
const (
	RoleUser   = "user"
	RoleAdmin  = "admin"
	RoleWorker = "worker"
)

// AuthMiddleware принимает список разрешённых ролей
func AuthMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenStr string

		// 1️⃣ — Пытаемся достать токен из cookie
		if cookie, err := c.Cookie("access_token"); err == nil && cookie != "" {
			tokenStr = cookie
		}

		// 2️⃣ — Если нет cookie, пробуем Authorization: Bearer
		if tokenStr == "" {
			authHeader := c.GetHeader("Authorization")
			if strings.HasPrefix(authHeader, "Bearer ") {
				tokenStr = strings.TrimPrefix(authHeader, "Bearer ")
			}
		}

		// 3️⃣ — Если токена нет вообще
		if tokenStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			c.Abort()
			return
		}

		// 4️⃣ — Парсим JWT токен
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		// 5️⃣ — Достаём claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid claims"})
			c.Abort()
			return
		}

		// 6️⃣ — Извлекаем user_id и role
		userID, _ := claims["user_id"].(float64)
		role, _ := claims["role"].(string)

		// 7️⃣ — Проверяем доступ по роли
		if len(allowedRoles) > 0 {
			allowed := false
			for _, r := range allowedRoles {
				if r == role {
					allowed = true
					break
				}
			}

			if !allowed {
				log.Printf("access denied: role=%s", role)
				c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
				c.Abort()
				return
			}
		}

		// 8️⃣ — Добавляем данные в контекст
		c.Set("user_id", uint(userID))
		c.Set("role", role)
		c.Request.Header.Set("X-User-Role", role)

		c.Next()
	}
}
