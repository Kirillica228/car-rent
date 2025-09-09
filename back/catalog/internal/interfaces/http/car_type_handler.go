package http

import (
	"catalog/internal/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TypeHandler struct {
	uc entity.CarTypeUseCase
}

func NewTypeHandler(uc entity.CarTypeUseCase) *TypeHandler {
	return &TypeHandler{uc: uc}
}

func (h *TypeHandler) RegisterPublicRoutes(r *gin.RouterGroup) {
	r.GET("/", h.listPublic)
}

func (h *TypeHandler) RegisterAdminRoutes(r *gin.RouterGroup) {
	r.GET("/:id", h.getByID)
	r.GET("/list", h.listAdmin)
	r.POST("/create", h.create)
	r.PUT("/:id", h.update)
	r.DELETE("/:id", h.delete)
}

func (h *TypeHandler) listPublic(c *gin.Context) {
	types, err := h.uc.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, types)
}

func (h *TypeHandler) listAdmin(c *gin.Context) {
	types, err := h.uc.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, types)
}


func (h *TypeHandler) getByID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	t, err := h.uc.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, t)
}

func (h *TypeHandler) create(c *gin.Context) {
	var t entity.CarType
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.uc.Create(t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, err)
}

func (h *TypeHandler) update(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	var t entity.CarType
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.uc.Update(uint(id), t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, err)
}

func (h *TypeHandler) delete(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	if err := h.uc.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
