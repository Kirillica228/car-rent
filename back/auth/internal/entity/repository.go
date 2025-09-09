package entity

type AuthRepository interface {
	Create(authUser *AuthUser) error
	GetByEmail(email string) (*AuthUser,error)
}
