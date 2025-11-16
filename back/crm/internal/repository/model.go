package repository

import (
	"time"
)

type TicketStatus string

const (
	StatusCreated    TicketStatus = "CREATED"    // создан
	StatusInProgress TicketStatus = "IN_PROGRESS" // в процессе
	StatusClosed     TicketStatus = "CLOSED"      // закрыт
)

type TicketModel struct {
	ID         uint         `gorm:"primaryKey;autoIncrement"`
	UserId     uint       	`gorm:"index;not null"`           
	RankID     uint         `gorm:"index;not null"`                     
	CaseID     uint         `gorm:"index;not null"`                     
	Status     TicketStatus `gorm:"type:varchar(100);default:'CREATED'"` 
	EmployeeID *uint        `gorm:"index;null"`                        
	CreatedAt  time.Time
	UpdatedAt  time.Time
}


type RankModel struct {
	ID          uint       `gorm:"primaryKey;autoIncrement"`
	Name        string     `gorm:"type:varchar(100);unique;not null"` 
	Description string     `gorm:"type:varchar(255)"`                 
	Cases       []CaseModel `gorm:"many2many:rank_cases;"`            
}


type CaseModel struct {
	ID    uint       `gorm:"primaryKey;autoIncrement"`
	Name  string     `gorm:"type:varchar(150);unique;not null"` 
	Ranks []RankModel `gorm:"many2many:rank_cases;"`            
}