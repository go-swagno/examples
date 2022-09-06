package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-swagno/examples/chi/handlers"
)

func main() {

	handler := handlers.NewHandler()

	r := chi.NewRouter()

	// set mock routes
	handler.SetRoutes(r)

	// set swagger routes
	handler.SetSwagger(r)

	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Listening on localhost:8080")
	log.Fatal(srv.ListenAndServe())

}
