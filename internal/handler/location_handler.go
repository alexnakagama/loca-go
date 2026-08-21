package handler

import (
	"encoding/json"
	"fmt"
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

	h.store.SetLocation(location)

	fmt.Println(location)
}
