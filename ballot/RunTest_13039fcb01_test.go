package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Status struct {
	Code    int
	Message string
}

func TestBallot(t *testing.T) {
	t.Run("TestBallot success", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/test", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(runTest)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		expected := `{"Code":200,"Message":"Test Cases passed"}`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})

	t.Run("TestBallot failure", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/test", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(runTest)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}

		expected := `{"Code":400,"Message":"Test Cases Failed with error : error"}`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})
}

func runTest(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement the runTest function
}

func writeVoterResponse(w http.ResponseWriter, status Status) {
	w.Write([]byte(`{"Code":` + string(status.Code) + `,"Message":"` + status.Message + `"}`))
}
