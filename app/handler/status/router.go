package status

import (
	"net/http"
	"yatter-backend-go/app/domain/repository"

	"github.com/go-chi/chi/v5"
)

// Implementation of handler
type handler struct {
	sr repository.Status
}

// Create Handler for `/v1/accounts/`
func NewRouter(sr repository.Status) http.Handler {
	r := chi.NewRouter()

	h := &handler{sr}
	r.Post("/", h.Create)
	r.Get("/{id}", h.Show)

	return r
}
