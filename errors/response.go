package errors

import (
	"encoding/json"
	"log"
	"net/http"
)

const DEFAULT_ERROR_LEVEL = "Guru Meditation"

type ErrorResponse struct {
	Error      string `json:"error"`
	Details    string `json:"details,omitempty"`
	ErrorLevel string `json:"error_level,omitempty"`
}

func Respond(w http.ResponseWriter, code int, errMsg string, details string) {
	log.Printf("ERROR %d: %s - %s", code, errMsg, details)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{
		Error:      errMsg,
		Details:    details,
		ErrorLevel: DEFAULT_ERROR_LEVEL,
	})
}
