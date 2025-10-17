package handlers

import (
	bjerrors "blackjack-api/errors"
	"blackjack-api/game"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
)

var session *game.GameSession

func StartGameHandler(w http.ResponseWriter, r *http.Request) {
	showLogs(r)
	w.Header().Set("Content-Type", "application/json")
	session = game.NewGameSession()
	json.NewEncoder(w).Encode(session.GetState())
}

func HitHandler(w http.ResponseWriter, r *http.Request) {
	showLogs(r)
	if session == nil || session.GameOver {
		//http.Error(w, "No hay partida activa", http.StatusBadRequest)
		bjerrors.Respond(w, http.StatusBadRequest, bjerrors.ErrNoActiveGame.Error(), "Debes iniciar una partida antes de pedir una carta")
		return
	}
	session.Hit()
	respondWithJSON(w, session.GetState())
}

func StandHandler(w http.ResponseWriter, r *http.Request) {
	showLogs(r)
	w.Header().Set("Content-Type", "application/json")

	if session == nil || session.GameOver {
		bjerrors.Respond(w, http.StatusBadRequest, bjerrors.ErrNoActiveGame.Error(), "Debes iniciar una partida antes de plantarte")
		return
	}
	session.Stand()
	respondWithJSON(w, session.GetState())
}

func StateHandler(w http.ResponseWriter, r *http.Request) {
	showLogs(r)
	w.Header().Set("Content-Type", "application/json")
	if session == nil {
		bjerrors.Respond(w, http.StatusBadRequest, bjerrors.ErrNoActiveGame.Error(), "Debes iniciar una partida antes de consultar el estado")
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
	showLogs(r)

	session = game.NewGameSession()
	respondWithJSON(w, session.GetState())
}

func ServiceStatus(w http.ResponseWriter, r *http.Request) {
	showLogs(r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"message": "Blackjack API is running...🚀",
	})
}

func getSessionID(r *http.Request) string {
	return r.Header.Get("X-Session-ID")
}

func showLogs(r *http.Request) {
	ip, port, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		ip = r.RemoteAddr
		port = "?"
	}
	// Detect localhost (IPv4 or IPv6)
	isLocal := ip == "127.0.0.1" || ip == "::1"

	// Logs
	if isLocal {
		fmt.Printf("📥 %s %s from localhost (IP: %s, Port: %s)\n", r.Method, r.URL.Path, ip, port)
	} else {
		fmt.Printf("📥 %s %s from %s:%s\n", r.Method, r.URL.Path, ip, port)
	}
}
