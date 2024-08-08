package http

import (
	"encoding/json"
	"net/http"
)

type errorReturn struct {
	Error string `json:"error"`
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	errorReturn := errorReturn{Error: message}
	RespondWithJSON(w, code, errorReturn)
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error marshalling JSON")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
