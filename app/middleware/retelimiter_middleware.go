package middleware

import (
	"go.uber.org/zap"
	ratelimiterusecase "kinopoisk/app/retelimiter/usecase"
	"net/http"
)

func RateLimiterMiddleware(logger *zap.SugaredLogger, rateLimiterUseCases ratelimiterusecase.RateLimiterUseCase, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestAddr := r.RemoteAddr
		canMakeRequest := rateLimiterUseCases.CheckRateLimit(requestAddr)
		if !canMakeRequest {
			return
		}
		next.ServeHTTP(w, r)
	})
}
