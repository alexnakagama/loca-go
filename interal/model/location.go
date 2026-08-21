package model

type Location struct {
	PersonID  int     `json:"person_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
