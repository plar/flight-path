package main

import (
	"reflect"
	"testing"
)

func TestNewFlights(t *testing.T) {
	// Define test cases
	tests := []struct {
		name   string
		req    Request
		exp    []Flight
		expErr error
	}{
		{
			name: "Valid Request",
			req:  Request{{"SFO", "LAX"}, {"LAX", "JFK"}},
			exp:  []Flight{{"SFO", "LAX"}, {"LAX", "JFK"}},
		},
		{
			name:   "Empty Request",
			req:    Request{},
			expErr: ErrNoFlights,
		},
		{
			name:   "Duplicate Flight Segment",
			req:    Request{{"SFO", "LAX"}, {"LAX", "JFK"}, {"SFO", "LAX"}},
			expErr: ErrDuplicateFlightSegment,
		},
		{
			name:   "Invalid Flight Segment",
			req:    Request{{"SFO", ""}, {"LAX", "JFK"}},
			expErr: ErrWrongFlightFormat,
		},
		{
			name:   "Invalid Airport code",
			req:    Request{{"SF-", "LAX"}},
			expErr: ErrInvalidAirportCode,
		},
		{
			name:   "Invalid Airport code",
			req:    Request{{"SFO", "SFO"}},
			expErr: ErrSameSourceAndDestination,
		},
	}

	// Run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			flights, err := NewFlights(test.req)

			// Check if error matches expected error
			if err != test.expErr {
				t.Errorf("Unexpected error. Got %v, want %v", err, test.expErr)
			}

			// Check if slice of flights matches expected flights
			if !reflect.DeepEqual(flights, test.exp) {
				t.Errorf("Unexpected flights. Got %v, want %v", flights, test.exp)
			}
		})
	}
}
