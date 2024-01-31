package reviewserviceusecse

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	errorreview "kinopoisk/service_review/error"
	review "kinopoisk/service_review/proto"
	reviewservicerepo "kinopoisk/service_review/repo/mysql"
	"sync"
)

const (
	ChangeRatingQueueName = "change_rating"
)

type ReviewGRPCServer struct {
	review.UnimplementedReviewMakerServer
	ReviewRepo reviewservicerepo.ReviewRepo
	mu         *sync.RWMutex
	rabbitChan *amqp.Channel
	logger     *zap.SugaredLogger
}

type ChangeRatingInfo struct {
	ChangeType string
	ReviewID   uint64
	OldMark    uint32
	NewMark    uint32
	FilmID     uint64
}

func NewReviewGRPCServer(reviewRepo reviewservicerepo.ReviewRepo, rabbitChan *amqp.Channel, logger *zap.SugaredLogger) *ReviewGRPCServer {
	return &ReviewGRPCServer{
		UnimplementedReviewMakerServer: review.UnimplementedReviewMakerServer{},
		ReviewRepo:                     reviewRepo,
		mu:                             &sync.RWMutex{},
		rabbitChan:                     rabbitChan,
		logger:                         logger,
	}
}

func (rs *ReviewGRPCServer) GetFilmReviews(_ context.Context, in *review.FilmID) (*review.Reviews, error) {
	rs.mu.RLock()
	reviews, err := rs.ReviewRepo.GetFilmReviewsRepo(in.GetID())
	rs.mu.RUnlock()
	if err != nil {
		rs.logger.Errorf("error in getting film reviews: %s", err)
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
		rs.logger.Errorf("review from user %d for film %d already exists", in.GetUserID().ID, in.GetFilmID().ID)
		return &review.Review{}, nil
	}
	if err != nil {
		if !errors.Is(err, errorreview.ErrorNoReview) {
			rs.logger.Errorf("error in getting review by user and film: %s", err)
			return &review.Review{}, err
		}
	}
	rs.mu.Lock()
	newReview, err := rs.ReviewRepo.NewReviewRepo(in.GetReview(), in.GetFilmID().ID, in.GetUserID().ID)
	rs.mu.Unlock()
	if err != nil {
		rs.logger.Errorf("error in adding new review: %s", err)
		return &review.Review{}, err
	}
	changeRatingInfo := &ChangeRatingInfo{
		NewMark:  newReview.Mark,
		ReviewID: newReview.ID.ID,
	}
	err = rs.putChangeRatingTaskToQueue(changeRatingInfo)
	if err != nil {
		rs.logger.Errorf("error in changing rating: %s", err)
	}
	return newReview, nil
}

func (rs *ReviewGRPCServer) DeleteReview(_ context.Context, in *review.DeleteReviewData) (*review.DeletedData, error) {
	rs.mu.RLock()
	rev, err := rs.ReviewRepo.GetUserReviewByID(in.ReviewID.ID, in.UserID.ID)
	rs.mu.RUnlock()
	if errors.Is(err, errorreview.ErrorNoReview) {
		rs.logger.Errorf("no review with id %d and userID: %d", in.ReviewID.ID, in.UserID.ID)
		return &review.DeletedData{
			IsDeleted: false,
		}, nil
	}
	if err != nil {
		rs.logger.Errorf("error in getting review by id and user id: %s", err)
		return &review.DeletedData{
			IsDeleted: false,
		}, err
	}
	rs.mu.Lock()
	isDeleted, err := rs.ReviewRepo.DeleteReviewRepo(rev.ID.ID)
	rs.mu.Unlock()
	if err != nil {
		rs.logger.Errorf("error in deleting review: %s", err)
		return &review.DeletedData{
			IsDeleted: false,
		}, err
	}

	changeRatingInfo := &ChangeRatingInfo{
		OldMark: rev.Mark,
		FilmID:  rev.FilmID.ID,
	}
	err = rs.putChangeRatingTaskToQueue(changeRatingInfo)
	if err != nil {
		rs.logger.Errorf("error in changing rating: %s", err)
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
		rs.logger.Errorf("no review with id %d and userID %d", in.Review.ID.ID, in.UserID.ID)
		return &review.Review{}, nil
	}
	if err != nil {
		rs.logger.Errorf("error in getting review by id and user id: %s", err)
		return &review.Review{}, err
	}
	rs.mu.Lock()
	updatedReview, err := rs.ReviewRepo.UpdateReviewRepo(in.Review)
	rs.mu.Unlock()
	if err != nil {
		rs.logger.Errorf("error in updating review: %s", err)
		return &review.Review{}, err
	}
	changeRatingInfo := &ChangeRatingInfo{
		OldMark:  oldReview.Mark,
		NewMark:  updatedReview.Mark,
		ReviewID: in.Review.ID.ID,
	}
	err = rs.putChangeRatingTaskToQueue(changeRatingInfo)
	if err != nil {
		rs.logger.Errorf("error in changing rating: %s", err)
	}
	return updatedReview, nil
}

func (rs *ReviewGRPCServer) putChangeRatingTaskToQueue(chInfo *ChangeRatingInfo) error {
	data, err := json.Marshal(chInfo)
	if err != nil {
		return err
	}
	err = rs.rabbitChan.Publish(
		"",
		ChangeRatingQueueName,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         data,
		})
	return err
}
