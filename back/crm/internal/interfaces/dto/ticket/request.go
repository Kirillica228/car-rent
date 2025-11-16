package ticket

type TicketDTO struct {
	ID         uint   `json:"id,omitempty"`
	UserId     uint   `json:"user_id"`
	RankID     uint   `json:"rank_id"`
	CaseID     uint   `json:"case_id"`
	Status     string `json:"status,omitempty"` // created / in_progress / closed
	EmployeeID *uint  `json:"employee_id,omitempty"`
}