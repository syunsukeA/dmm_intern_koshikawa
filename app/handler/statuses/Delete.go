package statuses

import (
	_"encoding/json"
	_"log"
	"net/http"
	"strconv"

	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/auth"

	"github.com/go-chi/chi/v5"
)

// Handle request for `POST /v1/status`
func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ao := auth.AccountOf(r)
	var err error

	account_id := ao.ID
	str_id := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(str_id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	status := new(object.Status)

	// DBからstatus情報を削除
	status, err = h.sr.DeleteStatus(ctx, id, account_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// DeleteStatus()で削除statusが見つからない場合は404を返す
	if status == nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}
