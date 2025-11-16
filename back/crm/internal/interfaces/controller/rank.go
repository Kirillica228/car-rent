package controller

import (
	"crm/internal/entity"
	"crm/internal/interfaces/dto/rank"
	"crm/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RankController struct {
	usecase usecase.RankUseCase
}

func NewRankController(usecase usecase.RankUseCase) RankController {
	return RankController{usecase: usecase}
}

func (rc *RankController) RegisterRoutes(r *gin.Engine) {
	r.GET("/ranks", rc.GetAll)
	r.GET("/ranks/:id", rc.GetByID)
	r.POST("/ranks", rc.Save)
}

func (rc *RankController) GetAll(c *gin.Context) {
	ranks, err := rc.usecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ranks)
}

func (rc *RankController) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	rank, err := rc.usecase.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "rank not found"})
		return
	}
	c.JSON(http.StatusOK, rank)
}

func (rc *RankController) Save(c *gin.Context) {
	var dto rank.RankDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	rank := &entity.Rank{
		Name:        dto.Name,
		Description: dto.Description,
	}

	if err := rc.usecase.Save(rank); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, rank)
}
