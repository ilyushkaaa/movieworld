package reviewservicerepo

import (
	"database/sql"
	"errors"
	"go.uber.org/zap"
	errorreview "kinopoisk/service_review/error"
	review "kinopoisk/service_review/proto"
)

type ReviewRepo interface {
	GetFilmReviewsRepo(filmID uint64) ([]*review.Review, error)
	NewReviewRepo(newReview *review.Review, filmID, userID uint64) (*review.Review, error)
	DeleteReviewRepo(reviewID uint64) (bool, error)
	UpdateReviewRepo(reviewToUpdate *review.Review) (*review.Review, error)
	GetReviewByFilmUser(filmID, userID uint64) (uint64, error)
	GetUserReviewByID(reviewID, userID uint64) (*review.Review, error)
	ChangeRatingAfterDeleteReview(oldReview *review.Review, reviewID uint64)
	ChangeRatingAfterUpdateReview(oldReview, newReview *review.Review, reviewID uint64)
	ChangeRatingAddReview(newReview *review.Review, reviewID uint64)
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
	reviews := []*review.Review{}
	rows, err := r.db.Query("SELECT r.id, r.mark, r.comment, r.user_id, u.username FROM reviews r JOIN users u ON r.user_id = u.id WHERE r.film_id = ?", filmID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			r.logger.Errorf("error in closing db rows")
		}
	}(rows)
	for rows.Next() {
		newReview := &review.Review{}
		newReview.ID = &review.ReviewID{}
		newReview.Author = &review.User{
			ID: &review.UserID{},
		}
		err = rows.Scan(&newReview.ID.ID, &newReview.Mark, &newReview.Comment, &newReview.Author.ID.ID, &newReview.Author.Username)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, newReview)
	}
	return reviews, nil

}

func (r *ReviewRepoMySQL) NewReviewRepo(newReview *review.Review, filmID, userID uint64) (*review.Review, error) {
	res, err := r.db.Exec(
		"INSERT INTO reviews (`mark`, `comment`, `user_id`, `film_id`) VALUES (?, ?, ?, ?)",
		newReview.Mark,
		newReview.Comment,
		userID,
		filmID,
	)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	newReview.ID = &review.ReviewID{
		ID: uint64(id),
	}
	err = r.getReviewAuthor(newReview)
	if err != nil {
		return nil, err
	}

	return newReview, nil

}

func (r *ReviewRepoMySQL) DeleteReviewRepo(reviewID uint64) (bool, error) {
	_, err := r.db.Exec(
		"DELETE FROM reviews WHERE id = ?",
		reviewID,
	)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *ReviewRepoMySQL) UpdateReviewRepo(reviewToUpdate *review.Review) (*review.Review, error) {
	_, err := r.db.Exec(
		"UPDATE reviews SET mark = ?, comment = ? where id = ?",
		reviewToUpdate.Mark,
		reviewToUpdate.Comment,
		reviewToUpdate.ID.ID,
	)
	if err != nil {
		return nil, err
	}
	err = r.getReviewAuthor(reviewToUpdate)
	if err != nil {
		return nil, err
	}
	return reviewToUpdate, nil
}

func (r *ReviewRepoMySQL) getReviewAuthor(reviewToUpdate *review.Review) error {
	reviewToUpdate.Author = &review.User{
		ID: &review.UserID{},
	}
	err := r.db.
		QueryRow("SELECT u.id, u.username from users u JOIN reviews r on u.id = r.user_id WHERE r.id = ?", reviewToUpdate.ID.ID).
		Scan(&reviewToUpdate.Author.ID.ID, &reviewToUpdate.Author.Username)
	return err
}

func (r *ReviewRepoMySQL) GetReviewByFilmUser(filmID, userID uint64) (uint64, error) {
	var ID uint64
	err := r.db.
		QueryRow("SELECT id from reviews WHERE film_id = ? AND user_id = ?", filmID, userID).
		Scan(&ID)

	if errors.Is(err, sql.ErrNoRows) {
		return 0, errorreview.ErrorNoReview
	}
	if err != nil {
		return 0, err
	}
	return ID, nil
}

func (r *ReviewRepoMySQL) GetUserReviewByID(reviewID, userID uint64) (*review.Review, error) {
	foundReview := &review.Review{}
	foundReview.ID = &review.ReviewID{}
	err := r.db.
		QueryRow("SELECT id, mark from reviews WHERE id = ? AND user_id = ?", reviewID, userID).
		Scan(&foundReview.ID.ID, &foundReview.Mark)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, errorreview.ErrorNoReview
	}
	if err != nil {
		return nil, err
	}
	return foundReview, nil
}

func (r *ReviewRepoMySQL) ChangeRatingAfterDeleteReview(oldReview *review.Review, reviewID uint64) {
	_, err := r.db.Exec(
		`UPDATE films 
                SET 
                   sum_mark = sum_mark - ?,
                num_of_marks = num_of_marks - 1,
                rating = CASE 
                              WHEN num_of_marks > 1 THEN (sum_mark - ?) / (num_of_marks - 1)
                            ELSE 0
                          END
                WHERE id IN (SELECT film_id FROM reviews WHERE id = ?)`,
		oldReview.Mark,
		oldReview.Mark,
		reviewID,
	)
	if err != nil {
		r.logger.Errorf("error in changing rating after delete review")
	}
}

func (r *ReviewRepoMySQL) ChangeRatingAfterUpdateReview(oldReview, newReview *review.Review, reviewID uint64) {
	_, err := r.db.Exec(
		"UPDATE films SET sum_mark = sum_mark + ? - ?, rating = (sum_mark + ? - ?) / num_of_marks WHERE id in (SELECT film_id from reviews WHERE id = ?)",
		newReview.Mark,
		oldReview.Mark,
		newReview.Mark,
		oldReview.Mark,
		reviewID,
	)
	if err != nil {
		r.logger.Errorf("error in changing rating after update review: %s", err)
	}
}

func (r *ReviewRepoMySQL) ChangeRatingAddReview(newReview *review.Review, reviewID uint64) {
	_, err := r.db.Exec(
		"UPDATE films SET sum_mark = sum_mark + ?, num_of_marks = num_of_marks + 1, rating = (sum_mark + ?) / (num_of_marks + 1) WHERE id in (SELECT film_id from reviews WHERE id = ?)",
		newReview.Mark,
		newReview.Mark,
		reviewID,
	)
	if err != nil {
		r.logger.Errorf("error in changing rating after add review")
	}
}
