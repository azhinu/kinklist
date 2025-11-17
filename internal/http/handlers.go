package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/azhinu/kinklist/internal/model"
	"github.com/azhinu/kinklist/internal/storage"
)

type Handler struct {
    db *storage.DB
}

func NewHandler(db *storage.DB) *Handler {
    return &Handler{db: db}
}

// GET /api/kinklist/{id}
func (h *Handler) GetKinkList(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")

    kl, err := h.db.GetKinkList(id)
    if err != nil {
        http.Error(w, "not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(kl)
}

// PUT /api/kinklist/{id}
func (h *Handler) UpdateKinkList(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")

    var kl model.KinkList

    if err := json.NewDecoder(r.Body).Decode(&kl); err != nil {
        http.Error(w, "bad json", http.StatusBadRequest)
        return
    }

    kl.ID = id
    kl.UpdatedAt = time.Now()

    if err := h.db.SaveKinkList(&kl); err != nil {
        http.Error(w, "db error", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(&kl)
}
