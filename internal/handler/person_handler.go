package handler

import "github.com/alexnakagama/loca-go/internal/store"

type PersonHandler struct {
	store *store.Store
}

func NewPersonHandler(store *store.Store) *PersonHandler {} 
