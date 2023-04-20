package main

// FindFlightPath takes a slice of unordered flights and returns a flight.
// If the function fails to find a valid flight path, it returns an error.
func FindFlightPath(flights []Flight) (flight *Flight, err error) {
	if flight = findFlightPathCandidate(flights); flight == nil {
		return nil, ErrFailedToFindFlightPath
	}

	if hasDisconnectedFlights(flight, flights) {
		return nil, ErrDisconnectedFlights
	}

	return flight, nil
}

// findFlightPathCandidate takes a slice of unordered flights and returns
// the candidate flight path.
// If the function fails to find a candidate flight path, it returns nil.
// The returned flight path candidate may not be a valid flight path
// if there are disconnected flights. A hasDisconnectedFlights function should be
// used to validate the connectivity of the flight path.
//
// This algorithm is based on the observation that a valid flight path must start at an airport
// that has more departing flights than arriving flights (or no arriving flights),
// and end at an airport that has more arriving flights than departing flights (or no departing flights).
// If there are multiple airports that meet these criteria, the algorithm simply chooses the first one it encounters.
//
// This function has a time complexity of O(n), where n is the number of flights in the input slice.
// The space complexity of the function is O(m), where m is the number of distinct airports.
func findFlightPathCandidate(flights []Flight) *Flight {
	airportCounts := make(map[string]int)
	for _, flight := range flights {
		airportCounts[flight.Source]++
		airportCounts[flight.Destination]--
	}

	var f Flight // no source, no dest
	for airport, count := range airportCounts {
		if f.Source == "" && count > 0 {
			f.Source = airport
		} else if f.Destination == "" && count < 0 {
			f.Destination = airport
		}

		if f.Source != "" && f.Destination != "" {
			break
		}
	}

	if f.Source == "" || f.Destination == "" {
		return nil
	}

	return &f
}

// hasDisconnectedFlights checks if there are any disconnected flights in the given
// flight path and flights list. It takes a flightPath slice containing the starting
// and ending airports and a flights slice containing all the flights. The function
// returns true if there are disconnected flights or the flight path is incomplete,
// otherwise, it returns false. A disconnected flight is one where there is no
// continuous path between the source and destination airports considering all flights.
func hasDisconnectedFlights(flight *Flight, flights []Flight) bool {
	src, dst := flight.Source, flight.Destination
	if len(flights) == 1 {
		osrc, odst := flights[0].Source, flights[0].Destination
		return !((osrc == src && odst == dst) || (osrc == dst && odst == src))
	}

	// a breadth-first search (BFS) algorithm to traverse the flights from the starting airport.
	// If the end airport is not visited during the BFS traversal or srd and dst airports belong
	// to the same flight it means there are disconnected flights.
	singleFlights := make(map[string]bool)
	visited := make(map[string]bool)
	queue := []string{src}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current] {
			continue
		}

		lazySFInit := len(singleFlights) == 0
		for _, flight := range flights {
			if lazySFInit {
				singleFlights[flight.ID()] = true
			}

			if flight.Source == current && !visited[flight.Destination] {
				queue = append(queue, flight.Destination)
			}
		}

		visited[current] = true
	}

	return dst == "" || !visited[dst] || singleFlights[NewFlightID(src, dst)]
}
