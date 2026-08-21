package repository

import "github.com/alexnakagama/loca-go/interal/model"

type PersonRepository interface {
	Create(person *model.Person) error
	GetByID(id int) (*model.Person, error)
	GetAll() ([]model.Person, error)
	Delete(id int) error
}
