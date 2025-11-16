package transport

import (
	"auth/internal/entity"
	"auth/internal/interfaces/http/dto"
	"auth/internal/jwt"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase entity.AuthCase
}

func NewAuthHandler(usecase entity.AuthCase) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/register", h.register)
	r.POST("/login", h.login)
	r.POST("logout",h.Logout)
	r.GET("/validate", h.validate)
}

func (h *Handler) register(c *gin.Context) {
	var jsonData dto.RegisterRequest
	if err:=c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(400,gin.H{"error":"Не валидные данные"})
		return
	}
	user := entity.AuthUser{
		Email: jsonData.Email,
		Password: jsonData.Password,
	}
	if err := h.usecase.Register(&user); err != nil {
		if err == entity.ErrUserAlreadyExists {
			c.JSON(409, gin.H{"error": "Такой пользователь уже существует"})
			return
		}
		c.JSON(500, gin.H{"error": "внутренняя ошибка сервиса"})
		return
	}

	c.JSON(201,gin.H{"message":"Пользователь зарегистрирован"})
}
func (h *Handler) login(c *gin.Context) {
	var jsonData dto.LoginRequest
	if err:=c.ShouldBindJSON(&jsonData);err != nil {
		c.JSON(400,gin.H{"error":"Не валидные данные"})
		return
	}
	user := entity.AuthUser{
		Email: jsonData.Email,
		Password: jsonData.Password,
	}
	token, err := h.usecase.Login(&user)
	if err != nil {
		if err == entity.ErrInvalidCredentials {
			c.JSON(401, gin.H{"error": "почта или пароль не верны"})
			return
		}
		c.JSON(500, gin.H{"error": "внутренняя ошибка сервиса"})
		return
	}

	c.SetCookie(
		"access_token", 
		token.Token,       
		3600,          
		"/",            
		"",            
		true,           
		true,          
	)

	c.JSON(200, gin.H{
		"message": "Пользователь авторизован",
	})

}
func (h *Handler) Logout(c *gin.Context) {
    // Сбрасываем cookie
    c.SetCookie(
        "access_token", // имя cookie
        "",             // пустое значение
        -1,             // max-age = -1 → удаление
        "/",            // path
        "",             // domain, если нужно
        true,           // secure (https)
        true,           // httpOnly
    )

    c.JSON(200, gin.H{
        "message": "Logged out",
    })
}

func (h *Handler) validate(c *gin.Context) {
	cookie, err := c.Cookie("access_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "cookie not found"})
		return
	}

	claims, err := jwt.ParseToken(cookie)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}

	// Устанавливаем заголовки — Nginx их передаст дальше
	c.Header("X-User-ID", fmt.Sprintf("%d", claims.UserID))
	c.Header("X-User-Role", claims.Role)

	c.Status(http.StatusOK)
}

