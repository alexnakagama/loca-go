package model

import "errors"

type Location struct {
	PersonID  int     `json:"person_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (l Location) Validate() error {
	if l.PersonID <= 0 {
		return errors.New("ERROR: invalid PersonID")
	}

	if l.Latitude < -90 || l.Latitude > 90 {
		return errors.New("ERROR: invalid Latitude")
	}

	if l.Longitude < -180 || l.Longitude > 180 {
		return errors.New("ERROR: invalid Longitude")
	}

	return nil
}
