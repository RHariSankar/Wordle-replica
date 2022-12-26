package main

import (
	"backend/routes/v1"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.V1Routes(router)

	log.Println("Starting server at 8000")

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatalln("There's an error with the server:", err)
	}
}
