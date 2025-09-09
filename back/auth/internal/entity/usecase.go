package entity

type AuthCase interface {
	Register(authUser *AuthUser) (error)
	Login(authUser *AuthUser) (*AuthToken, error)
}