package http

import "net/http"

func handleReadiness(w http.ResponseWriter, _ *http.Request) {
	type returnJson struct {
		Status string `json:"status"`
	}
	respondWithJSON(w, http.StatusOK, returnJson{Status: "ok"})
}

func handleError(w http.ResponseWriter, _ *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
