package accounts

import (
	"encoding/json"
	_"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Request body for `POST /v1/accounts`
// type ShowRequest struct {
// 	Username string
// 	Password string
// }

// Handle request for `POST /v1/accounts`
func (h *handler) Show(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	username := chi.URLParam(r, "username")

	account, err := h.ar.FindByUsername(ctx, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// FindByUsername()でaccountが見つからない場合は404を返す
	if account == nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
