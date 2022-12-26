package routes

import (
	"log"
	"net/http"
)

// Health is the standard health endpoint.
// It returns a 200 status code if success.
func Health(w http.ResponseWriter, r *http.Request) {
	log.Println("200 OK")
	w.Header().Set("Content-type", "plain/text")
}
