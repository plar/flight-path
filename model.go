package main

import (
	"errors"
	"regexp"
	"strings"
)

// Service errors
var (
	ErrWrongFlightFormat        = errors.New("flight should have source and destination")
	ErrNoFlights                = errors.New("at least one flight segment is required")
	ErrInvalidAirportCode       = errors.New("invalid airport code")
	ErrDuplicateFlightSegment   = errors.New("duplicate flight segment")
	ErrSameSourceAndDestination = errors.New("source and destination airports must be different")
	ErrFailedToFindFlightPath   = errors.New("failed to find flight path")
	ErrDisconnectedFlights      = errors.New("disconnected flights")
)

// An IATA airport code, also known as an IATA location identifier,
// IATA station code, or simply a location identifier,
// is a three-character alphanumeric geocode
var rxAirportCode = regexp.MustCompile(`^[A-Z0-9]{3}$`)

// Request is representing a list of flights provided as input,
// where each flight is represented by an array containing two strings:
// a source[0] and a destination[1] airport code.
type Request [][2]string

// Response is a struct to hold the response data for valid and invalid cases.
type Response struct {
	Status  string   `json:"status"`
	Path    []string `json:"path,omitempty"`
	Message string   `json:"message,omitempty"`
}

// Flight represents a single flight with a source and destination airport.
type Flight struct {
	Source      string
	Destination string
}

// ID returns unique flight ID.
func (f *Flight) ID() string {
	return NewFlightID(f.Source, f.Destination)
}

func (f *Flight) Validate() error {
	if len(f.Source) == 0 || len(f.Destination) == 0 {
		return ErrWrongFlightFormat
	}

	if !rxAirportCode.MatchString(f.Source) || !rxAirportCode.MatchString(f.Destination) {
		return ErrInvalidAirportCode
	}

	if f.Source == f.Destination {
		return ErrSameSourceAndDestination
	}
	return nil
}

// NewFlightID takes a source and destination airport code and generates a unique ID.
func NewFlightID(src, dst string) string {
	var sb strings.Builder
	sb.WriteString(src)
	sb.WriteString(dst)
	return sb.String()
}

// NewFlights takes a Request containing a list of raw flight data,
// and converts it into a slice of Flight structs.
// It returns an error if any flight entry in the raw data has
// an incorrect format, such as empty source or destination airport codes.
// The function also checks for any duplicate flight segments in the Request.
// If a duplicate is found, the function returns ErrDuplicateFlightSegment.
func NewFlights(req Request) (flights []Flight, err error) {
	if len(req) == 0 {
		return nil, ErrNoFlights
	}

	uniqFlights := make(map[string]struct{})

	for _, flight := range req {
		src, dst := flight[0], flight[1]
		flight := Flight{
			Source:      src,
			Destination: dst,
		}
		if err = flight.Validate(); err != nil {
			return nil, err
		}

		flightID := flight.ID()
		if _, found := uniqFlights[flightID]; found {
			return nil, ErrDuplicateFlightSegment
		}
		uniqFlights[flightID] = struct{}{}

		flights = append(flights, flight)
	}

	return
}
