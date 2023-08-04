package accounts

import (
	"encoding/json"
	"log"
	"net/http"

	"yatter-backend-go/app/handler/auth"

	"github.com/go-chi/chi/v5"
)

// Handle request for `GET /v1/accounts/{username}`
func (h *handler) Follow(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// ao情報の取得
	follower_ao, err := h.ar.FindByUsername(ctx, chi.URLParam(r, "username"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// FindByUsername()でaccountが見つからない場合は404を返す
	if follower_ao == nil {
		http.Error(w, "Account not found.", http.StatusNotFound)
		return
	}
	followee_ao := auth.AccountOf(r)
	if follower_ao == nil {
		http.Error(w, "Account not found.", http.StatusBadRequest)
		return
	}

	// ToDo: 適正な戻り値の型を定義する
	log.Println(followee_ao, follower_ao)
	esc, err := h.ar.AddFollow(ctx, followee_ao, follower_ao)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// AddFollow()で最初の戻り値がnilの場合404を返す
	if esc == nil {
		http.Error(w, "Account not found.", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(esc); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
