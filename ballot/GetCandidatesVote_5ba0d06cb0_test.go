package main

import (
	"sync"
	"testing"
)

var (
	once                sync.Once
	candidateVotesStore map[string]int
)

func getCandidatesVote() map[string]int {
	once.Do(func() {
		candidateVotesStore = make(map[string]int)
	})
	return candidateVotesStore
}

func TestGetCandidatesVote_5ba0d06cb0(t *testing.T) {
	// Test case 1: Check if the function returns an empty map on first call
	votes := getCandidatesVote()
	if len(votes) != 0 {
		t.Error("Expected empty map on first call, got", votes)
	}

	// Test case 2: Check if the function returns the same map on subsequent calls
	candidateVotesStore["John"] = 5
	votes = getCandidatesVote()
	if len(votes) != 1 {
		t.Error("Expected map with 1 entry, got", votes)
	}
	if votes["John"] != 5 {
		t.Error("Expected 5 votes for John, got", votes["John"])
	}

	// Test case 3: Check if the function returns the updated map after changes
	candidateVotesStore["Doe"] = 10
	votes = getCandidatesVote()
	if len(votes) != 2 {
		t.Error("Expected map with 2 entries, got", votes)
	}
	if votes["Doe"] != 10 {
		t.Error("Expected 10 votes for Doe, got", votes["Doe"])
	}
}
