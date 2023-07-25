package main

import (
	"sync"
	"testing"
)

var once sync.Once
var candidateVotesStore map[string]int

func getCandidatesVote() map[string]int {
	once.Do(func() {
		candidateVotesStore = make(map[string]int)
	})
	return candidateVotesStore
} 

func TestGetCandidatesVote_152da079ca(t *testing.T) {
	// Test case 1: Test the function when the candidateVotesStore is not initialized.
	candidateVotesStore = nil
	votes := getCandidatesVote()
	if votes == nil {
		t.Error("Expected a non-nil map, but got nil")
	}

	// Test case 2: Test the function when the candidateVotesStore is already initialized.
	candidateVotesStore = map[string]int{"John": 5, "Doe": 3}
	votes = getCandidatesVote()
	if len(votes) != 2 {
		t.Errorf("Expected a map with 2 elements, but got %d elements", len(votes))
	}
	if votes["John"] != 5 || votes["Doe"] != 3 {
		t.Error("Unexpected vote counts for candidates")
	}
}
