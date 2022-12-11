package repos

import (
	"github.com/TendonT52/e-learning-tendon/internal/core"
)

type UserDB interface {
	InsertUser(firstName, lastName, email, hashPassword, role string) (core.User, error)
	GetUserByEmail(email string) (core.User, error)
	GetUserById(id string) (core.User, error)
}
