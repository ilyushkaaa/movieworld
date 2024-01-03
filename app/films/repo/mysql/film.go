package filmrepo

import (
	"database/sql"
	"errors"
	"go.uber.org/zap"
	"kinopoisk/app/entity"
	"time"
)

type FilmRepo interface {
	GetFilmsRepo(genre, country, producer string) ([]*entity.Film, error)
	GetFilmByIDRepo(filmID uint64) (*entity.Film, error)
	GetFilmsByActorRepo(ID uint64) ([]*entity.Film, error)
	GetSoonFilmsRepo(date time.Time) ([]*entity.Film, error)
	GetFavouriteFilmsRepo(userID uint64) ([]*entity.Film, error)
	AddFavouriteFilmRepo(userID, filmID uint64) (bool, error)
	DeleteFavouriteFilmRepo(userID, filmID uint64) (bool, error)
	GetFilmActorsRepo(filmID uint64) ([]*entity.Actor, error)
	GetFilmGenresRepo(filmID uint64) ([]*entity.Genre, error)
}

type FilmRepoMySQL struct {
	db     *sql.DB
	logger *zap.SugaredLogger
}

func NewActorRepoMySQL(db *sql.DB, logger *zap.SugaredLogger) *FilmRepoMySQL {
	return &FilmRepoMySQL{
		db:     db,
		logger: logger,
	}
}

func (r *FilmRepoMySQL) GetFilmsRepo(genre, country, producer string) ([]*entity.Film, error) {
	var args []interface{}
	query := "SELECT f.id, f.name, f.description, f.duration, f.min_age, f.country, f.producer_name, f.date_of_release from films f"
	if genre != "" {
		query += " INNER JOIN film_genres fg ON f.id = fg.film_id INNER JOIN genres g ON g.id = fg.genre_id WHERE g.name = ?"
		args = append(args, genre)
	} else {
		query += " WHERE 1 = 1"
	}
	if country != "" {
		args = append(args, country)
		query += " AND f.country = ?"
	}
	if producer != "" {
		args = append(args, producer)
		query += " AND f.producer_name = ?"
	}
	rows, err := r.db.Query(query, args...)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			r.logger.Errorf("error in closing db rows in mysql")
		}
	}(rows)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	films := make([]*entity.Film, 0)
	for rows.Next() {
		film := &entity.Film{}
		err = rows.Scan(&film.ID, &film.Name, &film.Description, &film.Duration, &film.MinAge, &film.Country,
			&film.ProducerName, film.DateOfRelease)
		if err != nil {
			return nil, err
		}
		films = append(films, film)
	}
	return films, nil
}

func (r *FilmRepoMySQL) GetFilmByIDRepo(filmID uint64) (*entity.Film, error) {
	return nil, nil
}

func (r *FilmRepoMySQL) GetFilmsByActorRepo(ID uint64) ([]*entity.Film, error) {
	return nil, nil
}

func (r *FilmRepoMySQL) GetSoonFilmsRepo(date time.Time) ([]*entity.Film, error) {
	return nil, nil
}

func (r *FilmRepoMySQL) GetFavouriteFilmsRepo(userID uint64) ([]*entity.Film, error) {
	return nil, nil
}

func (r *FilmRepoMySQL) AddFavouriteFilmRepo(userID, filmID uint64) (bool, error) {
	return false, nil
}

func (r *FilmRepoMySQL) DeleteFavouriteFilmRepo(userID, filmID uint64) (bool, error) {
	return false, nil
}

func (r *FilmRepoMySQL) GetFilmActorsRepo(filmID uint64) ([]*entity.Actor, error) {
	return nil, nil
}

func (r *FilmRepoMySQL) GetFilmGenresRepo(filmID uint64) ([]*entity.Genre, error) {
	return nil, nil
}
