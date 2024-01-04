package reviewserviceusecse

import (
	"context"
	"errors"
	errorreview "kinopoisk/service_review/error"
	review "kinopoisk/service_review/proto"
	reviewservicerepo "kinopoisk/service_review/repo/mysql"
)

type ReviewGRPCServer struct {
	review.UnimplementedReviewMakerServer

	ReviewRepo reviewservicerepo.ReviewRepo
}

func NewReviewGRPCServer(reviewRepo reviewservicerepo.ReviewRepo) *ReviewGRPCServer {
	return &ReviewGRPCServer{
		UnimplementedReviewMakerServer: review.UnimplementedReviewMakerServer{},
		ReviewRepo:                     reviewRepo,
	}
}

func (rs *ReviewGRPCServer) GetFilmReviews(_ context.Context, in *review.FilmID) (*review.Reviews, error) {
	reviews, err := rs.ReviewRepo.GetFilmReviewsRepo(in.GetID())
	if err != nil {
		return nil, err
	}
	return &review.Reviews{
		Reviews: reviews,
	}, nil
}

func (rs *ReviewGRPCServer) NewReview(_ context.Context, in *review.NewReviewData) (*review.Review, error) {
	_, err := rs.ReviewRepo.GetReviewByFilmUser(in.GetFilmID().ID, in.GetUserID().ID)
	if errors.Is(err, errorreview.ErrorNoReview) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	newReview, err := rs.ReviewRepo.NewReviewRepo(in.GetReview(), in.GetFilmID().ID, in.GetUserID().ID)
	if err != nil {
		return nil, err
	}
	return newReview, nil
}

func (rs *ReviewGRPCServer) DeleteReview(_ context.Context, in *review.DeleteReviewData) (*review.IsDeleted, error) {
	id, err := rs.ReviewRepo.GetUserReviewByID(in.ReviewID.ID, in.UserID.ID)
	if errors.Is(err, errorreview.ErrorNoReview) {
		return &review.IsDeleted{
			IsDeleted: false,
		}, nil
	}
	if err != nil {
		return &review.IsDeleted{
			IsDeleted: false,
		}, err
	}
	isDeleted, err := rs.ReviewRepo.DeleteReviewRepo(id)
	if err != nil {
		return &review.IsDeleted{
			IsDeleted: false,
		}, err
	}
	return &review.IsDeleted{
		IsDeleted: isDeleted,
	}, nil
}

func (rs *ReviewGRPCServer) UpdateReview(_ context.Context, in *review.UpdateReviewData) (*review.Review, error) {
	_, err := rs.ReviewRepo.GetUserReviewByID(in.Review.ID.ID, in.UserID.ID)
	if errors.Is(err, errorreview.ErrorNoReview) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	updatedReview, err := rs.ReviewRepo.UpdateReviewRepo(in.Review)
	if err != nil {
		return nil, err
	}
	return updatedReview, nil
}
