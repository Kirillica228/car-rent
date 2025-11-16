package main

import (
	"crm/config"
	"crm/internal/interfaces/controller"
	"crm/internal/usecase"
	ticket_repo "crm/internal/repository/ticket"
	rank_repo "crm/internal/repository/rank"
	rankcase_repo "crm/internal/repository/rank_case"
	"fmt"

	"crm/pkg"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	// подключение
	db := pkg.ConnectDB(dsn)

	// Инициализация репозиториев
	rankRepo := rank_repo.NewRankRepo(db)
	caseRepo := rankcase_repo.NewRankCaseRepo(db)
	ticketRepo := ticket_repo.NewTicketRepo(db)

	// Инициализация usecase
	rankUC := usecase.NewRankUseCase(rankRepo)
	caseUC := usecase.NewRankCaseUseCase(caseRepo)
	ticketUC := usecase.NewTicketUseCase(ticketRepo)

	// Инициализация контроллеров
	rankController := controller.NewRankController(rankUC)
	caseController := controller.NewRankCaseController(caseUC)
	ticketController := controller.NewTicketController(ticketUC)

	// Gin роутер
	r := gin.Default()

	// Регистрация маршрутов
	rankController.RegisterRoutes(r)
	caseController.RegisterRoutes(r)
	ticketController.RegisterRoutes(r)

	// Запуск сервера
	if err := r.Run(":8080"); err != nil {
		panic("failed to run server: " + err.Error())
	}
}
