package handlers

import (
	"go.uber.org/zap"
	userusecase "kinopoisk/app/users/usecase"
)

type UserHandler struct {
	UserUseCases userusecase.UserUseCase
	Logger       *zap.SugaredLogger
}

func NewUserHandler(userUseCases userusecase.UserUseCase, logger *zap.SugaredLogger) *UserHandler {
	return &UserHandler{
		UserUseCases: userUseCases,
		Logger:       logger,
	}
}
