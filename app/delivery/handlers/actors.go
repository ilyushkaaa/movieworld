package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	actorusecase "kinopoisk/app/actors/usecase"
	"kinopoisk/app/delivery"
	"net/http"
	"strconv"
)

type ActorHandler struct {
	ActorUseCases actorusecase.ActorUseCase
	Logger        *zap.SugaredLogger
}

func NewActorHandler(actorUseCases actorusecase.ActorUseCase, logger *zap.SugaredLogger) *ActorHandler {
	return &ActorHandler{
		ActorUseCases: actorUseCases,
		Logger:        logger,
	}
}

func (ah *ActorHandler) GetActors(w http.ResponseWriter, _ *http.Request) {
	actors, err := ah.ActorUseCases.GetActors()
	if err != nil {
		errText := fmt.Sprintf(`{"message": "internal server error: %s"}`, err)
		delivery.WriteResponse(ah.Logger, w, []byte(errText), http.StatusInternalServerError)
		return
	}
	actorsJSON, err := json.Marshal(actors)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "error in coding actors: %s"}`, err)
		delivery.WriteResponse(ah.Logger, w, []byte(errText), http.StatusInternalServerError)
		return
	}
	delivery.WriteResponse(ah.Logger, w, actorsJSON, http.StatusOK)
}

func (ah *ActorHandler) GetActorByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	actorID := vars["ACTOR_ID"]
	actorIDInt, err := strconv.ParseUint(actorID, 10, 64)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "bad format of actor id: %s"}`, err)
		delivery.WriteResponse(ah.Logger, w, []byte(errText), http.StatusBadRequest)
		return
	}
	actor, err := ah.ActorUseCases.GetActorByID(actorIDInt)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "internal server error: %s"}`, err)
		delivery.WriteResponse(ah.Logger, w, []byte(errText), http.StatusInternalServerError)
		return
	}
	actorJSON, err := json.Marshal(actor)
	if err != nil {
		errText := fmt.Sprintf(`{"message": "error in coding actor: %s"}`, err)
		delivery.WriteResponse(ah.Logger, w, []byte(errText), http.StatusInternalServerError)
		return
	}
	delivery.WriteResponse(ah.Logger, w, actorJSON, http.StatusOK)
}
