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
	AccessSecret         string
	RefreshSecret         string
	AccesstokenDuration  time.Duration
	RefreshtokenDuration time.Duration
}
