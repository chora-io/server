package cosmos

// EventVoteWithVoter includes the voter address in addition to the proposal id, whereas
// EventVote only includes the proposal id. The voter address is pulled from MsgVote and
// made available in the returned event so that we can query a single vote by voter.
type EventVoteWithVoter struct {
	ProposalId uint64
	Voter      string
}
