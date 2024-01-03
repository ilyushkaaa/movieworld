package ratelimiterusecase

import (
	ratelimiterrepo "kinopoisk/app/ratelimiter/repo/redis"
	"sync"
)

type RateLimiterUseCase interface {
	CheckRateLimit(userAddr string) bool
}

type RateLimiterUseCaseStruct struct {
	mu              *sync.RWMutex
	RateLimiterRepo ratelimiterrepo.RateLimiterRepo
}

func NewFilmUseCaseStruct(repo ratelimiterrepo.RateLimiterRepo) *RateLimiterUseCaseStruct {
	return &RateLimiterUseCaseStruct{
		mu:              &sync.RWMutex{},
		RateLimiterRepo: repo,
	}
}
