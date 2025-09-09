package transport

import (
	"auth/internal/entity"
	"auth/internal/interfaces/http/dto"

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

	c.JSON(200, gin.H{"message": "Пользователь авторизован", "token": token})

}
