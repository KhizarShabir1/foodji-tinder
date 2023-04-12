package http

import (
	json "encoding/json"
	log "log"
	http "net/http"

	"github.com/KhizarShabir1/foodji-tinder/database"
	"github.com/KhizarShabir1/foodji-tinder/foodji"
	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

// Server represents an HTTP server. It wraps all HTTP functionality.
type Server struct {
	Router *chi.Mux

	ProductProvider foodji.ProductProvider
	SessionProvider foodji.SessionProvider
	VoteProvider    foodji.VoteProvider
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewServer(provider *database.Provider) *Server {
	server := &Server{
		Router: chi.NewRouter(),

		ProductProvider: provider,
		SessionProvider: provider,
		VoteProvider:    provider,
	}

	server.Router.Use(middleware.Logger)

	server.Router.Route("/", func(r chi.Router) {

		server.RegisterProductRoutes(r)
		server.RegisterSessionRoutes(r)
		server.RegisterVoteRoutes(r)
	})

	return server
}

// Start begins listening for HTTP requests on the configured port
func (s *Server) Start() error {
	srv := &http.Server{
		Handler: s.Router,
		Addr:    ":8080",
	}
	err := srv.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Add("Conent-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println("Failed to write response to client", zap.Error(err))
	}
}

func (s *Server) writeError(w http.ResponseWriter, statusCode int, err ErrorResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(err)
}
