package handler

import (
	"encoding/json"
	"net/http"

	"github.com/alexnakagama/loca-go/internal/model"
	"github.com/alexnakagama/loca-go/internal/store"
)

type PersonHandler struct {
	store *store.Store
}

func NewPersonHandler(store *store.Store) *PersonHandler {
	return &PersonHandler{
		store: store,
	}
}

func (h *PersonHandler) HandlePerson(w http.ResponseWriter, r * http.Request) {
	var person model.Person

	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	
	err = person.Validate()
	if err != nil {
		http.Error(w, "invalid person", http.StatusBadRequest)
		return
	}

	h.store.SetPerson(person)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(person)
}

func (h *PersonHandler) GetPerson(w http.ResponseWriter, r *http.Request) {}
