package core

import (
	"time"
)

const (
	Admin   = "admin"
	Teacher = "teacher"
	Student = "student"
)

type User struct {
	ID           string
	FirstName    string
	LastName     string
	Email        string
	HashPassword string
	Role         string
	UpdatedAt    time.Time
}
