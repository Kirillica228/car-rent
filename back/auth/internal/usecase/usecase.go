package usecase

import (
	"auth/internal/entity"
	"auth/internal/jwt"

	"golang.org/x/crypto/bcrypt"
)

type UseCase struct {
	repo entity.AuthRepository
}

func NewAuthUseCase(r entity.AuthRepository) entity.AuthCase{
	return &UseCase{repo : r}
}

func (u *UseCase) Register(authUser *entity.AuthUser) error {
	existing,err := u.repo.GetByEmail(authUser.Email)
	if err == nil && existing != nil {
		return entity.ErrUserAlreadyExists
	}
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(authUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	authUser.Password = string(hashedPassword)
	return u.repo.Create(authUser)
}

func (u *UseCase) Login(authUser *entity.AuthUser) (*entity.AuthToken,error) {
	user,err := u.repo.GetByEmail(authUser.Email)
	if err != nil || user == nil {
		return nil,entity.ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(authUser.Password));err != nil {
		return nil,entity.ErrInvalidCredentials
	}
	token,err := jwt.GenerateToken(user.ID,user.Role)
	if err!=nil{
		return nil,entity.ErrTokenCreate
	}
	response := entity.AuthToken{
		Token:token,
	}

	return &response,err
}