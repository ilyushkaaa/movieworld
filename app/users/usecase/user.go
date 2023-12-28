package userusecase

import "kinopoisk/app/entity"

type UserUseCase interface {
	Login(username, password string) (*entity.User, error)
	Register(username, password string) (*entity.User, error)
	CreateSession(user *entity.User) (string, error)
	GetSession(token string) (*entity.Session, error)
	DeleteSession(token string) error
}
