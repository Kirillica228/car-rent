package entity

type RankCase struct {
	ID    uint       
	Name  string    
	RankId  uint       
}

type RankCaseRepository interface {
	Save(rankCase *RankCase) error
	GetByID(id uint) (*RankCase, error)
	GetAll() ([]*RankCase, error)
	Delete(id uint) error
	Update(id uint, rankCase *RankCase) error
}