package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-swagno/examples/gorilla-mux/handlers"
	"github.com/gorilla/mux"
)

func main() {

	handler := handlers.NewHandler()

	r := mux.NewRouter()

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
