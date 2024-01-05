package ratelimiterusecase

import (
	ratelimiterrepo "kinopoisk/app/ratelimiter/repo/redis"
	"sync"
	"time"
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

func (rl *RateLimiterUseCaseStruct) CheckRateLimit(userAddr string) bool {
	currentTimeMillis := time.Now().UnixNano() / int64(time.Millisecond)
	rl.mu.RLock()
	numOfRequests := rl.RateLimiterRepo.CheckRateLimitRepo(userAddr, currentTimeMillis-2000)
	rl.mu.RUnlock()
	var canMakeRequest = false
	if numOfRequests < 3 {
		canMakeRequest = true
		rl.RateLimiterRepo.AddRateRepo(userAddr, currentTimeMillis)
	}
	return canMakeRequest
}
