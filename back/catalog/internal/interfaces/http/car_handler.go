package http

import (
	"catalog/internal/domain/entity"
	"catalog/internal/interfaces/http/dto"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CarHandler struct {
	uc entity.CarUseCase
}

func NewCarHandler(uc entity.CarUseCase) *CarHandler {
	return &CarHandler{uc: uc}
}

func (h *CarHandler) RegisterPublicRoutes(r *gin.RouterGroup) {
	r.GET("/list", h.listPublic)
	r.GET("/:id", h.getByID)
}

func (h *CarHandler) RegisterAdminRoutes(r *gin.RouterGroup) {
	r.GET("/:id", h.getByID)
	r.GET("/list", h.listAdmin)
	r.POST("/create", h.create)
	r.PUT("/:id", h.update)
	r.DELETE("/delete", h.delete)

	// фотки
	r.POST("/:id/photos", h.uploadPhotos)
	r.DELETE("/:id/photos", h.deletePhotos)
}

func (h *CarHandler) listPublic(c *gin.Context) {
	var params dto.ListCarsParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cars, err := h.uc.List(params, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cars)
}

func (h *CarHandler) listAdmin(c *gin.Context) {
	var params dto.ListCarsParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cars, err := h.uc.List(params, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cars)
}

func (h *CarHandler) getByID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	car, err := h.uc.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, car)
}

func (h *CarHandler) create(c *gin.Context) {
	if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
	c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка парсинга формы: " + err.Error()})
	return
	}

	// Извлекаем поля формы
	brandID, _ := strconv.Atoi(c.PostForm("brand_id"))
	typeID, _ := strconv.Atoi(c.PostForm("type_id"))
	year, _ := strconv.Atoi(c.PostForm("year"))
	price, _ := strconv.ParseFloat(c.PostForm("price"), 64)
	isVisible := c.PostForm("is_visible") == "true"

	// Новые поля
	color := c.PostForm("color")
	description := c.PostForm("description")
	transmission := c.PostForm("transmission")
	status := c.PostForm("status") // допустимые значения: "На диагностике", "Готова к аренде", "Недоступна"

	car := entity.Car{
		BrandID:      uint(brandID),
		TypeID:       uint(typeID),
		Model:        c.PostForm("model"),
		Year:         year,
		LicensePlate: c.PostForm("license_plate"),
		Status:       entity.CarStatus(status),
		Price:        price,
		IsVisible:    isVisible,
		Color:        color,
		Description:  description,
		Transmission: entity.Transmission(transmission),
	}
	log.Default().Println(car)
	// --- Создаём машину в БД ---
	if err := h.uc.Create(&car); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании машины: " + err.Error()})
		return
	}

	// --- Загружаем фото (если есть) ---
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка формы: " + err.Error()})
		return
	}

	files := form.File["photos"]
	var urls []string

	for _, header := range files {
		file, err := header.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer file.Close()

		// ⚡ используем уже готовую бизнес-логику
		url, err := h.uc.UploadPhoto(car.ID, file, header)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при загрузке фото: " + err.Error()})
			return
		}

		urls = append(urls, url)
	}

	// --- Ответ ---
	c.JSON(http.StatusCreated, gin.H{
		"message": "Машина успешно создана",
		"car":     car,
		"photos":  urls,
	})
}

func (h *CarHandler) update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
	c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка парсинга формы: " + err.Error()})
	return
	}

	// Извлекаем поля формы
	brandID, _ := strconv.Atoi(c.PostForm("brand_id"))
	typeID, _ := strconv.Atoi(c.PostForm("type_id"))
	year, _ := strconv.Atoi(c.PostForm("year"))
	price, _ := strconv.ParseFloat(c.PostForm("price"), 64)
	isVisible := c.PostForm("is_visible") == "true"

	// Новые поля
	color := c.PostForm("color")
	description := c.PostForm("description")
	transmission := c.PostForm("transmission")
	status := c.PostForm("status") // допустимые значения: "На диагностике", "Готова к аренде", "Недоступна"

	car := entity.Car{
		BrandID:      uint(brandID),
		TypeID:       uint(typeID),
		Model:        c.PostForm("model"),
		Year:         year,
		LicensePlate: c.PostForm("license_plate"),
		Status:       entity.CarStatus(status),
		Price:        price,
		IsVisible:    isVisible,
		Color:        color,
		Description:  description,
		Transmission: entity.Transmission(transmission),
	}

	// Обновляем основные данные машины
	if err := h.uc.Update(uint(id), car); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// --- Работа с фото ---
	form := c.Request.MultipartForm

	// 1️⃣ Удаляем фото, которые пользователь убрал
	removed := form.Value["removed_photos"]
	if len(removed) > 0 {
		if err := h.uc.DeletePhotos(removed); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete old photos"})
			return
		}
	}

	// 2️⃣ Добавляем новые фото
	files := form.File["photos"]
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to open uploaded file"})
			return
		}
		defer file.Close()

		if _, err := h.uc.UploadPhoto(uint(id), file, fileHeader); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save photo"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "car updated successfully"})
}


func (h *CarHandler) delete(c *gin.Context) {
	var req dto.DeleteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if len(req.ID) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no ids provided"})
		return
	}

	if err := h.uc.Delete(req.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully"})
}


func (h *CarHandler) uploadPhotos(c *gin.Context) {
    idStr := c.Param("id")
    id, _ := strconv.Atoi(idStr)

    form, err := c.MultipartForm()
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "неверная форма"})
        return
    }

    files := form.File["photos"]
    if len(files) == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "нужен хотя бы один файл"})
        return
    }

    var urls []string
    for _, header := range files {
        file, err := header.Open()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        defer file.Close()

        url, err := h.uc.UploadPhoto(uint(id), file, header)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        urls = append(urls, url)
    }

    c.JSON(http.StatusCreated, gin.H{"urls": urls})
}

func (h *CarHandler) deletePhotos(c *gin.Context) {
    var req struct {
        URLs []string `json:"urls" binding:"required"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if len(req.URLs) == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "нужен хотя бы один url"})
        return
    }

    if err := h.uc.DeletePhotos(req.URLs); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.Status(http.StatusNoContent)
}
