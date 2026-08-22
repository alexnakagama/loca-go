package handler

import (
	"net/http"

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

func (h *PersonHandler) HandlePerson(w http.ResponseWriter, r * http.Request) {}
