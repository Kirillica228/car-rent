package http

import (
	"catalog/internal/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BrandHandler struct {
	uc entity.CarBrandUseCase
}

func NewBrandHandler(uc entity.CarBrandUseCase) *BrandHandler {
	return &BrandHandler{uc: uc}
}

func (h *BrandHandler) RegisterPublicRoutes(r *gin.RouterGroup) {
	r.GET("/", h.listPublic)
}

func (h *BrandHandler) RegisterAdminRoutes(r *gin.RouterGroup) {
	r.GET("/:id", h.getByID)
	r.GET("/list", h.listAdmin)
	r.POST("/create", h.create)
	r.PUT("/:id", h.update)
	r.DELETE("/:id", h.delete)
}

func (h *BrandHandler) listPublic(c *gin.Context) {
	brands, err := h.uc.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, brands)
}

func (h *BrandHandler) listAdmin(c *gin.Context) {
	brands, err := h.uc.ListAdmin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, brands)
}

func (h *BrandHandler) getByID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	brand, err := h.uc.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, brand)
}

func (h *BrandHandler) create(c *gin.Context) {
	var brand entity.CarBrand
	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.uc.Create(brand)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, err)
}

func (h *BrandHandler) update(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	var brand entity.CarBrand
	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.uc.Update(uint(id), brand)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, err)
}

func (h *BrandHandler) delete(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	if err := h.uc.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
