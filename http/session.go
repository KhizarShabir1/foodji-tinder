package http

import (
	http "net/http"

	chi "github.com/go-chi/chi/v5"
)

func (s *Server) RegisterSessionRoutes(r chi.Router) {

	r.Post("/sessions/create", s.CreateSession())
	r.Get("/sessions/all", s.GetSessions())
}

func (s *Server) CreateSession() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Insert the new session into the database

		session, err := s.SessionProvider.CreateSession()
		if err != nil {
			s.writeError(w, http.StatusInternalServerError, ErrorResponse{
				Message: "Failed to create session",
			})
			return
		}

		// Return the newly created session to the client
		s.writeJSON(w, http.StatusCreated, session)
	}
}

func (s *Server) GetSessions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get all sessions from the database
		sessions, err := s.SessionProvider.GetSessions()
		if err != nil {
			s.writeError(w, http.StatusInternalServerError, ErrorResponse{
				Message: "Failed to get sessions",
			})
			return
		}

		// Return the sessions to the client
		s.writeJSON(w, http.StatusOK, sessions)
	}
}
