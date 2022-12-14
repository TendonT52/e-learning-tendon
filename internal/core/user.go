package core

import (
	"time"
)


type User struct {
	ID           string
	FirstName    string
	LastName     string
	Email        string
	HashPassword string
	Role         string
	UpdatedAt    time.Time
	Curricula    []string
}
