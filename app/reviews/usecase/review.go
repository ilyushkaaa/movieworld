package reviewusecase

import (
	"context"
	"kinopoisk/app/dto"
	"kinopoisk/app/entity"
	errorapp "kinopoisk/app/errors"
	filmrepo "kinopoisk/app/films/repo/mysql"
	review "kinopoisk/service_review/proto"
	"sync"
)

type ReviewUseCase interface {
	GetFilmReviews(filmID uint64) ([]*entity.Review, error)
	NewReview(newReview *dto.ReviewDTO, filmID, userID uint64) (*entity.Review, error)
	DeleteReview(reviewID, userID uint64) (bool, error)
	UpdateReview(reviewToUpdate *dto.ReviewDTO, reviewID, userID uint64) (*entity.Review, error)
}

type ReviewGRPCClient struct {
	mu         *sync.RWMutex
	grpcClient review.ReviewMakerClient
	filmRepo   filmrepo.FilmRepo
}

func NewReviewGRPCClient(grpcClient review.ReviewMakerClient, filmRepo filmrepo.FilmRepo) *ReviewGRPCClient {
	return &ReviewGRPCClient{
		grpcClient: grpcClient,
		filmRepo:   filmRepo,
		mu:         &sync.RWMutex{},
	}
}

func (r *ReviewGRPCClient) GetFilmReviews(filmID uint64) ([]*entity.Review, error) {
	reviews, err := r.grpcClient.GetFilmReviews(context.Background(), &review.FilmID{
		ID: filmID,
	})
	if err != nil {
		return nil, err
	}
	reviewsArr := reviews.GetReviews()
	reviewsApp := make([]*entity.Review, len(reviewsArr))
	for i, currentReview := range reviewsArr {
		newReviewApp := getReviewFromGRPCStruct(currentReview)
		reviewsApp[i] = newReviewApp
	}
	return reviewsApp, nil
}

func (r *ReviewGRPCClient) NewReview(newReview *dto.ReviewDTO, filmID, userID uint64) (*entity.Review, error) {
	film, err := r.filmRepo.GetFilmByIDRepo(filmID)
	if err != nil {
		return nil, err
	}
	if film == nil {
		return nil, errorapp.ErrorNoFilm
	}
	newReviewGRPC, err := r.grpcClient.NewReview(context.Background(), &review.NewReviewData{
		Review: getGRPCReviewFromDTO(newReview),
		FilmID: &review.FilmID{ID: filmID},
		UserID: &review.UserID{ID: userID},
	})
	if err != nil {
		return nil, err
	}
	if newReviewGRPC == nil {
		return nil, nil
	}
	reviewApp := getReviewFromGRPCStruct(newReviewGRPC)
	return reviewApp, nil
}

func (r *ReviewGRPCClient) DeleteReview(reviewID, userID uint64) (bool, error) {
	isDeletedGRPC, err := r.grpcClient.DeleteReview(context.Background(), &review.DeleteReviewData{
		ReviewID: &review.ReviewID{ID: reviewID},
		UserID:   &review.UserID{ID: userID},
	})
	if err != nil {
		return false, err
	}
	return isDeletedGRPC.IsDeleted, nil
}

func (r *ReviewGRPCClient) UpdateReview(reviewToUpdate *dto.ReviewDTO, reviewID, userID uint64) (*entity.Review, error) {
	grpcReview := getGRPCReviewFromDTO(reviewToUpdate)
	grpcReview.ID = &review.ReviewID{
		ID: reviewID,
	}
	updatedReviewGRPC, err := r.grpcClient.UpdateReview(context.Background(), &review.UpdateReviewData{
		Review: grpcReview,
		UserID: &review.UserID{ID: userID},
	})
	if err != nil {
		return nil, err
	}
	updatedReviewApp := getReviewFromGRPCStruct(updatedReviewGRPC)
	return updatedReviewApp, nil
}

func getReviewFromGRPCStruct(reviewGRPC *review.Review) *entity.Review {
	return &entity.Review{
		ID:      reviewGRPC.ID.ID,
		Mark:    reviewGRPC.Mark,
		Comment: reviewGRPC.Comment,
	}
}

func getGRPCReviewFromDTO(reviewDTO *dto.ReviewDTO) *review.Review {
	return &review.Review{
		Mark:    reviewDTO.Mark,
		Comment: reviewDTO.Comment,
	}
}
