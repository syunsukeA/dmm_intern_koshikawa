package statuses

import (
	"net/http"
	"yatter-backend-go/app/domain/repository"
	"yatter-backend-go/app/handler/auth"

	"github.com/go-chi/chi/v5"
)

// Implementation of handler
type handler struct {
	sr repository.Status
}

// Create Handler for `/v1/status/`
func NewRouter(ar repository.Account, sr repository.Status) http.Handler {
	r := chi.NewRouter()

	h := &handler{sr}
	r.With(auth.Middleware(ar)).Post("/", h.Create)
	r.With(auth.Middleware(ar)).Delete("/{id}", h.Delete)
	r.Get("/{id}", h.Show)

	return r
}
