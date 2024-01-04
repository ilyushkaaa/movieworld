package userusecase

import (
	"context"
	"kinopoisk/app/entity"
	errorapp "kinopoisk/app/errors"
	auth "kinopoisk/service_auth/proto"
)

type UserUseCase interface {
	Login(username, password string) (*entity.User, error)
	Register(username, password string) (*entity.User, error)
	CreateSession(user *entity.User) (string, error)
	GetSession(token string) (*entity.Session, error)
	DeleteSession(token string) (bool, error)
}

type AuthGRPCClient struct {
	grpcClient auth.AuthMakerClient
}

func NewAuthGRPCClient(grpcClient auth.AuthMakerClient) *AuthGRPCClient {
	return &AuthGRPCClient{
		grpcClient: grpcClient,
	}
}

func (a *AuthGRPCClient) Login(username, password string) (*entity.User, error) {
	loggedInUser, err := a.grpcClient.Login(context.Background(), &auth.AuthData{
		Username: username,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	if loggedInUser == nil {
		return nil, nil
	}
	newUserApp := getUserFromGRPCStruct(loggedInUser)
	return newUserApp, nil
}

func (a *AuthGRPCClient) Register(username, password string) (*entity.User, error) {
	newUser, err := a.grpcClient.Login(context.Background(), &auth.AuthData{
		Username: username,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	if newUser == nil {
		return nil, errorapp.ErrorUserExists
	}
	newUserApp := getUserFromGRPCStruct(newUser)
	return newUserApp, nil
}

func (a *AuthGRPCClient) CreateSession(user *entity.User) (string, error) {
	userGRPC := getGRPCUserFromEntityUser(user)
	token, err := a.grpcClient.CreateSession(context.Background(), userGRPC)
	if err != nil {
		return "", err
	}
	return token.Token, nil
}

func (a *AuthGRPCClient) GetSession(token string) (*entity.Session, error) {
	session, err := a.grpcClient.GetSession(context.Background(), &auth.Token{
		Token: token,
	})
	if err != nil {
		return nil, err
	}
	sessionApp := getSessionFromGRPCStruct(session)
	return sessionApp, nil
}

func (a *AuthGRPCClient) DeleteSession(token string) (bool, error) {
	isDeleted, err := a.grpcClient.DeleteSession(context.Background(), &auth.Token{
		Token: token,
	})
	if err != nil {
		return false, err
	}
	if !isDeleted.IsDeleted {
		return false, errorapp.ErrorNoSession
	}
	return true, nil
}

func getUserFromGRPCStruct(user *auth.User) *entity.User {
	return &entity.User{
		ID:       user.ID,
		Username: user.Username,
	}
}

func getGRPCUserFromEntityUser(user *entity.User) *auth.User {
	return &auth.User{
		ID:       user.ID,
		Username: user.Username,
	}
}

func getSessionFromGRPCStruct(sess *auth.Session) *entity.Session {
	return &entity.Session{
		ID:   sess.ID,
		User: getUserFromGRPCStruct(sess.User),
	}
}
