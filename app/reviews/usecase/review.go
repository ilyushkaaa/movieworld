package reviewusecase

import (
	"kinopoisk/app/dto"
	"kinopoisk/app/entity"
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
}

func NewReviewGRPCClient(grpcClient review.ReviewMakerClient) *ReviewGRPCClient {
	return &ReviewGRPCClient{
		mu:         &sync.RWMutex{},
		grpcClient: grpcClient,
	}
}

func (r *ReviewGRPCClient) GetFilmReviews(filmID uint64) ([]*entity.Review, error) {
	return nil, nil
}

func (r *ReviewGRPCClient) NewReview(newReview *dto.ReviewDTO, filmID, userID uint64) (*entity.Review, error) {
	return nil, nil
}

func (r *ReviewGRPCClient) DeleteReview(reviewID, userID uint64) (bool, error) {
	return false, nil
}

func (r *ReviewGRPCClient) UpdateReview(reviewToUpdate *dto.ReviewDTO, reviewID, userID uint64) (*entity.Review, error) {
	return nil, nil
}
