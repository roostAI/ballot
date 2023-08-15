package main

import (
	"testing"
)

// Mock Vote struct
type Vote struct {
	CandidateID string
}

// Mock candidateVotesStore
var candidateVotesStore map[string]int

// Mock getCandidatesVote function
func getCandidatesVote() map[string]int {
	return candidateVotesStore
}

// Method to be tested
func saveVote(vote Vote) error {
	candidateVotesStore = getCandidatesVote()
	candidateVotesStore[vote.CandidateID]++
	return nil
}

// Test method
func TestSaveVote_3f5729642d(t *testing.T) {
	// Test case 1: When vote is saved successfully
	candidateVotesStore = make(map[string]int)
	vote := Vote{CandidateID: "123"}
	err := saveVote(vote)
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	if candidateVotesStore[vote.CandidateID] != 1 {
		t.Error("Expected 1, got ", candidateVotesStore[vote.CandidateID])
	}

	// Test case 2: When vote is saved for the second time
	err = saveVote(vote)
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	if candidateVotesStore[vote.CandidateID] != 2 {
		t.Error("Expected 2, got ", candidateVotesStore[vote.CandidateID])
	}
}
