package http

import (
	json "encoding/json"
	http "net/http"

	"github.com/KhizarShabir1/foodji-tinder/foodji"
	chi "github.com/go-chi/chi/v5"
)

func (s *Server) RegisterVoteRoutes(r chi.Router) {

	r.Post("/votes/store", s.StoreVote())
	r.Get("/votes/session/{sessionID}", s.GetVotesForSession())
	r.Get("/votes/aggscores", s.GetAggregatedScores())
}

// GetAggregatedScores returns a handler for the "GET /votes/aggscores"
func (s *Server) GetAggregatedScores() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		scores, err := s.VoteProvider.GetAggregatedScores()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(scores)
	}
}

// GetVotesForSession returns a handler for the "GET /votes/session/{sessionID}" route.
func (s *Server) GetVotesForSession() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionID := chi.URLParam(r, "sessionID")

		votes, err := s.VoteProvider.GetVotesForSession(sessionID)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		// Marshal votes to JSON and write response
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(votes); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
}

// GetVotesForSession returns a handler for the "POST /votes/store" route.
func (s *Server) StoreVote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()

		var vote foodji.Vote
		err := decoder.Decode(&vote)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		err = s.VoteProvider.StoreVote(&vote)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		s.writeJSON(w, http.StatusCreated, vote)
	}
}
