package usecase

type UseCase[T any] interface{
	GetByID(id uint) (T, error)
	List() ([]T, error)
	Create(t T) error
	Update(id uint, t T) error
	Delete(ids []uint) error
}