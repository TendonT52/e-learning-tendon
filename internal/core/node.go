package core

import (
	"time"
)

type Node struct {
	ID        string
	Type      string
	Data      string
	CreateBy  string
	UpdatedAt time.Time
}

