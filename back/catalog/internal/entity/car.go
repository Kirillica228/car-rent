package entity

type Car struct {
	ID   uint
	Brand CarBrand
	Type CarType
	Model string
	Year int
	License_plate string
	Status 	string
	Price float32
	IsVisible bool
}

type CarType struct {
	ID   uint
	Name string
	IsVisible bool
}

type CarBrand struct {
	ID   uint
	Name string
	IsVisible bool
}

type CarRepository interface {
	List() ([]Car, error)
	GetByID(id uint) (Car, error)
	Save(car Car) error
	Delete(id uint) error
	Update(id uint, car Car) error
	ListAdmin() ([]Car, error)
}

type CarBrandRepository interface {
	List() ([]CarBrand, error)
	GetByID(id uint) (CarBrand, error)
	Save(car CarBrand) error
	Delete(id uint) error
	Update(id uint, car CarBrand) error
	ListAdmin() ([]CarBrand, error)
}

type CarTypeRepository interface {
	List() ([]CarType, error)
	GetByID(id uint) (CarType, error)
	Save(car CarType) error
	Delete(id uint) error
	Update(id uint, car CarType) error
	ListAdmin() ([]CarType, error)
}

type CarUseCase interface {
	GetByID(id uint) (Car, error)
	List() ([]Car, error)
	Create(car Car) error
	Update(id uint, car Car) error
	Delete(id uint) error
	ListAdmin() ([]Car, error)
}

type CarBrandUseCase interface {
	GetByID(id uint) (CarBrand, error)
	List() ([]CarBrand, error)
	Create(brand CarBrand) error
	Update(id uint, brand CarBrand) error
	Delete(id uint) error
	ListAdmin() ([]CarBrand, error)
}

type CarTypeUseCase interface {
	GetByID(id uint) (CarType, error)
	List() ([]CarType, error)
	Create(t CarType) error
	Update(id uint, t CarType) error
	Delete(id uint) error
	ListAdmin() ([]CarType, error)
}