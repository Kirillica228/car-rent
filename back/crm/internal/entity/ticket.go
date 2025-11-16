package entity

import "time"

type TicketStatus string

const (
	StatusCreated   TicketStatus = "CREATED"
	StatusInProcess TicketStatus = "IN_PROCESS"
	StatusClosed    TicketStatus = "CLOSED"
)

type Ticket struct {
	ID         uint
	UserId     uint
	RankID     uint
	CaseID     uint
	Subject    string
	Status     TicketStatus
	EmployeeID *uint     
	CreatedAt  time.Time
}

type TicketRepository interface{
	Save(ticket *Ticket) error
	GetByID(id uint) (*Ticket, error)
	GetAll() ([]*Ticket, error)
	GetByUserID(userID uint) ([]*Ticket, error)
	GetByEmployeeId(employeeID uint) ([]*Ticket, error)
}