package services

import "github.com/TendonT52/e-learning-tendon/internal/core"

type UserService interface {
	SignUp(string, lastName, email, password string) (core.User, core.Token, error) 
	SignIn(email, password string) (core.User, core.Token, error)
}
