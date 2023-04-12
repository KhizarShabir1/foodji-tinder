package database

import (
	"log"

	"github.com/KhizarShabir1/foodji-tinder/foodji"
)

var _ foodji.VoteProvider = (*Provider)(nil)

func (p *Provider) StoreVote(vote *foodji.Vote) error {
	_, err := p.db.Exec("INSERT INTO vote (session_id, product_id, liked) VALUES ($1, $2, $3)", vote.SessionID, vote.ProductID, vote.Liked)
	if err != nil {
		return err
	}

	return nil
}

func (p *Provider) GetVotesForSession(sessionID string) ([]*foodji.Vote, error) {
	// Prepare query to retrieve votes for the given session ID
	stmt := `
		SELECT id, session_id, product_id, liked
		FROM vote
		WHERE session_id = $1;
	`

	// Execute the query
	rows, err := p.db.Query(stmt, sessionID)
	if err != nil {
		log.Println("Error retrieving votes for session:", err)
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows and create Vote objects for each row
	votes := []*foodji.Vote{}
	for rows.Next() {
		vote := &foodji.Vote{}
		err := rows.Scan(&vote.ID, &vote.SessionID, &vote.ProductID, &vote.Liked)
		if err != nil {
			log.Println("Error scanning row for vote:", err)
			return nil, err
		}
		votes = append(votes, vote)
	}

	return votes, nil
}

func (p *Provider) GetAggregatedScores() (map[string]float64, error) {
	rows, err := p.db.Query("SELECT product_id, AVG(CAST(liked AS INT)) FROM vote GROUP BY product_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	scores := make(map[string]float64)
	for rows.Next() {
		var productID string
		var avgScore float64
		if err := rows.Scan(&productID, &avgScore); err != nil {
			return nil, err
		}
		scores[productID] = avgScore
	}

	return scores, nil
}
