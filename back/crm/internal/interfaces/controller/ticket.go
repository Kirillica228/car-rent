package controller

import (
	"crm/internal/entity"
	"crm/internal/interfaces/dto/ticket"
	"crm/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TicketController struct {
	usecase usecase.TicketUseCase
}

func NewTicketController(usecase usecase.TicketUseCase) TicketController {
	return TicketController{usecase: usecase}
}

func (tc *TicketController) RegisterRoutes(r *gin.Engine) {
	r.GET("/tickets", tc.GetAll)                   // только админ/работник
	r.GET("/tickets/:id", tc.GetByID)              // инфа о тикете
	r.GET("/my/tickets", tc.GetByUserID)           // тикеты юзера
	r.POST("/tickets", tc.CreateTicket)            // создание тикета
	r.GET("/employee/tickets", tc.GetByEmployeeID) // тикеты сотрудника
}

// Получить все тикеты (админ/CRM)
func (tc *TicketController) GetAll(c *gin.Context) {
	tickets, err := tc.usecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tickets)
}

// Получить тикет по ID
func (tc *TicketController) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	ticket, err := tc.usecase.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ticket not found"})
		return
	}
	c.JSON(http.StatusOK, ticket)
}

// Тикеты текущего пользователя
func (tc *TicketController) GetByUserID(c *gin.Context) {
	userID := c.GetUint("user_id")
	tickets, err := tc.usecase.GetByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tickets)
}

// Тикеты для сотрудника
func (tc *TicketController) GetByEmployeeID(c *gin.Context) {
	employeeID := c.GetUint("user_id")
	tickets, err := tc.usecase.GetByEmployeeId(employeeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tickets)
}

// Создать тикет
func (tc *TicketController) CreateTicket(c *gin.Context) {
	var dto ticket.TicketDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// привязываем тикет к юзеру из токена
	userID := c.GetUint("user_id")

	ticket := &entity.Ticket{
		UserId:     userID,
		RankID:     dto.RankID,
		CaseID:     dto.CaseID,
		Status:     entity.StatusCreated, // по умолчанию
		EmployeeID: nil,
	}

	if err := tc.usecase.Save(ticket); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, ticket)
}
