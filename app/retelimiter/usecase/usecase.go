package ratelimiterusecase

type RateLimiterUseCase interface {
	CheckRateLimit(userAddr string) bool
}
