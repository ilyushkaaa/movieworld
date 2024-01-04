package sessionrepo

import (
	"github.com/gomodule/redigo/redis"
	auth "kinopoisk/service_auth/proto"
)

type SessionRepo interface {
	CreateSessionRepo(session *auth.Session) error
	GetSessionRepo(token string) (*auth.Session, error)
	DeleteSessionRepo(token string) (bool, error)
}

type SessionRepoRedis struct {
	redisConn redis.Conn
}

func NewSessionRepoRedis(redisConn redis.Conn) *SessionRepoRedis {
	return &SessionRepoRedis{
		redisConn: redisConn,
	}
}

func (s *SessionRepoRedis) CreateSessionRepo(session *auth.Session) error {

}

func (s *SessionRepoRedis) GetSessionRepo(token string) (*auth.Session, error) {

}

func (s *SessionRepoRedis) DeleteSessionRepo(token string) (bool, error) {

}
