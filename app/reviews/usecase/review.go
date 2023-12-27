package reviewusecase

import "kinopoisk/app/entity"

type ReviewUseCase interface {
	GetFilmReviews(filmID uint64) ([]*entity.Review, error)
	NewReview(newReview *entity.Review, filmID, userID int) (*entity.Review, error)
	DeleteReview(filmID, userID int) (bool, error)
	UpdateReview(reviewToUpdate *entity.Review, userID int) (*entity.Review, error)
}
