package reviewservicerepo

import (
	"database/sql"
	"go.uber.org/zap"
	review "kinopoisk/service_review/proto"
)

type ReviewRepo interface {
	GetFilmReviewsRepo(filmID uint64) ([]*review.Review, error)
	NewReviewRepo(newReview *review.Review, filmID, userID uint64) (*review.Review, error)
	DeleteReviewRepo(reviewID uint64) (bool, error)
	UpdateReviewRepo(reviewToUpdate *review.Review) (*review.Review, error)
	GetReviewByFilmUser(filmID, userID uint64) (uint64, error)
	GetUserReviewByID(reviewID, userID uint64) (uint64, error)
}

type ReviewRepoMySQL struct {
	db     *sql.DB
	logger *zap.SugaredLogger
}

func NewReviewRepoMySQL(db *sql.DB, logger *zap.SugaredLogger) *ReviewRepoMySQL {
	return &ReviewRepoMySQL{
		db:     db,
		logger: logger,
	}
}

func (r *ReviewRepoMySQL) GetFilmReviewsRepo(filmID uint64) ([]*review.Review, error) {

}

func (r *ReviewRepoMySQL) NewReviewRepo(newReview *review.Review, filmID, userID uint64) (*review.Review, error) {

}

func (r *ReviewRepoMySQL) DeleteReviewRepo(reviewID uint64) (bool, error) {

}

func (r *ReviewRepoMySQL) UpdateReviewRepo(reviewToUpdate *review.Review) (*review.Review, error) {

}

func (r *ReviewRepoMySQL) GetReviewByFilmUser(filmID, userID uint64) (uint64, error) {

}

func (r *ReviewRepoMySQL) GetUserReviewByID(reviewID, userID uint64) (uint64, error) {

}
