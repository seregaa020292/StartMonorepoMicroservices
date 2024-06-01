package main

import (
	"log"
	"net"
	"net/http"

	"seregaa020292/StartMonorepoMicroservices/fines/monitoring/internal/hi"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Write(hi.Message())
	})
	mux.HandleFunc("GET /health/{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	mux.HandleFunc("GET /ping/{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "pong"}`))
	})

	log.Println("Starting server monitoring")
	http.ListenAndServe(net.JoinHostPort("", "8080"), mux)
}
