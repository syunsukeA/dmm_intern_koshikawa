package status

import (
	"encoding/json"
	"net/http"

	"yatter-backend-go/app/domain/object"
)

// Request body for `POST /v1/status`
type AddRequest struct {
	Content string
}

// Handle request for `POST /v1/status`
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error

	var req AddRequest
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // BadRequest返すべきなのか？
		return
	}

	status := new(object.Status)
	status.Content = req.Content

	// DBにstatus情報を保存
	status, err = h.sr.SaveStatus(ctx, status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
