package handlers

import "time"

var config Config

type Config struct {
	Port 				string
	Url                   string
	AccessCookieSecure    bool
	AccessCookieHttpOnly  bool
	AccessTokenDuration   time.Duration
	RefreshCookieSecure   bool
	RefreshCookieHttpOnly bool
	RefreshTokenDuration  time.Duration
}

func SetConfig(con Config) {
	config = con
}
