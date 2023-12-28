package middleware

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"kinopoisk/app/delivery"
	userusecase "kinopoisk/app/users/usecase"
	"net/http"
	"strings"
)

type userKey int
type tokenKey int

const (
	MyUserKey  userKey  = 1
	MyTokenKey tokenKey = 2
)

func AuthMiddleware(logger *zap.SugaredLogger, uc userusecase.UserUseCase, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("auth middleware start")
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			delivery.WriteResponse(logger, w, []byte(`{"message": "there is no access token or it is in wrong format"}`),
				http.StatusUnauthorized)
			return
		}
		tokenValue := strings.TrimPrefix(authHeader, "Bearer ")
		mySession, err := uc.GetSession(tokenValue)
		if err != nil || mySession == nil {
			errText := fmt.Sprintf(`{"message": "there is no session for token %s}`, tokenValue)
			delivery.WriteResponse(logger, w, []byte(errText), http.StatusUnauthorized)
			return
		}
		sessionUser := mySession.User
		ctx := r.Context()
		ctx = context.WithValue(ctx, MyUserKey, sessionUser)
		ctx = context.WithValue(ctx, MyTokenKey, tokenValue)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
