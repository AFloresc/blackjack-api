package main

import (
	"blackjack-api/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	api := router.PathPrefix("/blackjack/api/v1").Subrouter()
	api.HandleFunc("/", handlers.ServiceStatus).Methods("GET")
	api.HandleFunc("/start", handlers.StartGameHandler).Methods("POST")
	api.HandleFunc("/hit", handlers.HitHandler).Methods("POST")
	api.HandleFunc("/stand", handlers.StandHandler).Methods("POST")
	api.HandleFunc("/state", handlers.StateHandler).Methods("GET")
	api.HandleFunc("/restart", handlers.RestartGameHandler).Methods("POST")

	log.Println("Servidor escuchando en :8080")
	//log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))
	log.Fatal(http.ListenAndServe(":8080", router))
}
