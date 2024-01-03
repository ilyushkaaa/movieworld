package ratelimiterrepo

import "github.com/gomodule/redigo/redis"

type RateLimiterRepo interface {
	CheckRateLimitRepo(userAddr string) bool
}

type RateLimiterRepoRedis struct {
	RedisConn redis.Conn
}

func NewRateLimiterRepoRedis(redisConn redis.Conn) *RateLimiterRepoRedis {
	return &RateLimiterRepoRedis{
		RedisConn: redisConn,
	}
}
