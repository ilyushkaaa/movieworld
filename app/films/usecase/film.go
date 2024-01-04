package filmusecase

import (
	"errors"
	"kinopoisk/app/entity"
	errorapp "kinopoisk/app/errors"
	filmrepo "kinopoisk/app/films/repo/mysql"
	"sync"
	"time"
)

type FilmUseCase interface {
	GetFilms(genre, country, producer string) ([]*entity.Film, error)
	GetFilmByID(filmID uint64) (*entity.Film, error)
	GetFilmsByActor(ID uint64) ([]*entity.Film, error)
	GetSoonFilms() ([]*entity.Film, error)
	GetFavouriteFilms(userID uint64) ([]*entity.Film, error)
	AddFavouriteFilm(userID, filmID uint64) (bool, error)
	DeleteFavouriteFilm(userID, filmID uint64) (bool, error)
	GetFilmActors(filmID uint64) ([]*entity.Actor, error)
	GetFilmGenres(filmID uint64) ([]*entity.Genre, error)
}

type FilmUseCaseStruct struct {
	mu       *sync.RWMutex
	FilmRepo filmrepo.FilmRepo
}

func NewFilmUseCaseStruct(filmRepo filmrepo.FilmRepo) *FilmUseCaseStruct {
	return &FilmUseCaseStruct{
		mu:       &sync.RWMutex{},
		FilmRepo: filmRepo,
	}
}

func (f *FilmUseCaseStruct) GetFilms(genre, country, producer string) ([]*entity.Film, error) {
	films, err := f.FilmRepo.GetFilmsRepo(genre, country, producer)
	if err != nil {
		return nil, err
	}
	return films, nil
}

func (f *FilmUseCaseStruct) GetFilmByID(filmID uint64) (*entity.Film, error) {
	film, err := f.FilmRepo.GetFilmByIDRepo(filmID)
	if err != nil {
		return nil, err
	}
	if film == nil {
		return nil, nil
	}
	return film, nil
}

func (f *FilmUseCaseStruct) GetFilmsByActor(ID uint64) ([]*entity.Film, error) {
	films, err := f.FilmRepo.GetFilmsByActorRepo(ID)
	if err != nil {
		return nil, err
	}
	return films, nil
}

func (f *FilmUseCaseStruct) GetSoonFilms() ([]*entity.Film, error) {
	currentDate := time.Now().Format("2006-01-02")
	films, err := f.FilmRepo.GetSoonFilmsRepo(currentDate)
	if err != nil {
		return nil, err
	}
	return films, nil
}

func (f *FilmUseCaseStruct) GetFavouriteFilms(userID uint64) ([]*entity.Film, error) {
	films, err := f.FilmRepo.GetFavouriteFilmsRepo(userID)
	if err != nil {
		return nil, err
	}
	return films, nil
}

func (f *FilmUseCaseStruct) AddFavouriteFilm(userID, filmID uint64) (bool, error) {
	film, err := f.FilmRepo.GetFilmByIDRepo(filmID)
	if err != nil {
		return false, err
	}
	if film == nil {
		return false, errorapp.ErrorNoFilm
	}
	_, err = f.FilmRepo.GetFilmInFavourites(filmID, userID)
	if errors.Is(err, errorapp.ErrorNoFilm) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	wasAdded, err := f.FilmRepo.AddFavouriteFilmRepo(userID, filmID)
	if err != nil {
		return false, err
	}
	return wasAdded, nil
}

func (f *FilmUseCaseStruct) DeleteFavouriteFilm(userID, filmID uint64) (bool, error) {
	ID, err := f.FilmRepo.GetFilmInFavourites(filmID, userID)
	if errors.Is(err, errorapp.ErrorNoFilm) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	wasDeleted, err := f.FilmRepo.DeleteFavouriteFilmRepo(ID)
	if err != nil {
		return false, err
	}
	return wasDeleted, nil
}

func (f *FilmUseCaseStruct) GetFilmActors(filmID uint64) ([]*entity.Actor, error) {
	film, err := f.GetFilmByID(filmID)
	if err != nil {
		return nil, err
	}
	if film == nil {
		return nil, errorapp.ErrorNoFilm
	}
	actors, err := f.FilmRepo.GetFilmActorsRepo(filmID)
	if err != nil {
		return nil, err
	}
	return actors, nil
}

func (f *FilmUseCaseStruct) GetFilmGenres(filmID uint64) ([]*entity.Genre, error) {
	film, err := f.GetFilmByID(filmID)
	if err != nil {
		return nil, err
	}
	if film == nil {
		return nil, errorapp.ErrorNoFilm
	}
	genres, err := f.FilmRepo.GetFilmGenresRepo(filmID)
	if err != nil {
		return nil, err
	}
	return genres, nil
}
