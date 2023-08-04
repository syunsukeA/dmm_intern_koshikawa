package statuses

import (
	"encoding/json"
	_"log"
	"net/http"

	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/auth"
)

// Request body for `POST /v1/status`
type AddRequest struct {
	AccountID int64
	Content   string
}

// Handle request for `POST /v1/status`
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ao := auth.AccountOf(r)
	var err error

	var req AddRequest
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // BadRequest返すべきなのか？
		return
	}

	status := new(object.Status)
	status.Content = req.Content
	status.AccountID = ao.ID

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
