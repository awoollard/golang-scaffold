package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Main function
func main() {
	r := mux.NewRouter()

	// Endpoints and their respective handlers
	r.HandleFunc("/", GetRoot).Methods("GET")
	r.HandleFunc("/status", GetStatus).Methods("GET")

	// Start server
	log.Fatal(http.ListenAndServe(":80", r))
}
