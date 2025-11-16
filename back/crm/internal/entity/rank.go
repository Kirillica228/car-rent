package entity

type Rank struct {
	ID          uint       
	Name        string    
	Description string    
	Cases       []string         
}

type RankRepository interface {
	Save(rank *Rank) error
	GetByID(id uint) (*Rank, error)
	GetAll() ([]*Rank, error)
	Delete(id uint) error
	Update(id uint, rank *Rank) error
}