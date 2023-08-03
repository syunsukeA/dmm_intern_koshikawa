package timelines

import (
	"encoding/json"
	_"log"
	"net/http"
	"strconv"
	_"time"

	_"github.com/go-chi/chi/v5"
)

// Handle request for `GET /v1/timeline/home`
func (h *handler) ShowHome(w http.ResponseWriter, r *http.Request) {
}
// Handle request for `GET /v1/timeline/public`
func (h *handler) ShowPublic(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// GETパラメータの取得
	str_only_media := r.URL.Query().Get("only_media")
	str_max_id 	  := r.URL.Query().Get("max_id")
	str_since_id  := r.URL.Query().Get("since_id")
	str_limit 	  := r.URL.Query().Get("limit")
	// 初期値の設定
	if len(str_limit) <= 0 {
		str_limit = "40"
	}
	// エラーハンドリング（paramが指定されていない等の不正リクエストの検出）
	only_media, err := strconv.ParseBool(str_only_media)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	max_id, err := strconv.ParseInt(str_max_id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	since_id, err := strconv.ParseInt(str_since_id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	limit, err := strconv.ParseInt(str_limit, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// 最大値の判定
	if limit > 80 {
		limit = 80
	}
	timeline, err := h.tr.FindByID(ctx, only_media, max_id, since_id, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// FindByID()でtimelineが空の場合は404を返す
	if timeline == nil {
		http.Error(w, "Not Found.", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(timeline); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
