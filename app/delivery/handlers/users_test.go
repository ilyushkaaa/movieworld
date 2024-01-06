package handlers

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"go.uber.org/zap"
	"io"
	"kinopoisk/app/entity"
	userusecase "kinopoisk/app/users/usecase"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testUseCase := userusecase.NewMockUserUseCase(ctrl)
	testHandler := NewUserHandler(testUseCase, zap.NewNop().Sugar())

	// can not read request body
	request := httptest.NewRequest(http.MethodPost, "/login", &errorReader{})
	respWriter := httptest.NewRecorder()
	testHandler.Login(respWriter, request)
	resp := respWriter.Result()
	_, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("unable to read response body")
		return
	}
	err = resp.Body.Close()
	if err != nil {
		t.Fatalf("failed to close response body")
	}
	if resp.StatusCode != 401 {
		t.Errorf("expected status %d, got status %d", http.StatusUnauthorized, resp.StatusCode)
	}

	request = httptest.NewRequest(http.MethodPost, "/login", strings.NewReader("{"))
	respWriter = httptest.NewRecorder()
	testHandler.Login(respWriter, request)
	resp = respWriter.Result()
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("unable to read response body")
		return
	}
	err = resp.Body.Close()
	if err != nil {
		t.Fatalf("failed to close response body")
	}
	if resp.StatusCode != 401 {
		t.Errorf("expected status %d, got status %d", http.StatusUnauthorized, resp.StatusCode)
	}

	request = httptest.NewRequest(http.MethodPost, "/review/1", strings.NewReader(`{"username":"hello12"}`))
	respWriter = httptest.NewRecorder()
	testHandler.Login(respWriter, request)
	resp = respWriter.Result()
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("unable to read response body")
		return
	}
	err = resp.Body.Close()
	if err != nil {
		t.Fatalf("failed to close response body")
	}
	if resp.StatusCode != 401 {
		t.Errorf("expected status %d, got status %d", http.StatusUnauthorized, resp.StatusCode)
	}

	testUseCase.EXPECT().Login("hello12", "qqqqqqqqq").Return(nil, fmt.Errorf("internal server error"))
	request = httptest.NewRequest(http.MethodPost, "/review/1", strings.NewReader(`{"username":"hello12","password":"qqqqqqqqq"}`))
	respWriter = httptest.NewRecorder()
	testHandler.Login(respWriter, request)
	resp = respWriter.Result()
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("unable to read response body")
		return
	}
	err = resp.Body.Close()
	if err != nil {
		t.Fatalf("failed to close response body")
	}
	if resp.StatusCode != 500 {
		t.Errorf("expected status %d, got status %d", http.StatusInternalServerError, resp.StatusCode)
	}

	testUseCase.EXPECT().Login("hello12", "qqqqqqqqq").Return(nil, nil)
	request = httptest.NewRequest(http.MethodPost, "/review/1", strings.NewReader(`{"username":"hello12","password":"qqqqqqqqq"}`))
	respWriter = httptest.NewRecorder()
	testHandler.Login(respWriter, request)
	resp = respWriter.Result()
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("unable to read response body")
		return
	}
	err = resp.Body.Close()
	if err != nil {
		t.Fatalf("failed to close response body")
	}
	if resp.StatusCode != 401 {
		t.Errorf("expected status %d, got status %d", http.StatusUnauthorized, resp.StatusCode)
	}

	loggedInUser := &entity.User{
		ID:       1,
		Username: "some_username",
	}
	testUseCase.EXPECT().Login("some_username", "aaaaaaaa").Return(loggedInUser, nil)
	testUseCase.EXPECT().CreateSession(loggedInUser).Return("", fmt.Errorf("error"))
	request = httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"username":"some_username", "password":"aaaaaaaa"}`))
	respWriter = httptest.NewRecorder()
	testHandler.Login(respWriter, request)
	resp = respWriter.Result()
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("unable to read response body")
		return
	}
	err = resp.Body.Close()
	if err != nil {
		t.Fatalf("failed to close response body")
	}
	if resp.StatusCode != 500 {
		t.Errorf("expected status %d, got status %d", http.StatusInternalServerError, resp.StatusCode)
	}

	testUseCase.EXPECT().Login("some_username", "aaaaaaaa").Return(loggedInUser, nil)
	testUseCase.EXPECT().CreateSession(loggedInUser).Return("some_token", nil)
	request = httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"username":"some_username", "password":"aaaaaaaa"}`))
	respWriter = httptest.NewRecorder()
	testHandler.Login(respWriter, request)
	resp = respWriter.Result()
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("unable to read response body")
		return
	}
	err = resp.Body.Close()
	if err != nil {
		t.Fatalf("failed to close response body")
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected status %d, got status %d", http.StatusOK, resp.StatusCode)
	}

}