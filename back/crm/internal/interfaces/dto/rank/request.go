package rank

type RankDTO struct {
	ID          uint   `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
}