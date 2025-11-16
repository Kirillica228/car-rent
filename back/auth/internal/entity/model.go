package entity


type AuthUser struct {
	ID		uint
	Email    string 
	Password string 
	Role     Role
}

type Role struct {
	ID   uint
	Name string
}

type AuthToken struct{
	Token	string
}