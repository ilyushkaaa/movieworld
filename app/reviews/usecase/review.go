package reviewusecase

import (
	"kinopoisk/app/dto"
	"kinopoisk/app/entity"
)

type ReviewUseCase interface {
	GetFilmReviews(filmID uint64) ([]*entity.Review, error)
	NewReview(newReview *dto.ReviewDTO, filmID, userID uint64) (*entity.Review, error)
	DeleteReview(reviewID, userID uint64) (bool, error)
	UpdateReview(reviewToUpdate *dto.ReviewDTO, reviewID, userID uint64) (*entity.Review, error)
}
