package ticket

import (
	"crm/internal/entity"
	"crm/internal/repository"

	"gorm.io/gorm"
)

type TicketRepo struct {
	db *gorm.DB
}

func NewTicketRepo(db *gorm.DB) entity.TicketRepository {
	return &TicketRepo{db: db}
}

// Save — создание нового тикета
func (t *TicketRepo) Save(ticket *entity.Ticket) error {
	model := toModel(ticket)
	if err := t.db.Create(model).Error; err != nil {
		return err
	}
	ticket.ID = model.ID // обновляем ID в entity
	return nil
}

// GetByID — получить тикет по ID
func (t *TicketRepo) GetByID(id uint) (*entity.Ticket, error) {
	var model repository.TicketModel
	if err := t.db.First(&model, id).Error; err != nil {
		return nil, err
	}
	return toEntity(&model), nil
}

// GetAll — список всех тикетов
func (t *TicketRepo) GetAll() ([]*entity.Ticket, error) {
	var models []repository.TicketModel
	if err := t.db.Find(&models).Error; err != nil {
		return nil, err
	}

	result := make([]*entity.Ticket, len(models))
	for i, m := range models {
		result[i] = toEntity(&m)
	}
	return result, nil
}

// GetByUserID — все тикеты пользователя
func (t *TicketRepo) GetByUserID(userID uint) ([]*entity.Ticket, error) {
	var models []repository.TicketModel
	if err := t.db.Where("user_id = ?", userID).Find(&models).Error; err != nil {
		return nil, err
	}

	result := make([]*entity.Ticket, len(models))
	for i, m := range models {
		result[i] = toEntity(&m)
	}
	return result, nil
}

// GetByEmployeeId — все тикеты назначенные сотруднику
func (t *TicketRepo) GetByEmployeeId(employeeID uint) ([]*entity.Ticket, error) {
	var models []repository.TicketModel
	if err := t.db.Where("employee_id = ?", employeeID).Find(&models).Error; err != nil {
		return nil, err
	}

	result := make([]*entity.Ticket, len(models))
	for i, m := range models {
		result[i] = toEntity(&m)
	}
	return result, nil
}
