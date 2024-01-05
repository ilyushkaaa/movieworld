package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"io"
	"kinopoisk/app/delivery"
	"kinopoisk/app/dto"
	"kinopoisk/app/entity"
	errorapp "kinopoisk/app/errors"
	"kinopoisk/app/middleware"
	reviewusecase "kinopoisk/app/reviews/usecase"
	"net/http"
	"strconv"
)

type ReviewHandler struct {
	ReviewUseCases reviewusecase.ReviewUseCase
	Logger         *zap.SugaredLogger
}

func NewReviewHandler(reviewUseCases reviewusecase.ReviewUseCase, logger *zap.SugaredLogger) *ReviewHandler {
	return &ReviewHandler{
		ReviewUseCases: reviewUseCases,
		Logger:         logger,
	}
}

func (rh *ReviewHandler) GetReviewsForFilm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filmID := vars["FILM_ID"]
	filmIDInt, err := strconv.ParseUint(filmID, 10, 64)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "bad format of film id: %s"}`, err)
		delivery.WriteResponse(rh.Logger, w, []byte(errText), http.StatusBadRequest)
		return
	}
	reviews, err := rh.ReviewUseCases.GetFilmReviews(filmIDInt)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "internal server error: %s"}`, err)
		delivery.WriteResponse(rh.Logger, w, []byte(errText), http.StatusInternalServerError)
		return
	}
	reviewsJSON, err := json.Marshal(reviews)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "error in coding reviews: %s"}`, err)
		delivery.WriteResponse(rh.Logger, w, []byte(errText), http.StatusInternalServerError)
		return
	}
	delivery.WriteResponse(rh.Logger, w, reviewsJSON, http.StatusOK)
}

func (rh *ReviewHandler) AddReview(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filmID := vars["FILM_ID"]
	filmIDInt, err := strconv.ParseUint(filmID, 10, 64)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "bad format of film id: %s"}`, err)
		delivery.WriteResponse(rh.Logger, w, []byte(errText), http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	user, ok := ctx.Value(middleware.MyUserKey).(*entity.User)
	if !ok {
		delivery.WriteResponse(rh.Logger, w, []byte(`{"message": "can not cast context value to user"}`), http.StatusInternalServerError)
		return
	}
	reviewDTO := &dto.ReviewDTO{}
	rBody, err := io.ReadAll(r.Body)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "error in reading request body: %s"}`, err)
		delivery.WriteResponse(rh.Logger, w, []byte(errText), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(rBody, reviewDTO)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "error in decoding posts: %s"}`, err)
		delivery.WriteResponse(rh.Logger, w, []byte(errText), http.StatusBadRequest)
		return
	}

	if validationErrors := reviewDTO.Validate(); len(validationErrors) != 0 {
		var errorsJSON []byte
		errorsJSON, err = json.Marshal(validationErrors)
		if err != nil {
			errText := fmt.Sprintf(`{"message": "error in json decoding: %s"}`, err)
			delivery.WriteResponse(rh.Logger, w, []byte(errText), http.StatusInternalServerError)
			return
		}
		delivery.WriteResponse(rh.Logger, w, errorsJSON, http.StatusUnprocessableEntity)
		return
	}
	addedReview, err := rh.ReviewUseCases.NewReview(reviewDTO, filmIDInt, user)
	if errors.Is(err, errorapp.ErrorNoFilm) {
		errText := fmt.Sprintf(`{"message": "no film with id: %d"}`, filmIDInt)
		delivery.WriteResponse(rh.Logger, w, []byte(errText), http.StatusNotFound)
		return
	}
	if err != nil {
		errText := fmt.Sprintf(`{"message": "internal server error: %s"}`, err)
		delivery.WriteResponse(rh.Logger, w, []byte(errText), http.StatusInternalServerError)
		return
	}
	if addedReview == nil {
		errText := fmt.Sprintf(`{"message": "film with id %d has been already reviewed"}`, filmIDInt)
		delivery.WriteResponse(rh.Logger, w, []byte(errText), http.StatusBadRequest)
		return
	}
	reviewJSON, err := json.Marshal(addedReview)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "error in coding reviews: %s"}`, err)
		delivery.WriteResponse(rh.Logger, w, []byte(errText), http.StatusInternalServerError)
		return
	}
	delivery.WriteResponse(rh.Logger, w, reviewJSON, http.StatusOK)
}

func (rh *ReviewHandler) DeleteReview(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, ok := ctx.Value(middleware.MyUserKey).(*entity.User)
	if !ok {
		delivery.WriteResponse(rh.Logger, w, []byte(`{"message": "can not cast context value to user"}`), http.StatusInternalServerError)
		return
	}
	vars := mux.Vars(r)
	reviewID := vars["REVIEW_ID"]
	reviewIDInt, err := strconv.ParseUint(reviewID, 10, 64)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "bad format of actor id: %s"}`, err)
		delivery.WriteResponse(rh.Logger, w, []byte(errText), http.StatusBadRequest)
		return
	}
	wasDeleted, err := rh.ReviewUseCases.DeleteReview(reviewIDInt, user.ID)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "internal server error: %s"}`, err)
		delivery.WriteResponse(rh.Logger, w, []byte(errText), http.StatusInternalServerError)
		return
	}
	if !wasDeleted {
		result := fmt.Sprintf(`{"meassage": "not found"}`)
		delivery.WriteResponse(rh.Logger, w, []byte(result), http.StatusNotFound)
		return
	}
	result := fmt.Sprintf(`{"result": "success"}`)
	delivery.WriteResponse(rh.Logger, w, []byte(result), http.StatusOK)
}

func (rh *ReviewHandler) UpdateReview(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reviewID := vars["REVIEW_ID"]
	reviewIDInt, err := strconv.ParseUint(reviewID, 10, 64)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "bad format of film id: %s"}`, err)
		delivery.WriteResponse(rh.Logger, w, []byte(errText), http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	user, ok := ctx.Value(middleware.MyUserKey).(*entity.User)
	if !ok {
		delivery.WriteResponse(rh.Logger, w, []byte(`{"message": "can not cast context value to user"}`), http.StatusInternalServerError)
		return
	}
	reviewDTO := &dto.ReviewDTO{}
	rBody, err := io.ReadAll(r.Body)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "error in reading request body: %s"}`, err)
		delivery.WriteResponse(rh.Logger, w, []byte(errText), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(rBody, reviewDTO)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "error in decoding posts: %s"}`, err)
		delivery.WriteResponse(rh.Logger, w, []byte(errText), http.StatusBadRequest)
		return
	}
	if validationErrors := reviewDTO.Validate(); len(validationErrors) != 0 {
		var errorsJSON []byte
		errorsJSON, err = json.Marshal(validationErrors)
		if err != nil {
			errText := fmt.Sprintf(`{"message": "error in json decoding: %s"}`, err)
			delivery.WriteResponse(rh.Logger, w, []byte(errText), http.StatusInternalServerError)
			return
		}
		delivery.WriteResponse(rh.Logger, w, errorsJSON, http.StatusUnprocessableEntity)
		return
	}
	updatedReview, err := rh.ReviewUseCases.UpdateReview(reviewDTO, reviewIDInt, user)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "internal server error: %s"}`, err)
		delivery.WriteResponse(rh.Logger, w, []byte(errText), http.StatusInternalServerError)
		return
	}
	if updatedReview == nil {
		errText := fmt.Sprintf(`{"message": "review with id %d is not found"}`, reviewIDInt)
		delivery.WriteResponse(rh.Logger, w, []byte(errText), http.StatusNotFound)
		return
	}
	reviewJSON, err := json.Marshal(updatedReview)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "error in coding review: %s"}`, err)
		delivery.WriteResponse(rh.Logger, w, []byte(errText), http.StatusInternalServerError)
		return
	}
	delivery.WriteResponse(rh.Logger, w, reviewJSON, http.StatusOK)
}
