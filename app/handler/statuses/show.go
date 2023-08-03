package statuses

import (
	"encoding/json"
	_ "log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// Handle request for `GET /v1/status/{id}`
func (h *handler) Show(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	str_id := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(str_id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	status, err := h.sr.FindByID(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// FindByID()でstatusが見つからない場合は404を返す
	if status == nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
