package services

import "github.com/TendonT52/e-learning-tendon/internal/core"

type UserService interface {
	SignUp(user *core.User, password string) (core.Token, error)
}
