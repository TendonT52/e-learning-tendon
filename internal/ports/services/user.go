package services

import "github.com/TendonT52/e-learning-tendon/internal/core"

type UserService interface {
	CreateUser(core.User) (string, error)
	GetUserByEmail(string) (core.User, error)
	GetUserById(string) (core.User, error)
}
