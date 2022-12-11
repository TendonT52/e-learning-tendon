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
	NodeId    string
	Type      string
	Data      string
	Priority  string
	UpdatedAt time.Time
}
