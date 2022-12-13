package core

import (
	"time"
)

const (
	Required  = "required"
	Extension = "extension"
	Optional  = "optional"
)

const (
	Text  = "text"
	Image = "image"
	Video = "video"
	Pdf   = "pdf"
	Sound = "sound"
)

type Node struct {
	ID        string
	Type      string
	Data      string
	CreateBy  string
	UpdatedAt time.Time
}
