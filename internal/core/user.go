package core

import (
	"time"
)

type User struct {
	ID             string
	FirstName      string
	LastName       string
	Email          string
	HashedPassword string
	Role           string
	Curricula      []string
	UpdatedAt      time.Time
}
