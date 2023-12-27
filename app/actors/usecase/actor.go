package actorusecase

import "kinopoisk/app/entity"

type ActorUseCase interface {
	GetActorByID(ID uint64) (*entity.Actor, error)
	GetActors() ([]*entity.Actor, error)
}
