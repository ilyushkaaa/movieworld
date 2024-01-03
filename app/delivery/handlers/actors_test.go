package handlers

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
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
	testHandler := NewActorHandler(testUseCase, zap.NewNop().Sugar())
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

func TestGetActorByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testUseCase := actorusecase.NewMockActorUseCase(ctrl)
	testHandler := NewActorHandler(testUseCase, zap.NewNop().Sugar())
	// bad id format
	request := httptest.NewRequest(http.MethodGet, "/actor/bad", nil)
	request = mux.SetURLVars(request, map[string]string{"ACTOR_ID": "bad"})
	respWriter := httptest.NewRecorder()
	testHandler.GetActorByID(respWriter, request)
	resp := respWriter.Result()
	_, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("unable to read response body")
		return
	}
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status %d, got status %d", http.StatusBadRequest, resp.StatusCode)
		return
	}

	// usecase returns error
	var ID uint64 = 1
	testUseCase.EXPECT().GetActorByID(ID).Return(nil, fmt.Errorf("error"))
	request = httptest.NewRequest(http.MethodGet, "/actor/1", nil)
	request = mux.SetURLVars(request, map[string]string{"ACTOR_ID": "1"})

	respWriter = httptest.NewRecorder()
	testHandler.GetActorByID(respWriter, request)
	resp = respWriter.Result()
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("unable to read response body")
		return
	}
	if resp.StatusCode != 500 {
		t.Errorf("expected status %d, got status %d", http.StatusInternalServerError, resp.StatusCode)
		return
	}

	// usecase returns nil actor

	testUseCase.EXPECT().GetActorByID(ID).Return(nil, nil)
	request = httptest.NewRequest(http.MethodGet, "/actor/1", nil)
	request = mux.SetURLVars(request, map[string]string{"ACTOR_ID": "1"})
	respWriter = httptest.NewRecorder()
	testHandler.GetActorByID(respWriter, request)
	resp = respWriter.Result()
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("unable to read response body")
		return
	}
	if resp.StatusCode != 404 {
		t.Errorf("expected status %d, got status %d", http.StatusNotFound, resp.StatusCode)
		return
	}

	// all is ok
	actor := &entity.Actor{
		ID:          1,
		Name:        "Sergey",
		Surname:     "Burunov",
		Nationality: "Russian",
		Birthday:    time.Now(),
	}
	testUseCase.EXPECT().GetActorByID(ID).Return(actor, nil)
	request = httptest.NewRequest(http.MethodGet, "/actor/1", nil)
	request = mux.SetURLVars(request, map[string]string{"ACTOR_ID": "1"})
	respWriter = httptest.NewRecorder()
	testHandler.GetActorByID(respWriter, request)
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
