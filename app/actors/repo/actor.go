package actorrepo

import "kinopoisk/app/entity"

type ActorRepo interface {
	GetActorByIDRepo(ID uint64) (*entity.Actor, error)
	GetActorsRepo() ([]*entity.Actor, error)
}
