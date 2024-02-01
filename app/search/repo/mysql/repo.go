package searchrepo

import (
	"database/sql"
	"go.uber.org/zap"
	"kinopoisk/app/entity"
)

type SearchRepo interface {
	MakeSearchDB(inputStr string) (*entity.SearchResult, error)
}

type SearchRepoMySQL struct {
	db     *sql.DB
	logger *zap.SugaredLogger
}

func NewSearchRepoMySQL(db *sql.DB, logger *zap.SugaredLogger) *SearchRepoMySQL {
	return &SearchRepoMySQL{
		db:     db,
		logger: logger,
	}
}

func (sr *SearchRepoMySQL) MakeSearchDB(inputStr string) (*entity.SearchResult, error) {
	return &entity.SearchResult{}, nil
}
