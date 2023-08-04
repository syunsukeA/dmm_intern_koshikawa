package timelines

import (
	"net/http"
	"yatter-backend-go/app/domain/repository"
	"yatter-backend-go/app/handler/auth"

	"github.com/go-chi/chi/v5"
)

// Implementation of handler
type handler struct {
	tr repository.Timeline
}

// Create Handler for `/v1/status/`
func NewRouter(ar repository.Account, tr repository.Timeline) http.Handler {
	r := chi.NewRouter()

	h := &handler{tr}
	r.With(auth.Middleware(ar)).Get("/home", h.ShowHome)
	r.Get("/public", h.ShowPublic)

	return r
}
