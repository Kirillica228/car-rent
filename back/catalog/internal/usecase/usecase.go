package usecase

import (
	"catalog/internal/domain/repository"
	"catalog/internal/domain/usecase"
)

type UseCase[T any] struct {
	Repo repository.Repository[T]
}

func NewUseCase[T any](repo repository.Repository[T]) usecase.UseCase[T] {
	return UseCase[T]{Repo: repo}
}

func (u UseCase[T]) Create(t T) error {
	return u.Repo.Save(t)
}

func (u UseCase[T]) Delete(ids []uint) error {
	return u.Repo.Delete(ids)
}

func (u UseCase[T]) GetByID(id uint) (T, error) {
	return u.Repo.GetByID(id)
}

func (u UseCase[T]) List() ([]T, error) {
	return u.Repo.List()
}

func (u UseCase[T]) Update(id uint, t T) error {
	return u.Repo.Update(id, t)
}
