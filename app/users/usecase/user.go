package userusecase

import (
	"kinopoisk/app/entity"
	auth "kinopoisk/service_auth/proto"
	"sync"
)

type UserUseCase interface {
	Login(username, password string) (*entity.User, error)
	Register(username, password string) (*entity.User, error)
	CreateSession(user *entity.User) (string, error)
	GetSession(token string) (*entity.Session, error)
	DeleteSession(token string) (bool, error)
}

type AuthGRPCClient struct {
	mu         *sync.RWMutex
	grpcClient auth.AuthMakerClient
}

func NewAuthGRPCClient(grpcClient auth.AuthMakerClient) *AuthGRPCClient {
	return &AuthGRPCClient{
		mu:         &sync.RWMutex{},
		grpcClient: grpcClient,
	}
}

func (a *AuthGRPCClient) Login(username, password string) (*entity.User, error) {
	return nil, nil
}

func (a *AuthGRPCClient) Register(username, password string) (*entity.User, error) {
	return nil, nil
}

func (a *AuthGRPCClient) CreateSession(user *entity.User) (string, error) {
	return "", nil
}

func (a *AuthGRPCClient) GetSession(token string) (*entity.Session, error) {
	return nil, nil
}

func (a *AuthGRPCClient) DeleteSession(token string) (bool, error) {
	return false, nil
}
