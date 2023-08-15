package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRunTest_a2bb790205(t *testing.T) {
	t.Run("Test case passed", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/test", bytes.NewBuffer([]byte{}))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(runTest)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		expected := `{"Message":"Test Cases passed","Code":200}`
		if rr.Body.String() != expected {
			t.Errorf("Handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})

	t.Run("Test case failed", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/test", bytes.NewBuffer([]byte{}))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(runTest)

		// TODO: Make TestBallot return an error to simulate a failure
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}

		expected := `{"Message":"Test Cases Failed with error : test error","Code":400}`
		if rr.Body.String() != expected {
			t.Errorf("Handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})
}
