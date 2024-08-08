package http

import (
	"net/http"
)

func HandleReadiness(w http.ResponseWriter, _ *http.Request) {
	type returnJson struct {
		Status string `json:"status"`
	}
	RespondWithJSON(w, http.StatusOK, returnJson{Status: "ok"})
}

func HandleError(w http.ResponseWriter, _ *http.Request) {
	RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
