package core

import "time"

type Course struct {
	ID          string
	Name        string
	Description string
	Access      string
	CreateBy    string
	UpdatedAt   time.Time
	Lessons     []string
}
