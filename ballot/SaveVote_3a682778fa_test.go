package main

import (
	"testing"
)

type Vote struct {
	CandidateID string
}

var candidateVotesStore map[string]int

func getCandidatesVote() map[string]int {
	if candidateVotesStore == nil {
		candidateVotesStore = make(map[string]int)
	}
	return candidateVotesStore
}

func saveVote(vote Vote) error {
	candidateVotesStore = getCandidatesVote()
	candidateVotesStore[vote.CandidateID]++
	return nil
}

func TestSaveVote_3a682778fa(t *testing.T) {
	// Test case 1: Single vote for a candidate
	vote := Vote{CandidateID: "candidate1"}
	err := saveVote(vote)
	if err != nil {
		t.Error("Failed to save vote", err)
	}
	if candidateVotesStore["candidate1"] != 1 {
		t.Error("Vote count mismatch. Expected 1, got ", candidateVotesStore["candidate1"])
	}

	// Test case 2: Multiple votes for a candidate
	vote = Vote{CandidateID: "candidate2"}
	for i := 0; i < 5; i++ {
		err = saveVote(vote)
		if err != nil {
			t.Error("Failed to save vote", err)
		}
	}
	if candidateVotesStore["candidate2"] != 5 {
		t.Error("Vote count mismatch. Expected 5, got ", candidateVotesStore["candidate2"])
	}

	// Test case 3: No vote for a candidate
	if candidateVotesStore["candidate3"] != 0 {
		t.Error("Vote count mismatch. Expected 0, got ", candidateVotesStore["candidate3"])
	}
}
