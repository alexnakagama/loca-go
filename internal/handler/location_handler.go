package handler

import (
	"encoding/json"
	"net/http"

	"github.com/alexnakagama/loca-go/internal/model"
	"github.com/alexnakagama/loca-go/internal/store"
)

type LocationHandler struct {
	store *store.Store
}

func NewLocationHandler(store *store.Store) *LocationHandler {
	return &LocationHandler{
		store: store,
	}
}

func (h *LocationHandler) HandleLocation(w http.ResponseWriter, r *http.Request) {
	var location model.Location

	err := json.NewDecoder(r.Body).Decode(&location)
	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	err = location.Validate()
	if err != nil {
		http.Error(w, "invalid location", http.StatusBadRequest)
		return
	}

	h.store.SetLocation(location)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(location)
}

func (h *LocationHandler) HandleGetLocation(w http.ResponseWriter, r *http.Request) {
}
