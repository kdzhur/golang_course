package main

import (
	"fmt"
	"voting_blockchain/internal/vote"
)

func main() {

	voteArea := vote.NewVotingArea()

	candidates := map[int]string{
		1: "Carlos",
		2: "Alex",
		3: "Rossie",
	}

	srlzdCandidates := vote.InitCandidates(candidates)

	votes := map[string]vote.Candidate{
		"Stephani Voter": srlzdCandidates[0],
		"Antonny Voter":  srlzdCandidates[0],
		"Joseph Voter":   srlzdCandidates[2],
	}

	voteArea.ProcBallots(votes)
	voteArea.ApplyBallots()

	if voteArea.Chain.IsValid() {
		fmt.Println("The Vote is valid!")
	} else {
		fmt.Println("The Vote is not valid!")
	}
}
