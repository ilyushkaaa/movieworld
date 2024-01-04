package userrepo

import (
	"database/sql"
	auth "kinopoisk/service_auth/proto"
)

type UserRepo interface {
	LoginRepo(username, password string) (*auth.User, error)
	RegisterRepo(username, password string) (*auth.User, error)
	FindUserByUsername(username string) (*auth.User, error)
}

type UserRepoMySQL struct {
	db *sql.DB
}

func NewUserRepoMySQL(db *sql.DB) *UserRepoMySQL {
	return &UserRepoMySQL{
		db: db,
	}
}

func (u *UserRepoMySQL) LoginRepo(username, password string) (*auth.User, error) {

}

func (u *UserRepoMySQL) RegisterRepo(username, password string) (*auth.User, error) {

}

func (u *UserRepoMySQL) FindUserByUsername(username string) (*auth.User, error) {

}
