// Code generated by MockGen. DO NOT EDIT.
// Source: ./review.go

// Package reviewusecase is a generated GoMock package.
package reviewusecase

import (
	dto "kinopoisk/app/dto"
	entity "kinopoisk/app/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockReviewUseCase is a mock of ReviewUseCase interface.
type MockReviewUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockReviewUseCaseMockRecorder
}

// MockReviewUseCaseMockRecorder is the mock recorder for MockReviewUseCase.
type MockReviewUseCaseMockRecorder struct {
	mock *MockReviewUseCase
}

// NewMockReviewUseCase creates a new mock instance.
func NewMockReviewUseCase(ctrl *gomock.Controller) *MockReviewUseCase {
	mock := &MockReviewUseCase{ctrl: ctrl}
	mock.recorder = &MockReviewUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReviewUseCase) EXPECT() *MockReviewUseCaseMockRecorder {
	return m.recorder
}

// DeleteReview mocks base method.
func (m *MockReviewUseCase) DeleteReview(reviewID, userID uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteReview", reviewID, userID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteReview indicates an expected call of DeleteReview.
func (mr *MockReviewUseCaseMockRecorder) DeleteReview(reviewID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteReview", reflect.TypeOf((*MockReviewUseCase)(nil).DeleteReview), reviewID, userID)
}

// GetFilmReviews mocks base method.
func (m *MockReviewUseCase) GetFilmReviews(filmID uint64) ([]*entity.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilmReviews", filmID)
	ret0, _ := ret[0].([]*entity.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFilmReviews indicates an expected call of GetFilmReviews.
func (mr *MockReviewUseCaseMockRecorder) GetFilmReviews(filmID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilmReviews", reflect.TypeOf((*MockReviewUseCase)(nil).GetFilmReviews), filmID)
}

// NewReview mocks base method.
func (m *MockReviewUseCase) NewReview(newReview *dto.ReviewDTO, filmID uint64, user *entity.User) (*entity.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewReview", newReview, filmID, user)
	ret0, _ := ret[0].(*entity.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewReview indicates an expected call of NewReview.
func (mr *MockReviewUseCaseMockRecorder) NewReview(newReview, filmID, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewReview", reflect.TypeOf((*MockReviewUseCase)(nil).NewReview), newReview, filmID, user)
}

// UpdateReview mocks base method.
func (m *MockReviewUseCase) UpdateReview(reviewToUpdate *dto.ReviewDTO, reviewID uint64, user *entity.User) (*entity.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateReview", reviewToUpdate, reviewID, user)
	ret0, _ := ret[0].(*entity.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateReview indicates an expected call of UpdateReview.
func (mr *MockReviewUseCaseMockRecorder) UpdateReview(reviewToUpdate, reviewID, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateReview", reflect.TypeOf((*MockReviewUseCase)(nil).UpdateReview), reviewToUpdate, reviewID, user)
}
