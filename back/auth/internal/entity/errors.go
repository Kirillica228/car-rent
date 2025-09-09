package entity

import "errors"

var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrDataJSON = errors.New("не валидные данные")
	ErrTokenCreate = errors.New("ошибка токена")
)
