package handlers

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"go.uber.org/zap"
	"io"
	actorusecase "kinopoisk/app/actors/usecase"
	"kinopoisk/app/entity"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetActors(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testUseCase := actorusecase.NewMockActorUseCase(ctrl)
	testHandler := &ActorHandler{
		Logger:        zap.NewNop().Sugar(),
		ActorUseCases: testUseCase,
	}
	// usecase returns error
	testUseCase.EXPECT().GetActors().Return(nil, fmt.Errorf("error"))
	request := httptest.NewRequest(http.MethodGet, "/actors/", nil)
	respWriter := httptest.NewRecorder()
	testHandler.GetActors(respWriter, request)
	resp := respWriter.Result()
	_, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("unable to read response body")
		return
	}
	if resp.StatusCode != 500 {
		t.Errorf("expected status %d, got status %d", http.StatusInternalServerError, resp.StatusCode)
		return
	}

	//usecase returns actors without error
	actors := []*entity.Actor{
		{
			ID:          1,
			Name:        "Sergey",
			Surname:     "Burunov",
			Nationality: "Russian",
			Birthday:    time.Now(),
		},
	}
	testUseCase.EXPECT().GetActors().Return(actors, nil)
	request = httptest.NewRequest(http.MethodGet, "/actors/", nil)
	respWriter = httptest.NewRecorder()
	testHandler.GetActors(respWriter, request)
	resp = respWriter.Result()
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("unable to read response body")
		return
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected status %d, got status %d", http.StatusOK, resp.StatusCode)
		return
	}

}
