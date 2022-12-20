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
	Courses        []string
	UpdatedAt      time.Time
}
