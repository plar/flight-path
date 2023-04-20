package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestCalculateHandler(t *testing.T) {
	// Disable logging output during tests
	log.SetOutput(ioutil.Discard)

	tests := []struct {
		name           string
		method         string
		payload        []byte
		expectedStatus int
		expectedRes    Response
	}{
		{
			name:           "Invalid Method",
			method:         "GET",
			payload:        nil,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedRes: Response{
				Status:  "error",
				Message: "Method not allowed",
			},
		},
		{
			name:           "Invalid JSON",
			method:         "POST",
			payload:        []byte(`{"bad": "json}`),
			expectedStatus: http.StatusBadRequest,
			expectedRes: Response{
				Status:  "error",
				Message: "Invalid JSON",
			},
		},
		{
			name:           "Invalid Flights",
			method:         "POST",
			payload:        []byte(`[["AUS","LAX"],["MIA"],["JFK","LGA"]]`),
			expectedStatus: http.StatusBadRequest,
			expectedRes: Response{
				Status:  "error",
				Message: "Invalid request format",
			},
		},
		{
			name:           "Same Source and Desination airport",
			method:         "POST",
			payload:        []byte(`[["SFO", "ATL"],["ATL", "EWR"],["EWR", "SFO"]]`),
			expectedStatus: http.StatusBadRequest,
			expectedRes: Response{
				Status:  "error",
				Message: "Failed to find flight path",
			},
		},
		{
			name:           "Valid Flights",
			method:         "POST",
			payload:        []byte(`[["AUS","LAX"],["LAX","JFK"]]`),
			expectedStatus: http.StatusOK,
			expectedRes: Response{
				Status: "success",
				Path:   []string{"AUS", "JFK"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, "/calculate", bytes.NewBuffer(tt.payload))
			if err != nil {
				t.Fatalf("Error creating request: %v", err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(calculate)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("Handler returned wrong status code: got %v, want %v", status, tt.expectedStatus)
			}

			var res Response
			err = json.Unmarshal(rr.Body.Bytes(), &res)
			if err != nil {
				t.Errorf("Error unmarshalling response body: %v", err)
			}

			if res.Status != tt.expectedRes.Status || res.Message != tt.expectedRes.Message || !reflect.DeepEqual(res.Path, tt.expectedRes.Path) {
				t.Errorf("Handler returned wrong response: got %v, want %v", res, tt.expectedRes)
			}
		})
	}
}
