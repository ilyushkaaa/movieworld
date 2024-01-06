package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"kinopoisk/app/delivery"
	"kinopoisk/app/entity"
	errorapp "kinopoisk/app/errors"
	filmusecase "kinopoisk/app/films/usecase"
	"kinopoisk/app/middleware"
	"net/http"
	"net/url"
	"strconv"
)

type FilmHandler struct {
	FilmUseCases filmusecase.FilmUseCase
	Logger       *zap.SugaredLogger
}

func NewFilmHandler(filmUseCases filmusecase.FilmUseCase, logger *zap.SugaredLogger) *FilmHandler {
	return &FilmHandler{
		FilmUseCases: filmUseCases,
		Logger:       logger,
	}
}

func checkUnknownParams(query url.Values) error {
	for key := range query {
		if key != "genre" && key != "country" && key != "director" {
			return fmt.Errorf("unknown param")
		}
	}
	return nil
}

func (fh *FilmHandler) GetFilms(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	err := checkUnknownParams(query)
	if err != nil {
		errText := `{"message": "bad params in query"}`
		delivery.WriteResponse(fh.Logger, w, []byte(errText), http.StatusBadRequest)
		return
	}
	genre := query.Get("genre")
	country := query.Get("country")
	director := query.Get("producer")
	films, err := fh.FilmUseCases.GetFilms(genre, country, director)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "internal server error: %s"}`, err)
		delivery.WriteResponse(fh.Logger, w, []byte(errText), http.StatusInternalServerError)
		return
	}
	filmsJSON, err := json.Marshal(films)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "error in coding films: %s"}`, err)
		delivery.WriteResponse(fh.Logger, w, []byte(errText), http.StatusInternalServerError)
		return
	}
	delivery.WriteResponse(fh.Logger, w, filmsJSON, http.StatusOK)
}

func (fh *FilmHandler) GetFilmByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filmID := vars["FILM_ID"]
	filmIDInt, err := strconv.ParseUint(filmID, 10, 64)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "bad format of film id: %s"}`, err)
		delivery.WriteResponse(fh.Logger, w, []byte(errText), http.StatusBadRequest)
		return
	}
	film, err := fh.FilmUseCases.GetFilmByID(filmIDInt)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "internal server error: %s"}`, err)
		delivery.WriteResponse(fh.Logger, w, []byte(errText), http.StatusInternalServerError)
		return
	}
	if film == nil {
		errText := fmt.Sprintf(`{"message": "film with ID %d is not found"}`, filmIDInt)
		delivery.WriteResponse(fh.Logger, w, []byte(errText), http.StatusNotFound)
		return
	}
	filmJSON, err := json.Marshal(film)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "error in coding film: %s"}`, err)
		delivery.WriteResponse(fh.Logger, w, []byte(errText), http.StatusInternalServerError)
		return
	}
	delivery.WriteResponse(fh.Logger, w, filmJSON, http.StatusOK)
}

func (fh *FilmHandler) GetFilmsByActor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	actorID := vars["ACTOR_ID"]
	actorIDInt, err := strconv.ParseUint(actorID, 10, 64)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "bad format of actor id: %s"}`, err)
		delivery.WriteResponse(fh.Logger, w, []byte(errText), http.StatusBadRequest)
		return
	}
	films, err := fh.FilmUseCases.GetFilmsByActor(actorIDInt)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "internal server error: %s"}`, err)
		delivery.WriteResponse(fh.Logger, w, []byte(errText), http.StatusInternalServerError)
		return
	}
	filmsJSON, err := json.Marshal(films)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "error in coding films: %s"}`, err)
		delivery.WriteResponse(fh.Logger, w, []byte(errText), http.StatusInternalServerError)
		return
	}
	delivery.WriteResponse(fh.Logger, w, filmsJSON, http.StatusOK)
}

func (fh *FilmHandler) GetFilmsSoon(w http.ResponseWriter, r *http.Request) {
	films, err := fh.FilmUseCases.GetSoonFilms()
	if err != nil {
		errText := fmt.Sprintf(`{"message": "internal server error: %s"}`, err)
		delivery.WriteResponse(fh.Logger, w, []byte(errText), http.StatusInternalServerError)
		return
	}
	filmsJSON, err := json.Marshal(films)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "error in coding films: %s"}`, err)
		delivery.WriteResponse(fh.Logger, w, []byte(errText), http.StatusInternalServerError)
		return
	}
	delivery.WriteResponse(fh.Logger, w, filmsJSON, http.StatusOK)
}

func (fh *FilmHandler) GetFavouriteFilms(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, ok := ctx.Value(middleware.MyUserKey).(*entity.User)
	if !ok {
		delivery.WriteResponse(fh.Logger, w, []byte(`{"message": "can not cast context value to user"}`), http.StatusInternalServerError)
		return
	}
	films, err := fh.FilmUseCases.GetFavouriteFilms(user.ID)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "internal server error: %s"}`, err)
		delivery.WriteResponse(fh.Logger, w, []byte(errText), http.StatusInternalServerError)
		return
	}
	filmsJSON, err := json.Marshal(films)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "error in coding films: %s"}`, err)
		delivery.WriteResponse(fh.Logger, w, []byte(errText), http.StatusInternalServerError)
		return
	}
	delivery.WriteResponse(fh.Logger, w, filmsJSON, http.StatusOK)
}

func (fh *FilmHandler) AddFavouriteFilm(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, ok := ctx.Value(middleware.MyUserKey).(*entity.User)
	if !ok {
		delivery.WriteResponse(fh.Logger, w, []byte(`{"message": "can not cast context value to user"}`), http.StatusInternalServerError)
		return
	}
	vars := mux.Vars(r)
	filmID := vars["FILM_ID"]
	filmIDInt, err := strconv.ParseUint(filmID, 10, 64)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "bad format of actor id: %s"}`, err)
		delivery.WriteResponse(fh.Logger, w, []byte(errText), http.StatusBadRequest)
		return
	}
	wasAdded, err := fh.FilmUseCases.AddFavouriteFilm(user.ID, filmIDInt)
	if errors.Is(err, errorapp.ErrorNoFilm) {
		errText := fmt.Sprintf(`{"message": "%s"}`, err)
		delivery.WriteResponse(fh.Logger, w, []byte(errText), http.StatusNotFound)
		return
	}
	if err != nil {
		errText := fmt.Sprintf(`{"message": "internal server error: %s"}`, err)
		delivery.WriteResponse(fh.Logger, w, []byte(errText), http.StatusInternalServerError)
		return
	}
	if !wasAdded {
		result := `{"result": "was not added"}`
		delivery.WriteResponse(fh.Logger, w, []byte(result), http.StatusOK)
		return
	}
	result := `{"result": "was added"}`
	delivery.WriteResponse(fh.Logger, w, []byte(result), http.StatusOK)
}

func (fh *FilmHandler) DeleteFavouriteFilm(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, ok := ctx.Value(middleware.MyUserKey).(*entity.User)
	if !ok {
		delivery.WriteResponse(fh.Logger, w, []byte(`{"message": "can not cast context value to user"}`), http.StatusInternalServerError)
		return
	}
	vars := mux.Vars(r)
	filmID := vars["FILM_ID"]
	filmIDInt, err := strconv.ParseUint(filmID, 10, 64)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "bad format of actor id: %s"}`, err)
		delivery.WriteResponse(fh.Logger, w, []byte(errText), http.StatusBadRequest)
		return
	}
	wasDeleted, err := fh.FilmUseCases.DeleteFavouriteFilm(user.ID, filmIDInt)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "internal server error: %s"}`, err)
		delivery.WriteResponse(fh.Logger, w, []byte(errText), http.StatusInternalServerError)
		return
	}
	if !wasDeleted {
		result := fmt.Sprintf(`{"meassage": "film with ID %d is not found in favourites"}`, filmIDInt)
		delivery.WriteResponse(fh.Logger, w, []byte(result), http.StatusNotFound)
		return
	}
	result := `{"result": "success"}`
	delivery.WriteResponse(fh.Logger, w, []byte(result), http.StatusOK)
}

func (fh *FilmHandler) GetFilmActors(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filmID := vars["FILM_ID"]
	filmIDint, err := strconv.ParseUint(filmID, 10, 64)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "bad format of actor id: %s"}`, err)
		delivery.WriteResponse(fh.Logger, w, []byte(errText), http.StatusBadRequest)
		return
	}
	actors, err := fh.FilmUseCases.GetFilmActors(filmIDint)
	if errors.Is(err, errorapp.ErrorNoFilm) {
		errText := fmt.Sprintf(`{"message": "no film with id: %d"}`, filmIDint)
		delivery.WriteResponse(fh.Logger, w, []byte(errText), http.StatusNotFound)
		return
	}
	if err != nil {
		errText := fmt.Sprintf(`{"message": "internal server error: %s"}`, err)
		delivery.WriteResponse(fh.Logger, w, []byte(errText), http.StatusInternalServerError)
		return
	}
	actorsJSON, err := json.Marshal(actors)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "error in coding actors: %s"}`, err)
		delivery.WriteResponse(fh.Logger, w, []byte(errText), http.StatusInternalServerError)
		return
	}
	delivery.WriteResponse(fh.Logger, w, actorsJSON, http.StatusOK)
}

func (fh *FilmHandler) GetFilmGenres(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filmID := vars["FILM_ID"]
	filmIDint, err := strconv.ParseUint(filmID, 10, 64)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "bad format of actor id: %s"}`, err)
		delivery.WriteResponse(fh.Logger, w, []byte(errText), http.StatusBadRequest)
		return
	}
	genres, err := fh.FilmUseCases.GetFilmGenres(filmIDint)
	if errors.Is(err, errorapp.ErrorNoFilm) {
		errText := fmt.Sprintf(`{"message": "no film with id: %d"}`, filmIDint)
		delivery.WriteResponse(fh.Logger, w, []byte(errText), http.StatusNotFound)
		return
	}
	if err != nil {
		errText := fmt.Sprintf(`{"message": "internal server error: %s"}`, err)
		delivery.WriteResponse(fh.Logger, w, []byte(errText), http.StatusInternalServerError)
		return
	}
	actorsJSON, err := json.Marshal(genres)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "error in coding actors: %s"}`, err)
		delivery.WriteResponse(fh.Logger, w, []byte(errText), http.StatusInternalServerError)
		return
	}
	delivery.WriteResponse(fh.Logger, w, actorsJSON, http.StatusOK)
}
