package filmusecase

import "kinopoisk/app/entity"

type FilmUseCase interface {
	GetFilms(genre, country, producer string) ([]*entity.Film, error)
	GetFilmByID(filmID uint64) (*entity.Film, error)
	GetFilmsByActor(ID uint64) ([]*entity.Film, error)
	GetFilmsByGenre(genre string) ([]*entity.Film, error)
	GetFilmsByCountry(country string) ([]*entity.Film, error)
	GetFilmsByProducer(producer string) ([]*entity.Film, error)
	GetSoonFilms() ([]*entity.Film, error)
	GetFavouriteFilms(userID uint64) ([]*entity.Film, error)
	AddFavouriteFilm(userID, filmID uint64) (bool, error)
	DeleteFavouriteFilm(userID, filmID uint64) (bool, error)
}
