package controller

import (
	"crm/internal/entity"
	"crm/internal/interfaces/dto/rankcase"
	"crm/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RankCaseController struct {
	usecase usecase.RankCaseUseCase
}

func NewRankCaseController(usecase usecase.RankCaseUseCase) RankCaseController {
	return RankCaseController{usecase: usecase}
}

func (rc *RankCaseController) RegisterRoutes(r *gin.Engine) {
	r.GET("/cases", rc.GetAll)
	r.GET("/cases/:id", rc.GetByID)
	r.POST("/cases", rc.Save)
}

func (rc *RankCaseController) GetAll(c *gin.Context) {
	cases, err := rc.usecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cases)
}

func (rc *RankCaseController) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	rankCase, err := rc.usecase.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "case not found"})
		return
	}
	c.JSON(http.StatusOK, rankCase)
}

func (rc *RankCaseController) Save(c *gin.Context) {
	var dto rankcase.RankCaseDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	rankCase := &entity.RankCase{
		Name: dto.Name,
	}

	if err := rc.usecase.Save(rankCase); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, rankCase)
}
