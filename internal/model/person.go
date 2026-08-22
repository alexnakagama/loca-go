package model

import "errors"

type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (p Person) Validate() error {
	if p.ID <= 0 {
		return errors.New("ERROR: invalid ID")
	}

	if p.Name == "" {
		return errors.New("ERROR: invalid name")
	}

	return nil
}
