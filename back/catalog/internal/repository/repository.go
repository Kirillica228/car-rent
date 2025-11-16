package repository

import (
	"catalog/internal/domain/repository"

	"gorm.io/gorm"
)

type postgresRepository[T any] struct {
	db *gorm.DB
}

func NewPostgresRepository[T any](db *gorm.DB) repository.Repository[T] {
	return &postgresRepository[T]{db: db}
}

func (p *postgresRepository[T]) Save(t T) error {
	return p.db.Create(&t).Error
}

func (p *postgresRepository[T]) List() ([]T, error) {
	var items []T
	if err := p.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (p *postgresRepository[T]) GetByID(id uint) (T, error) {
	var item T
	if err := p.db.First(&item, id).Error; err != nil {
		return item, err
	}
	return item, nil
}

func (p *postgresRepository[T]) Update(id uint, t T) error {
	return p.db.Debug().Model(&t).Where("id = ?", id).Select("*").Updates(&t).Error
}

func (p *postgresRepository[T]) Delete(ids []uint) error {
	var t T
	return p.db.Unscoped().Delete(&t, ids).Error
}