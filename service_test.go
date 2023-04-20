package main

import (
	"reflect"
	"testing"
)

func TestService(t *testing.T) {
	testCases := []struct {
		name     string
		flights  []Flight
		expected *Flight
		err      error
	}{
		{
			name: "Single Flight",
			flights: []Flight{
				{"SFO", "EWR"},
			},
			expected: &Flight{"SFO", "EWR"},
			err:      nil,
		},
		{
			name: "Multiple Flights",
			flights: []Flight{
				{"ATL", "EWR"},
				{"SFO", "ATL"},
			},
			expected: &Flight{"SFO", "EWR"},
			err:      nil,
		},
		{
			name: "Multiple Flights With Hops",
			flights: []Flight{
				{"IND", "EWR"},
				{"SFO", "ATL"},
				{"GSO", "IND"},
				{"ATL", "GSO"},
			},
			expected: &Flight{"SFO", "EWR"},
			err:      nil,
		},
		{
			name: "Invalid Flights",
			flights: []Flight{
				{"SFO", "ATL"},
				{"ATL", "EWR"},
				{"EWR", "SFO"},
			},
			expected: nil,
			err:      ErrFailedToFindFlightPath,
		},
		{
			name: "Disconnected Flights",
			flights: []Flight{
				{"SFO", "ATL"},
				{"EWR", "JFK"},
			},
			expected: nil,
			err:      ErrDisconnectedFlights,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := FindFlightPath(tc.flights)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected flight path %v, got %v", tc.expected, result)
			}

			if (err != nil && tc.err == nil) || (err == nil && tc.err != nil) || (err != nil && tc.err != nil && err.Error() != tc.err.Error()) {
				t.Errorf("Expected error '%v', got '%v'", tc.err, err)
			}
		})
	}
}
