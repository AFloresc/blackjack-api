package handlers

import (
	"blackjack-api/game"
	"encoding/json"
	"net/http"
)

var session *game.GameSession

func StartGameHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	session = game.NewGameSession()
	json.NewEncoder(w).Encode(session.GetState())
}

func HitHandler(w http.ResponseWriter, r *http.Request) {
	if session == nil || session.GameOver {
		http.Error(w, "No hay partida activa", http.StatusBadRequest)
		return
	}
	session.Hit()
	respondWithJSON(w, session.GetState())
}

func StandHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if session == nil || session.GameOver {
		http.Error(w, "No hay partida activa", http.StatusBadRequest)
		return
	}
	session.Stand()
	respondWithJSON(w, session.GetState())
}

func StateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if session == nil {
		http.Error(w, "No hay partida activa", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(session.GetState())
}

func respondWithJSON(w http.ResponseWriter, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, "Error al serializar JSON", http.StatusInternalServerError)
	}
}

func RestartGameHandler(w http.ResponseWriter, r *http.Request) {
	session = game.NewGameSession()
	respondWithJSON(w, session.GetState())
}

func getSessionID(r *http.Request) string {
	return r.Header.Get("X-Session-ID")
}
