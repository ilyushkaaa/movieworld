package reviewserviceusecse

import (
	"context"
	"errors"
	errorreview "kinopoisk/service_review/error"
	review "kinopoisk/service_review/proto"
	reviewservicerepo "kinopoisk/service_review/repo/mysql"
	"sync"
)

type ReviewGRPCServer struct {
	review.UnimplementedReviewMakerServer

	ReviewRepo reviewservicerepo.ReviewRepo
	mu         *sync.RWMutex
}

func NewReviewGRPCServer(reviewRepo reviewservicerepo.ReviewRepo) *ReviewGRPCServer {
	return &ReviewGRPCServer{
		UnimplementedReviewMakerServer: review.UnimplementedReviewMakerServer{},
		ReviewRepo:                     reviewRepo,
		mu:                             &sync.RWMutex{},
	}
}

func (rs *ReviewGRPCServer) GetFilmReviews(_ context.Context, in *review.FilmID) (*review.Reviews, error) {
	rs.mu.RLock()
	reviews, err := rs.ReviewRepo.GetFilmReviewsRepo(in.GetID())
	rs.mu.RUnlock()
	if err != nil {
		return &review.Reviews{}, err
	}
	return &review.Reviews{
		Reviews: reviews,
	}, nil
}

func (rs *ReviewGRPCServer) NewReview(_ context.Context, in *review.NewReviewData) (*review.Review, error) {
	rs.mu.RLock()
	_, err := rs.ReviewRepo.GetReviewByFilmUser(in.GetFilmID().ID, in.GetUserID().ID)
	rs.mu.RUnlock()
	if err == nil {
		return &review.Review{}, nil
	}
	if err != nil {
		if !errors.Is(err, errorreview.ErrorNoReview) {
			return &review.Review{}, err
		}
	}
	rs.mu.Lock()
	newReview, err := rs.ReviewRepo.NewReviewRepo(in.GetReview(), in.GetFilmID().ID, in.GetUserID().ID)
	rs.mu.Unlock()
	if err != nil {
		return &review.Review{}, err
	}
	rs.mu.Lock()
	rs.ReviewRepo.ChangeRatingAddReview(newReview, newReview.ID.ID)
	rs.mu.Unlock()
	return newReview, nil
}

func (rs *ReviewGRPCServer) DeleteReview(_ context.Context, in *review.DeleteReviewData) (*review.DeletedData, error) {
	rs.mu.RLock()
	rev, err := rs.ReviewRepo.GetUserReviewByID(in.ReviewID.ID, in.UserID.ID)
	rs.mu.RUnlock()
	if errors.Is(err, errorreview.ErrorNoReview) {
		return &review.DeletedData{
			IsDeleted: false,
		}, nil
	}
	if err != nil {
		return &review.DeletedData{
			IsDeleted: false,
		}, err
	}
	rs.mu.Lock()
	rs.ReviewRepo.ChangeRatingAfterDeleteReview(rev, in.ReviewID.ID)
	rs.mu.Unlock()
	rs.mu.Lock()
	isDeleted, err := rs.ReviewRepo.DeleteReviewRepo(rev.ID.ID)
	rs.mu.Unlock()
	if err != nil {
		return &review.DeletedData{
			IsDeleted: false,
		}, err
	}
	return &review.DeletedData{
		IsDeleted: isDeleted,
		Review:    rev,
	}, nil
}

func (rs *ReviewGRPCServer) UpdateReview(_ context.Context, in *review.UpdateReviewData) (*review.Review, error) {
	rs.mu.RLock()
	oldReview, err := rs.ReviewRepo.GetUserReviewByID(in.Review.ID.ID, in.UserID.ID)
	rs.mu.RUnlock()
	if errors.Is(err, errorreview.ErrorNoReview) {
		return &review.Review{}, nil
	}
	if err != nil {
		return &review.Review{}, err
	}
	rs.mu.Lock()
	updatedReview, err := rs.ReviewRepo.UpdateReviewRepo(in.Review)
	rs.mu.Unlock()
	if err != nil {
		return &review.Review{}, err
	}
	rs.mu.Lock()
	rs.ReviewRepo.ChangeRatingAfterUpdateReview(oldReview, updatedReview, in.Review.ID.ID)
	rs.mu.Unlock()
	return updatedReview, nil
}
