package ratelimiterrepo

type RateLimiterRepo interface {
	CheckRateLimitRepo(userAddr string) bool
}
