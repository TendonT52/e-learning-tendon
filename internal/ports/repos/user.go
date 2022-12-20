package repos

import (
	"github.com/TendonT52/e-learning-tendon/internal/core"
)

type UserDB interface {
	InsertUser(user *core.User) error
	InsertManyUser(users []core.User) error
	FindUserByEmail(email string) (core.User, error)
	FindUser(hexID string) (core.User, error)
	FindManyUser(hexIDs []string) ([]core.User, error)
	UpdateUser(user *core.User) error
	DeleteUser(hexID string) error
	DeleteManyUser(hexIDs []string) error
}
