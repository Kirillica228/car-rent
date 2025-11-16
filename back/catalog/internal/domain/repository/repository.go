package repository

type Repository[T any] interface{
	Save(t T) error
	List() ([]T, error)
	GetByID(id uint) (T, error)
	Update(id uint, t T) error
	Delete(ids []uint) error
}