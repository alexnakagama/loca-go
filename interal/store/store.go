package store

import "github.com/alexnakagama/loca-go/interal/model"

type Store struct {
	locations map[int]model.Location
}

func NewStore() *Store {
	return &Store{
		locations: make(map[int]model.Location),
	}
}
