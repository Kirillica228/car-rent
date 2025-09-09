package http

import (
	"catalog/internal/entity"
	"catalog/internal/interfaces/http/dto"
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
	r.GET("/", h.listPublic)
}

func (h *CarHandler) RegisterAdminRoutes(r *gin.RouterGroup) {
	r.GET("/:id", h.getByID)
	r.GET("/list", h.listAdmin)
	r.POST("/create", h.create)
	r.PUT("/:id", h.update)
	r.DELETE("/:id", h.delete)
}

func (h *CarHandler) listPublic(c *gin.Context) {
	cars, err := h.uc.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cars)
}

func (h *CarHandler) listAdmin(c *gin.Context) {
	cars, err := h.uc.ListAdmin()
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
	var req dto.CreateCarRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	car := entity.Car{
		Brand:         entity.CarBrand{ID: req.BrandID},
		Type:          entity.CarType{ID: req.TypeID},
		Model:         req.ModelCar,
		Year:          req.Year,
		License_plate: req.LicensePlate,
		Status:        req.Status,
		Price:         req.Price,
		IsVisible:     req.IsVisible,
	}

	if err := h.uc.Create(car); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, car)
}

func (h *CarHandler) update(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	var req dto.CreateCarRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	car := entity.Car{
		Brand:         entity.CarBrand{ID: req.BrandID},
		Type:          entity.CarType{ID: req.TypeID},
		Model:         req.ModelCar,
		Year:          req.Year,
		License_plate: req.LicensePlate,
		Status:        req.Status,
		Price:         req.Price,
		IsVisible:     req.IsVisible,
	}

	err := h.uc.Update(uint(id), car)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, err)
}

func (h *CarHandler) delete(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	if err := h.uc.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
