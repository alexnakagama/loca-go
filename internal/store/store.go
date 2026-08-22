package store

import (
	"errors"

	"github.com/alexnakagama/loca-go/internal/model"
)

type Store struct {
	locations map[int]model.Location
	persons map[int]model.Person
}

func NewStore() *Store {
	return &Store{
		locations: make(map[int]model.Location),
		persons: make(map[int]model.Person),
	}
}

func (s *Store) SetLocation(location model.Location) {
	s.locations[location.PersonID] = location
}

func (s *Store) SetPerson(person model.Person) {
	s.persons[person.ID] = person
}

func (s *Store) GetPerson(id int) (model.Person, error) {
	person, ok := s.persons[id]
	if !ok {
		return model.Person{}, errors.New("person not found")
	}

	return person, nil
}
