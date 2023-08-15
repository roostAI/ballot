package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Vote struct {
	VoterID     string
	CandidateID string
}

type MockSaveVote struct{}

func (m *MockSaveVote) saveVote(vote Vote) error {
	if vote.VoterID == "" || vote.CandidateID == "" {
		return errors.New("Vote is not valid. Vote can not be saved")
	}
	return nil
}

func serveRoot(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
	case "POST":
		var vote Vote
		err := json.NewDecoder(r.Body).Decode(&vote)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		m := MockSaveVote{}
		err = m.saveVote(vote)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func TestServeRoot_da512bfd0e(t *testing.T) {
	// Test case 1: Successful GET request
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

	// Test case 2: Successful POST request
	vote := Vote{VoterID: "voter1", CandidateID: "candidate1"}
	jsonVote, _ := json.Marshal(vote)
	req, err = http.NewRequest("POST", "/", bytes.NewBuffer(jsonVote))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(serveRoot)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Test case 3: Unsuccessful POST request with invalid vote
	vote = Vote{VoterID: "", CandidateID: ""}
	jsonVote, _ = json.Marshal(vote)
	req, err = http.NewRequest("POST", "/", bytes.NewBuffer(jsonVote))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(serveRoot)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	// Test case 4: Unsupported method
	req, err = http.NewRequest("PUT", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(serveRoot)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}
}
