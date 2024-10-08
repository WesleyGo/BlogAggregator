package main

import (
	"log"
	"net/http"

	httpHelper "github.com/WESLEYGO/BlogAggregator/internal/http"
)

// InitServer initializes the server and starts serving files on the specified port.
func InitServer(port string, apiCfg *apiConfig) {
	// This function will initialize the server
	mux := http.NewServeMux()
	fsHandler := http.FileServer(http.Dir("."))
	mux.Handle("/app/*", fsHandler)
	mux.HandleFunc("GET /v1/healthz", httpHelper.HandleReadiness)
	mux.HandleFunc("GET /v1/err", httpHelper.HandleError)
	mux.HandleFunc("POST /v1/users", apiCfg.handlerUsersCreate)
	mux.HandleFunc("GET /v1/users", apiCfg.handlerUsersGet)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files on this port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
