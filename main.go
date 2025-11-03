package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Status string `json:"status"`
	Name   string `json:"name,omitempty"`
}

func main() {
	mux := http.NewServeMux()

	// Return JSON response with status
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{
			Status: "success",
		})
	})

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{
			Status: "success",
			Name:   name,
		})
	})


	
	// I create a server, no framework, just pure Go
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// I start the server, if fails, I log the error
	log.Fatal(srv.ListenAndServe())
}
