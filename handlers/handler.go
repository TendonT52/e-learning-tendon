package handlers

import "github.com/TendonT52/e-learning-tendon/internal/ports/services"

var HandlerConfigInstance *handlerConfig

type Config struct {
	Url            string
	CookieSecure   bool
	CookieHttpOnly bool
}

type handlerConfig struct {
	config  Config
	userApp services.UserService
	jwtApp  services.JwtService
}

func NewHandlerConfig(userApp services.UserService, jwtApp services.JwtService, config Config) {
	HandlerConfigInstance = &handlerConfig{
		userApp: userApp,
		jwtApp:  jwtApp,
		config:  config,
	}
}
