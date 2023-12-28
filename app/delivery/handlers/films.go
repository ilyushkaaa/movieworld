package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"kinopoisk/app/delivery"
	"kinopoisk/app/entity"
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
	for key, _ := range query {
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
		errText := fmt.Sprintf(`{"message": "bad params in query"}`)
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
	filmJSON, err := json.Marshal(film)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "error in coding film: %s"}`, err)
		delivery.WriteResponse(fh.Logger, w, []byte(errText), http.StatusInternalServerError)
		return
	}
	delivery.WriteResponse(fh.Logger, w, filmJSON, http.StatusOK)
}

func (fh *FilmHandler) GetFilmByActor(w http.ResponseWriter, r *http.Request) {
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
	if err != nil {
		errText := fmt.Sprintf(`{"message": "internal server error: %s"}`, err)
		delivery.WriteResponse(fh.Logger, w, []byte(errText), http.StatusInternalServerError)
		return
	}
	if !wasAdded {
		result := fmt.Sprintf(`{"result": "was not added"}`)
		delivery.WriteResponse(fh.Logger, w, []byte(result), http.StatusOK)
		return
	}
	result := fmt.Sprintf(`{"result": "was added"}`)
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
		result := fmt.Sprintf(`{"meassage": "not found"}`)
		delivery.WriteResponse(fh.Logger, w, []byte(result), http.StatusNotFound)
		return
	}
	result := fmt.Sprintf(`{"result": "success"}`)
	delivery.WriteResponse(fh.Logger, w, []byte(result), http.StatusOK)
}
