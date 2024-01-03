package actorusecase

import (
	actorrepo "kinopoisk/app/actors/repo/mysql"
	"kinopoisk/app/entity"
	"sync"
)

type ActorUseCase interface {
	GetActorByID(ID uint64) (*entity.Actor, error)
	GetActors() ([]*entity.Actor, error)
}

type ActorUseCaseStruct struct {
	mu        *sync.RWMutex
	ActorRepo actorrepo.ActorRepo
}

func NewActorUseCaseStruct(actorRepo actorrepo.ActorRepo) *ActorUseCaseStruct {
	return &ActorUseCaseStruct{
		mu:        &sync.RWMutex{},
		ActorRepo: actorRepo,
	}
}

func (a *ActorUseCaseStruct) GetActors() ([]*entity.Actor, error) {
	actors, err := a.ActorRepo.GetActorsRepo()
	if err != nil {
		return nil, err
	}
	return actors, nil
}

func (a *ActorUseCaseStruct) GetActorByID(ID uint64) (*entity.Actor, error) {
	actor, err := a.ActorRepo.GetActorByIDRepo(ID)
	if err != nil {
		return nil, err
	}
	if actor == nil {
		return nil, nil
	}
	return actor, nil
}
