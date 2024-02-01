package userusecase

import (
	"context"
	"go.uber.org/zap"
	"kinopoisk/app/entity"
	errorapp "kinopoisk/app/errors"
	auth "kinopoisk/service_auth/proto"
)

type UserUseCase interface {
	Login(username, password string, logger *zap.SugaredLogger) (*entity.User, error)
	Register(username, password string, logger *zap.SugaredLogger) (*entity.User, error)
	CreateSession(user *entity.User, logger *zap.SugaredLogger) (string, error)
	GetSession(token string, logger *zap.SugaredLogger) (*entity.Session, error)
	DeleteSession(token string, logger *zap.SugaredLogger) (bool, error)
}

type AuthGRPCClient struct {
	grpcClient auth.AuthMakerClient
}

type loggerKey int

const MyLoggerKey loggerKey = 3

func NewAuthGRPCClient(grpcClient auth.AuthMakerClient) *AuthGRPCClient {
	return &AuthGRPCClient{
		grpcClient: grpcClient,
	}
}

func (a *AuthGRPCClient) Login(username, password string, logger *zap.SugaredLogger) (*entity.User, error) {
	ctx := context.WithValue(context.Background(), MyLoggerKey, logger)
	loggedInUser, err := a.grpcClient.Login(ctx, &auth.AuthData{
		Username: username,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	if loggedInUser.ID == 0 {
		return nil, nil
	}
	newUserApp := getUserFromGRPCStruct(loggedInUser)
	return newUserApp, nil
}

func (a *AuthGRPCClient) Register(username, password string, logger *zap.SugaredLogger) (*entity.User, error) {
	ctx := context.WithValue(context.Background(), MyLoggerKey, logger)
	newUser, err := a.grpcClient.Register(ctx, &auth.AuthData{
		Username: username,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	if newUser.ID == 0 {
		return nil, errorapp.ErrorUserExists
	}
	newUserApp := getUserFromGRPCStruct(newUser)
	return newUserApp, nil
}

func (a *AuthGRPCClient) CreateSession(user *entity.User, logger *zap.SugaredLogger) (string, error) {
	userGRPC := getGRPCUserFromEntityUser(user)
	ctx := context.WithValue(context.Background(), MyLoggerKey, logger)
	token, err := a.grpcClient.CreateSession(ctx, userGRPC)
	if err != nil {
		return "", err
	}
	return token.Token, nil
}

func (a *AuthGRPCClient) GetSession(token string, logger *zap.SugaredLogger) (*entity.Session, error) {
	ctx := context.WithValue(context.Background(), MyLoggerKey, logger)
	session, err := a.grpcClient.GetSession(ctx, &auth.Token{
		Token: token,
	})
	if err != nil {
		return nil, err
	}
	sessionApp := getSessionFromGRPCStruct(session)
	return sessionApp, nil
}

func (a *AuthGRPCClient) DeleteSession(token string, logger *zap.SugaredLogger) (bool, error) {
	ctx := context.WithValue(context.Background(), MyLoggerKey, logger)
	isDeleted, err := a.grpcClient.DeleteSession(ctx, &auth.Token{
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
