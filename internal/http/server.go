package http

import (
	"log"
	"net/http"
)

// InitServer initializes the server and starts serving files on the specified port.
func InitServer(port string) {
	// This function will initialize the server
	mux := http.NewServeMux()
	fsHandler := http.FileServer(http.Dir("."))
	mux.Handle("/app/*", fsHandler)
	mux.HandleFunc("GET /v1/healthz", handleReadiness)
	mux.HandleFunc("GET /v1/err", handleError)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files on this port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
