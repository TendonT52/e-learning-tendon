package core

import "time"

type Lesson struct {
	ID          string
	Name        string
	Description string
	Access      string
	CreateBy    string
	UpdatedAt   time.Time
	Nodes       []string
	NextLessons []string
	PrevLessons []string
}