package main

import (
	"backend/routes/v1"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.V1Routes(router)

	const timeout = 3
	server := &http.Server{
		Addr:              ":8000",
		ReadHeaderTimeout: time.Duration(timeout) * time.Second,
		Handler:           router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln("There's an error with the server:", err)
	}
	log.Println("Starting server at 8000")
}
