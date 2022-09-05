package main

import (
	"log"
	"net/http"

	"github.com/go-swagno/examples/net-http/handlers"
)

func main() {

	handler := handlers.NewHandler()

	httpMux := http.NewServeMux()

	// set mock routes
	handler.SetRoutes(httpMux)

	// set swagger routes
	handler.SetSwagger(httpMux)

	log.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", httpMux))

}
