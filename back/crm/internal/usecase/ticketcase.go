package usecase

import "crm/internal/entity"

type TicketUseCase interface {
	GetAll() ([]*entity.Ticket, error)
	GetByID(id uint) (*entity.Ticket, error)
	GetByUserID(userID uint) ([]*entity.Ticket, error)
	GetByEmployeeId(employeeID uint) ([]*entity.Ticket, error)
	Save(ticket *entity.Ticket) error
}

type TicketUC struct {
	repo entity.TicketRepository
}

func NewTicketUseCase(repo entity.TicketRepository) TicketUseCase {
	return &TicketUC{repo: repo}
}

func (t *TicketUC) GetAll() ([]*entity.Ticket, error) {
	return t.repo.GetAll()
}

func (t *TicketUC) GetByID(id uint) (*entity.Ticket, error) {
	return t.repo.GetByID(id)
}

func (t *TicketUC) GetByUserID(userID uint) ([]*entity.Ticket, error) {
	return t.repo.GetByUserID(userID)
}

func (t *TicketUC) GetByEmployeeId(employeeID uint) ([]*entity.Ticket, error) {
	return t.repo.GetByEmployeeId(employeeID)
}

func (t *TicketUC) Save(ticket *entity.Ticket) error {
	return t.repo.Save(ticket)
}
