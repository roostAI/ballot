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
		t.Error("Expected a non-nil map, but got nil")
	}

	// Test case 2: Check if the map is empty
	if len(votes) != 0 {
		t.Errorf("Expected an empty map, but got a map of length %d", len(votes))
	}

	// Test case 3: Check if the function returns the same map on subsequent calls
	votes["candidate1"] = 10
	votes2 := getCandidatesVote()
	if len(votes2) != 1 {
		t.Errorf("Expected a map of length 1, but got a map of length %d", len(votes2))
	}
	if votes2["candidate1"] != 10 {
		t.Errorf("Expected candidate1 to have 10 votes, but got %d votes", votes2["candidate1"])
	}
}
