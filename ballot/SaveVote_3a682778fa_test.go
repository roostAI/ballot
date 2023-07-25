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
	vote := Vote{CandidateID: "candidate1"}
	err := saveVote(vote)
	if err != nil {
		t.Error("Expected no error, got ", err)
	}

	if candidateVotesStore[vote.CandidateID] != 1 {
		t.Error("Expected 1 vote, got ", candidateVotesStore[vote.CandidateID])
	}

	// Test with multiple votes
	err = saveVote(vote)
	if err != nil {
		t.Error("Expected no error, got ", err)
	}

	if candidateVotesStore[vote.CandidateID] != 2 {
		t.Error("Expected 2 votes, got ", candidateVotesStore[vote.CandidateID])
	}
}
