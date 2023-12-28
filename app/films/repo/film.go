package filmrepo

import "kinopoisk/app/entity"

type FilmRepo interface {
	GetFilmsRepo(genre, country, producer string) ([]*entity.Film, error)
	GetFilmByIDRepo(filmID uint64) (*entity.Film, error)
	GetFilmsByActorRepo(ID uint64) ([]*entity.Film, error)
	GetFilmsByGenreRepo(genre string) ([]*entity.Film, error)
	GetFilmsByCountryRepo(country string) ([]*entity.Film, error)
	GetFilmsByProducerRepo(producer string) ([]*entity.Film, error)
	GetSoonFilmsRepo() ([]*entity.Film, error)
	GetFavouriteFilmsRepo(userID uint64) ([]*entity.Film, error)
	AddFavouriteFilmRepo(userID, filmID uint64) (bool, error)
	DeleteFavouriteFilmRepo(userID, filmID uint64) (bool, error)
}
