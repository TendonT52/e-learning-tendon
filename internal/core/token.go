package core

import (
	"time"

)

type Token struct {
	Access  string
	Refresh string
}

type JwtConfig struct {
	AppName              string
	SigningKey           string
	AccesstokenDuration  time.Duration
	RefreshtokenDuration time.Duration
}
