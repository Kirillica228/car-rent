package ticket

import (
	"crm/internal/entity"
	"crm/internal/repository"
)

func toEntity(ticket *repository.TicketModel) *entity.Ticket {
	return &entity.Ticket{
		ID:         ticket.ID,
		UserId:    ticket.UserId,
		RankID:     ticket.RankID,
		CaseID:     ticket.CaseID,
		Status:     entity.TicketStatus(ticket.Status),
		EmployeeID: ticket.EmployeeID,
		CreatedAt:  ticket.CreatedAt,
	}
}

func toModel(ticket *entity.Ticket) *repository.TicketModel {
	return &repository.TicketModel{
		ID:         ticket.ID,
		UserId:     ticket.UserId,
		RankID:     ticket.RankID,
		CaseID:     ticket.CaseID,
		Status:     repository.TicketStatus(ticket.Status),
		EmployeeID: ticket.EmployeeID,
		CreatedAt:  ticket.CreatedAt,
	}
}