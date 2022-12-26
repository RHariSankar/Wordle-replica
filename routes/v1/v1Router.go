package routes

import (
	"github.com/gorilla/mux"
)

// Handles all the v1 routes.
func V1Routes(router *mux.Router) {
	v1Router := router.PathPrefix("/v1").Subrouter()
	v1Router.HandleFunc("/health", Health)
	v1Router.HandleFunc("/isValid", IsValidWord).Methods("POST")
}
