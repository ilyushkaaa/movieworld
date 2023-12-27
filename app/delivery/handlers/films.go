package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"kinopoisk/app/delivery"
	filmusecase "kinopoisk/app/films/usecase"
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
	film, err := strconv.ParseUint(filmID, 10, 64)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "bad format of film id: %s"}`, err)
		delivery.WriteResponse(fh.Logger, w, []byte(errText), http.StatusBadRequest)
		return
	}
	actor, err := fh.FilmUseCases.GetFilmByID(film)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "internal server error: %s"}`, err)
		delivery.WriteResponse(fh.Logger, w, []byte(errText), http.StatusInternalServerError)
		return
	}
	filmJSON, err := json.Marshal(actor)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "error in coding film: %s"}`, err)
		delivery.WriteResponse(fh.Logger, w, []byte(errText), http.StatusInternalServerError)
		return
	}
	delivery.WriteResponse(fh.Logger, w, filmJSON, http.StatusOK)
}
