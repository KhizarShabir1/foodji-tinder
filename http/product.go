package http

import (
	"database/sql"
	json "encoding/json"
	"errors"
	"io/ioutil"
	"log"
	http "net/http"

	"github.com/KhizarShabir1/foodji-tinder/foodji"
	chi "github.com/go-chi/chi/v5"
)

func (s *Server) RegisterProductRoutes(r chi.Router) {

	r.Get("/products/{productID}", s.GetProduct())
	r.Get("/products", s.ListProducts())
	r.Post("/products/create", s.CreateProduct())
}

// ListProducts returns a handler for the "GET /products" route.
func (s *Server) ListProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get all products from the provider
		products, count, err := s.ProductProvider.ListProducts()
		if err != nil {
			// If there was an error retrieving the products from the database, return a 500 error
			s.writeError(w, http.StatusInternalServerError, ErrorResponse{
				Message: "Internal server error",
			})
			return
		}

		// If the products were found, write them to the response
		s.writeJSON(w, http.StatusOK, struct {
			Count    int              `json:"count"`
			Products []foodji.Product `json:"products"`
		}{
			Count:    count,
			Products: products,
		})
	}
}

// GetProduct returns a handler for the "GET /products/{productID}" route.
func (s *Server) GetProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "productID")

		product, err := s.ProductProvider.GetProduct(id)
		//log.Println(*product)
		if err != nil {
			// If the product wasn't found, return a 404 error with a custom error message
			if errors.Is(err, sql.ErrNoRows) {
				s.writeError(w, http.StatusNotFound, ErrorResponse{
					Message: "Product not found",
				})
				return
			}

			// If there was an error retrieving the product from the database, return a 500 error
			s.writeError(w, http.StatusInternalServerError, ErrorResponse{
				Message: "Internal server error",
			})
			return
		}

		// If the product was found, write it to the response
		s.writeJSON(w, http.StatusOK, *product)
	}
}

// CreateProduct returns a handler for the "POST /products/create route.
func (s *Server) CreateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var product foodji.Product
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Error reading request body:", err)
			return
		}
		defer r.Body.Close()
		err = json.Unmarshal(body, &product)
		if err != nil {
			log.Println("Error unmarshalling JSON:", err)
			return
		}

		product, _ = s.ProductProvider.CreateProduct(product)
		s.writeJSON(w, http.StatusOK, product)
	}
}
