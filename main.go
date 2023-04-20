package main

import (
	"log"
	"net/http"
)

// The main function initializes and starts the HTTP server.
func main() {
	http.HandleFunc("/calculate", calculate)

	log.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
