package accounts

import (
	"encoding/json"
	_"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Handle request for `GET /v1/accounts/{username}`
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
		http.Error(w, "Account not found.", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
