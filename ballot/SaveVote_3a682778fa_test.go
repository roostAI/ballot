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
		t.Error("Failed to save vote")
	}

	if candidateVotesStore[vote.CandidateID] != 1 {
		t.Error("Vote count mismatch")
	}

	vote2 := Vote{CandidateID: "candidate2"}
	err = saveVote(vote2)
	if err != nil {
		t.Error("Failed to save vote")
	}

	if candidateVotesStore[vote2.CandidateID] != 1 {
		t.Error("Vote count mismatch")
	}

	err = saveVote(vote)
	if err != nil {
		t.Error("Failed to save vote")
	}

	if candidateVotesStore[vote.CandidateID] != 2 {
		t.Error("Vote count mismatch")
	}
}
