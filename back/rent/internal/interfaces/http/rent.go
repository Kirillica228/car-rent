package http

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"rent/internal/entity"
	"rent/internal/interfaces/http/dto"
	"rent/internal/usecase"

	"github.com/gin-gonic/gin"
)

type RentHandler struct {
	useCase usecase.RentUseCase
}

func NewRentHandler(u usecase.RentUseCase) *RentHandler {
	return &RentHandler{useCase: u}
}

func (h *RentHandler) RegisterPublicRoutes(r *gin.RouterGroup) {
	r.GET("/list", h.GetAllByUserID) // список рентов текущего юзера
	r.GET("/:id", h.GetByID)         // один рент по id
	r.POST("/save", h.Save)          // создать новый рент
	r.PUT("/cancel/:id", h.CancelRent)
	r.GET("/bookings/:car_id", h.GetCarBookings)
	r.POST("/busy_cars", h.GetBusyCars)
}

func (h *RentHandler) RegisterAdminRoutes(r *gin.RouterGroup) {
	r.GET("/list", h.ListAdmin)
	r.GET("/:id", h.GetByIDAdmin)
	r.POST("/create", h.Create)
	r.PUT("/:id", h.UpdateAdmin)
	r.DELETE("/:id", h.DeleteAdmin)
}

func (h *RentHandler) ListAdmin(c *gin.Context) {
	rents, err := h.useCase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rents)
}

func (h *RentHandler) GetByIDAdmin(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	rent, err := h.useCase.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rent)
}

func (h *RentHandler) Create(c *gin.Context) {
	var rent entity.Rent
	if err := c.ShouldBindJSON(&rent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.useCase.Save(rent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, rent)
}

func (h *RentHandler) UpdateAdmin(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var rent entity.Rent
	if err := c.ShouldBindJSON(&rent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rent.ID = uint(id)
	if err := h.useCase.Update(rent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rent)
}

func (h *RentHandler) DeleteAdmin(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.useCase.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *RentHandler) GetAllByUserID(c *gin.Context) {
	userID := c.GetUint("X-User-Id") 
	rents, err := h.useCase.GetAllByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rents)
}

func (h *RentHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	rent, err := h.useCase.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rent)
}

func (h *RentHandler) Save(c *gin.Context) {
	var rent dto.CreateRentDTO

	if err := c.ShouldBindJSON(&rent); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rent.StartDate = time.Date(rent.StartDate.Year(), rent.StartDate.Month(), rent.StartDate.Day(), 0, 0, 0, 0, time.UTC)
	rent.EndDate = time.Date(rent.EndDate.Year(), rent.EndDate.Month(), rent.EndDate.Day(), 0, 0, 0, 0, time.UTC)

	result := entity.Rent{
		UserID:    c.GetUint("X-User-Id"),
		CarID:     rent.CarID,
		StartDate: rent.StartDate,
		EndDate:   rent.EndDate,
	}

	if err := h.useCase.Save(result); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (h *RentHandler) CancelRent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID"})
		return
	}

	status, err := h.useCase.CancelRent(uint(id))
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
func (h *RentHandler) GetCarBookings(c *gin.Context) {
    carIDStr := c.Param("car_id")
    carID, err := strconv.Atoi(carIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid car_id"})
        return
    }

    intervals, err := h.useCase.GetBookingIntervals(uint(carID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, intervals)
}

func (h *RentHandler) GetBusyCars(c *gin.Context) {
	var req dto.GetBusyCarsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	carIDs, err := h.useCase.GetBusyCars(req.StartDate, req.EndDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := make([]dto.BusyCarResponse, 0, len(carIDs))
	for _, id := range carIDs {
		res = append(res, dto.BusyCarResponse{CarID: id})
	}

	c.JSON(http.StatusOK, res)
}