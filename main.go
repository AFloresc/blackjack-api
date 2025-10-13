package main

import (
	"blackjack-api/routes"
	"log"
	"net/http"

	ghandlers "github.com/gorilla/handlers"
)

func main() {
	router := routes.InitRouter()

	cors := ghandlers.CORS(
		ghandlers.AllowedOrigins([]string{"*"}),
		ghandlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
		ghandlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	log.Println("Servidor escuchando en :8080")
	//log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))
	//log.Fatal(http.ListenAndServe(":8080", router))
	log.Fatal(http.ListenAndServe(":8080", cors(router)))
}
