package model

import (
	"errors"
	"strings"
	"time"
)

type Person struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *Person) Validate() error {
	if strings.TrimSpace(p.Name) == "" {
		return errors.New("name cannot be empty")
	}

	return nil
}
