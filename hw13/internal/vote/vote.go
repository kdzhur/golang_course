package vote

import "voting_blockchain/pkg/blockchain"

type VotingArea struct {
	Chain   *blockchain.Blockchain
	Ballots []Ballot
}

type Ballot struct {
	VoterName string
	Candidate *Candidate
}

type Candidate struct {
	Name   string
	Number int
}

func NewVotingArea() *VotingArea {
	return &VotingArea{
		Chain:   blockchain.CreateBlockchain(3),
		Ballots: []Ballot{},
	}
}

func InitCandidates(candidates map[int]string) []Candidate {
	var srlzdCandidates []Candidate

	for number, name := range candidates {
		srlzdCandidates = append(srlzdCandidates, Candidate{
			Name:   name,
			Number: number,
		})
	}

	return srlzdCandidates
}

func (b *Ballot) handle() (voterName, candidateName string, candidateNumber int) {
	return b.VoterName, b.Candidate.Name, b.Candidate.Number
}

func (v *VotingArea) Vote(b *Ballot) {
	v.Ballots = append(v.Ballots, *b)
}

func (v *VotingArea) ProcBallots(votes map[string]Candidate) {
	for voter, candidate := range votes {
		v.Vote(&Ballot{
			VoterName: voter,
			Candidate: &candidate,
		})
	}
}

func (v *VotingArea) ApplyBallots() {
	for _, b := range v.Ballots {
		v.Chain.AddBlock(b.handle())
	}
}
