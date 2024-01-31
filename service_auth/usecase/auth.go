package authserviceusecase

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
	auth "kinopoisk/service_auth/proto"
	userrepo "kinopoisk/service_auth/repo/mysql"
	sessionrepo "kinopoisk/service_auth/repo/redis"
	"os"
	"sync"
	"time"
)

type AuthGRPCServer struct {
	auth.UnimplementedAuthMakerServer

	mu          *sync.RWMutex
	UserRepo    userrepo.UserRepo
	SessionRepo sessionrepo.SessionRepo
	secret      []byte
	logger      *zap.SugaredLogger
}

func NewAuthGRPCServer(userRepo userrepo.UserRepo, sessionRepo sessionrepo.SessionRepo, logger *zap.SugaredLogger) *AuthGRPCServer {
	return &AuthGRPCServer{
		UnimplementedAuthMakerServer: auth.UnimplementedAuthMakerServer{},
		UserRepo:                     userRepo,
		SessionRepo:                  sessionRepo,
		secret:                       []byte(os.Getenv("SECRET")),
		mu:                           &sync.RWMutex{},
		logger:                       logger,
	}
}

func (a *AuthGRPCServer) Login(_ context.Context, in *auth.AuthData) (*auth.User, error) {
	hashPassword, err := getHashPassword(in.Password)
	if err != nil {
		a.logger.Errorf("error in getting hash password: %s", err)
		return &auth.User{}, err
	}
	a.mu.RLock()
	loggedInUser, err := a.UserRepo.LoginRepo(in.Username, hashPassword)
	a.mu.RUnlock()
	if err != nil {
		a.logger.Errorf("error in login user in db: %s", err)
		return &auth.User{}, err
	}
	if loggedInUser == nil {
		loggedInUser = &auth.User{}
	}
	return loggedInUser, nil
}

func (a *AuthGRPCServer) Register(_ context.Context, in *auth.AuthData) (*auth.User, error) {
	a.mu.RLock()
	loggedInUser, err := a.UserRepo.FindUserByUsername(in.Username)
	a.mu.RUnlock()
	if err != nil {
		a.logger.Errorf("error in getting user by username: %s", err)
		return &auth.User{}, err
	}
	if loggedInUser != nil {
		a.logger.Errorf("user with login %s already exists", in.Username)
		return &auth.User{}, nil
	}
	hashPassword, err := getHashPassword(in.Password)
	if err != nil {
		a.logger.Errorf("error in getting hash password: %s", err)
		return &auth.User{}, err
	}
	a.mu.Lock()
	newUser, err := a.UserRepo.RegisterRepo(in.Username, hashPassword)
	a.mu.Unlock()
	if err != nil {
		a.logger.Errorf("error in register user in db: %s", err)
		return &auth.User{}, err
	}
	if newUser == nil {
		newUser = &auth.User{}
	}
	return newUser, nil
}

func (a *AuthGRPCServer) CreateSession(_ context.Context, in *auth.User) (*auth.Token, error) {
	token, err := a.newToken(in)
	if err != nil {
		a.logger.Errorf("error in getting session token: %s", err)
		return &auth.Token{}, err
	}
	newSession := &auth.Session{
		ID:   token,
		User: in,
	}
	a.mu.Lock()
	err = a.SessionRepo.CreateSessionRepo(newSession)
	a.mu.Unlock()
	if err != nil {
		a.logger.Errorf("error in creating session: %s", err)
		return &auth.Token{}, err
	}
	return &auth.Token{Token: token}, nil

}

func (a *AuthGRPCServer) GetSession(_ context.Context, in *auth.Token) (*auth.Session, error) {
	hashSecretGetter := func(token *jwt.Token) (interface{}, error) {
		method, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok || method.Alg() != "HS256" {
			a.logger.Errorf("bad sign in")
			return nil, fmt.Errorf("bad sign method")
		}
		return a.secret, nil
	}
	token, err := jwt.Parse(in.Token, hashSecretGetter)
	if err != nil || !token.Valid {
		a.logger.Errorf("bad secret")
		return &auth.Session{}, nil
	}
	a.mu.RLock()
	sess, err := a.SessionRepo.GetSessionRepo(in.Token)
	a.mu.RUnlock()
	if err != nil {
		a.logger.Errorf("error in getting session from db: %s", err)
		return &auth.Session{}, err
	}
	return sess, nil

}

func (a *AuthGRPCServer) DeleteSession(_ context.Context, in *auth.Token) (*auth.IsDeleted, error) {
	a.mu.Lock()
	idDeleted, err := a.SessionRepo.DeleteSessionRepo(in.Token)
	a.mu.Unlock()
	if err != nil {
		a.logger.Errorf("error in deleting session in db: %s", err)
		return &auth.IsDeleted{IsDeleted: false}, err
	}
	return &auth.IsDeleted{IsDeleted: idDeleted}, nil
}

func getHashPassword(password string) (string, error) {
	hash := sha256.New()
	_, err := hash.Write([]byte(password))
	if err != nil {
		return "", err
	}
	hashBytes := hash.Sum(nil)
	hashPass := hex.EncodeToString(hashBytes)
	return hashPass, nil
}

func (a *AuthGRPCServer) newToken(user *auth.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour * 24 * 7).Unix(),
	})
	tokenString, err := token.SignedString(a.secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
