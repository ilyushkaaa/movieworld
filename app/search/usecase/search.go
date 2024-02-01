package searchusecase

import (
	"go.uber.org/zap"
	"kinopoisk/app/entity"
	searchrepo "kinopoisk/app/search/repo/mysql"
	"sync"
)

type SearchUseCase interface {
	MakeSearch(inputStr string, logger *zap.SugaredLogger) (*entity.SearchResult, error)
}

type SearchUseCaseStruct struct {
	mu         *sync.RWMutex
	searchRepo searchrepo.SearchRepo
}

func NewSearchUseCaseStruct(searchRepo searchrepo.SearchRepo) *SearchUseCaseStruct {
	return &SearchUseCaseStruct{
		mu:         &sync.RWMutex{},
		searchRepo: searchRepo,
	}
}

func (sr *SearchUseCaseStruct) MakeSearch(inputStr string, logger *zap.SugaredLogger) (*entity.SearchResult, error) {
	result, err := sr.searchRepo.MakeSearchDB(inputStr)
	if err != nil {
		logger.Errorf("error in search database method: %s", err)
		return nil, err
	}
	return result, nil
}
