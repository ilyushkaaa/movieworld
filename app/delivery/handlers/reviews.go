package handlers

import (
	"go.uber.org/zap"
	reviewusecase "kinopoisk/app/reviews/usecase"
)

type ReviewHandler struct {
	ReviewUseCases reviewusecase.ReviewUseCase
	Logger         *zap.SugaredLogger
}

func NewReviewHandler(reviewUseCases reviewusecase.ReviewUseCase, logger *zap.SugaredLogger) *ReviewHandler {
	return &ReviewHandler{
		ReviewUseCases: reviewUseCases,
		Logger:         logger,
	}
}
