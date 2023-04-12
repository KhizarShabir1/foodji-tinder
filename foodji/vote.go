package foodji

// VoteProvider gives access to votes
type VoteProvider interface {
	StoreVote(vote *Vote) error
	GetVotesForSession(sessionID string) ([]*Vote, error)
	GetAggregatedScores() (map[string]float64, error)
}

type Vote struct {
	ID        int    `json:"id"`
	SessionID string `json:"sessionId"`
	ProductID string `json:"productId"`
	Liked     bool   `json:"liked"`
}
