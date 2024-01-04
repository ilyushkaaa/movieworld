package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"io"
	"kinopoisk/app/delivery"
	"kinopoisk/app/dto"
	"kinopoisk/app/entity"
	errorapp "kinopoisk/app/errors"
	"kinopoisk/app/middleware"
	userusecase "kinopoisk/app/users/usecase"
	"net/http"
)

type UserHandler struct {
	UserUseCases userusecase.UserUseCase
	Logger       *zap.SugaredLogger
}

func NewUserHandler(userUseCases userusecase.UserUseCase, logger *zap.SugaredLogger) *UserHandler {
	return &UserHandler{
		UserUseCases: userUseCases,
		Logger:       logger,
	}
}

func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	userFromLoginForm, err := checkRequestFormat(uh.Logger, w, r)
	if err != nil || userFromLoginForm == nil {
		errText := fmt.Sprintf(`{"message": "error in login request: %s"}`, err)
		delivery.WriteResponse(uh.Logger, w, []byte(errText), http.StatusUnauthorized)
		return
	}
	loggedInUser, err := uh.UserUseCases.Login(userFromLoginForm.Username, userFromLoginForm.Password)

	if errors.Is(err, errorapp.ErrorNoUser) {
		delivery.WriteResponse(uh.Logger, w, []byte(`{"message": "user not found"}`), http.StatusUnauthorized)
		return
	}
	if errors.Is(err, errorapp.ErrorBadPassword) {
		delivery.WriteResponse(uh.Logger, w, []byte(`{"message": "invalid password"}`), http.StatusUnauthorized)
		return
	}
	if err != nil {
		errText := fmt.Sprintf(`{"message": "error in getting user by login and password: %s"}`, err)
		delivery.WriteResponse(uh.Logger, w, []byte(errText), http.StatusInternalServerError)
		return
	}
	uh.HandleGetToken(w, loggedInUser)

}

func (uh *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	userFromLoginForm, err := checkRequestFormat(uh.Logger, w, r)
	if err != nil || userFromLoginForm == nil {
		return
	}
	newUser, err := uh.UserUseCases.Register(userFromLoginForm.Username, userFromLoginForm.Password)

	if errors.Is(err, errorapp.ErrorUserExists) {
		delivery.WriteResponse(uh.Logger, w, []byte(`{"message": "user already exists"}`), http.StatusUnprocessableEntity)
		return
	}
	if err != nil {
		errText := fmt.Sprintf(`{"message": "unknown error occured in register: %s"}`, err)
		delivery.WriteResponse(uh.Logger, w, []byte(errText), http.StatusInternalServerError)
		return
	}
	uh.HandleGetToken(w, newUser)
}

func (uh *UserHandler) HandleGetToken(w http.ResponseWriter, newUser *entity.User) {
	token, err := uh.UserUseCases.CreateSession(newUser)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "error in session creation: %s"}`, err)
		delivery.WriteResponse(uh.Logger, w, []byte(errText), http.StatusInternalServerError)
		return
	}
	resp := dto.AuthResponseDTO{
		Token: token,
	}
	tokenJSON, err := json.Marshal(&resp)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "error in coding response: %s"}`, err)
		delivery.WriteResponse(uh.Logger, w, []byte(errText), http.StatusInternalServerError)
		return
	}
	uh.Logger.Infof("new token: %s", token)
	delivery.WriteResponse(uh.Logger, w, tokenJSON, http.StatusOK)
}

func checkRequestFormat(logger *zap.SugaredLogger, w http.ResponseWriter, r *http.Request) (*dto.AuthRequestDTO, error) {
	rBody, err := io.ReadAll(r.Body)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "error in reading request body: %s"}`, err)
		delivery.WriteResponse(logger, w, []byte(errText), http.StatusBadRequest)
		return nil, err
	}
	userFromLoginForm := &dto.AuthRequestDTO{}
	err = json.Unmarshal(rBody, userFromLoginForm)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "error in decoding user: %s"}`, err)
		delivery.WriteResponse(logger, w, []byte(errText), http.StatusInternalServerError)
		return nil, err
	}
	if validationErrors := userFromLoginForm.Validate(); len(validationErrors) != 0 {
		errorsJSON, err := json.Marshal(validationErrors)
		if err != nil {
			errText := fmt.Sprintf(`{"message": "error in decoding validation errors: %s"}`, err)
			delivery.WriteResponse(logger, w, []byte(errText), http.StatusInternalServerError)
			return nil, err
		}
		logger.Errorf("login form did not pass validation: %s", err)
		delivery.WriteResponse(logger, w, errorsJSON, http.StatusUnprocessableEntity)
		return nil, err
	}
	return userFromLoginForm, nil
}

func (uh *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	token, ok := ctx.Value(middleware.MyTokenKey).(string)
	if !ok {
		delivery.WriteResponse(uh.Logger, w, []byte(`{"message": "can not cast context value to user"}`), http.StatusInternalServerError)
		return
	}
	isDeleted, err := uh.UserUseCases.DeleteSession(token)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "error in logging out: %s"}`, err)
		delivery.WriteResponse(uh.Logger, w, []byte(errText), http.StatusInternalServerError)
		return
	}
	if !isDeleted {
		errText := fmt.Sprintf(`{"message": "no session with token: %s"}`, token)
		delivery.WriteResponse(uh.Logger, w, []byte(errText), http.StatusNotFound)
		return
	}
	message := `{"result":"success"}`
	delivery.WriteResponse(uh.Logger, w, []byte(message), http.StatusOK)
}
