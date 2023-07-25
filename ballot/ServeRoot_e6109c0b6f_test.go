package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type Vote struct {
	VoterID     string
	CandidateID string
}

func serveRoot(w http.ResponseWriter, r *http.Request) {
	// Mock function to be implemented
}

func TestServeRoot_e6109c0b6f(t *testing.T) {
	t.Run("Test GET Method", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(serveRoot)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		expected := `{"Results":[{"CandidateID":"1","Votes":2},{"CandidateID":"2","Votes":1}],"TotalVotes":3}`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
		}
	})

	t.Run("Test POST Method", func(t *testing.T) {
		vote := &Vote{
			VoterID:     "1",
			CandidateID: "1",
		}
		jsonVote, _ := json.Marshal(vote)
		req, err := http.NewRequest("POST", "/", bytes.NewBuffer(jsonVote))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(serveRoot)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
		}

		expected := `{"Code":201,"Message":"Vote saved sucessfully"}`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
		}
	})

	t.Run("Test Invalid Method", func(t *testing.T) {
		req, err := http.NewRequest("PUT", "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(serveRoot)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
		}

		expected := `{"Code":405,"Message":"Bad Request. Vote can not be saved"}`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
		}
	})
}
