package routes

import (
	"blackjack-api/handlers"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()

	api := router.PathPrefix("/blackjack/api/v1").Subrouter()
	api.HandleFunc("/", handlers.ServiceStatus).Methods("GET")
	api.HandleFunc("/start", handlers.StartGameHandler).Methods("POST")
	api.HandleFunc("/hit", handlers.HitHandler).Methods("POST")
	api.HandleFunc("/stand", handlers.StandHandler).Methods("POST")
	api.HandleFunc("/state", handlers.StateHandler).Methods("GET")
	api.HandleFunc("/restart", handlers.RestartGameHandler).Methods("POST")

	return router
}
