package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// calculate is an HTTP handler function that processes incoming requests to the
// '/calculate' endpoint. It expects a JSON payload containing a list of flights
// in the request body. The function calls the FindFlightPath function to
// calculate the flight path. If successful, it returns a JSON response with
// the calculated path. In case of errors, appropriate error messages and status
// codes are returned.
func calculate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(Response{Status: "error", Message: "Method not allowed"})
		return
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding request body: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Status: "error", Message: "Invalid JSON"})
		return
	}

	flights, err := NewFlights(req)
	if err != nil {
		log.Printf("Error decoding request body: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Status: "error", Message: "Invalid request format"})
		return
	}

	flight, err := FindFlightPath(flights)
	if err != nil {
		log.Printf("Error finding flight path: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Status: "error", Message: "Failed to find flight path"})
		return
	}

	json.NewEncoder(w).Encode(Response{
		Status: "success",
		Path: []string{
			flight.Source,
			flight.Destination,
		}})
}
