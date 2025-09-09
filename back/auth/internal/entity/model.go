package entity


type AuthUser struct {
	ID		uint
	Email    string 
	Password string 
	Role     int
}

type AuthToken struct{
	Token	string
}