package model

import (
	"time"
)

type Location struct {
	ID        int       `json:"id"`
	PersonID  int       `json:"person_id"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	CreatedAt time.Time `json:"created_at"`
}
