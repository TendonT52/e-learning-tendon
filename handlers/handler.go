package handlers

import "github.com/TendonT52/e-learning-tendon/internal/ports/services"


var appService AppService

type AppService struct {
	userService services.UserService
	jwtService  services.JwtService
}
