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
	// Test case 1: Check if the function returns a non-nil map
	votes := getCandidatesVote()
	if votes == nil {
		t.Error("Expected a non-nil map, got nil")
	}

	// Test case 2: Check if the function returns an empty map on first call
	if len(votes) != 0 {
		t.Errorf("Expected an empty map, got a map with size: %d", len(votes))
	}

	// Test case 3: Check if the function returns the same map on subsequent calls
	candidateVotesStore["John"] = 10
	votes = getCandidatesVote()
	if len(votes) != 1 {
		t.Errorf("Expected a map with size: 1, got a map with size: %d", len(votes))
	}
	if votes["John"] != 10 {
		t.Errorf("Expected John to have 10 votes, got: %d", votes["John"])
	}

	// Test case 4: Check if the function returns the updated map
	candidateVotesStore["Jane"] = 20
	votes = getCandidatesVote()
	if len(votes) != 2 {
		t.Errorf("Expected a map with size: 2, got a map with size: %d", len(votes))
	}
	if votes["Jane"] != 20 {
		t.Errorf("Expected Jane to have 20 votes, got: %d", votes["Jane"])
	}
}
