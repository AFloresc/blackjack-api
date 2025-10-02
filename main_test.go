package main

import (
	"blackjack-api/handlers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestRouterStart(t *testing.T) {
	router := mux.NewRouter()
	api := router.PathPrefix("/blackjack/api/v1").Subrouter()
	api.HandleFunc("/start", handlers.StartGameHandler).Methods("POST")

	req := httptest.NewRequest(http.MethodPost, "/blackjack/api/v1/start", nil)
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	if res.Result().StatusCode != http.StatusOK {
		t.Errorf("esperado 200, recibido %d", res.Result().StatusCode)
	}
}
