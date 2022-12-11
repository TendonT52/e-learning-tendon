package mock

import (
	"time"

	"github.com/TendonT52/e-learning-tendon/internal/core"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockUserDB struct {
	Insert     error
	GetByEmail error
	GetById    error
}

func (u *MockUserDB) InsertUser(firstName, lastName, email, hashPassword, role string) (core.User, error) {
	switch u.Insert {
	case nil:
		user := core.User{
			ID:           primitive.NewObjectID().Hex(),
			FirstName:    "mock First Name",
			LastName:     "mock Last Name",
			Email:        "mock@email.com",
			HashPassword: "mockHashPassword",
			Role:         core.Student,
			UpdatedAt:    time.Now(),
		}
		return user, nil
	default:
		return core.User{}, u.Insert
	}
}

func (u *MockUserDB) GetUserByEmail(email string) (core.User, error) {
	switch u.GetByEmail {
	case nil:
		user := core.User{
			ID:           primitive.NewObjectID().Hex(),
			FirstName:    "mock First Name",
			LastName:     "mock Last Name",
			Email:        "mock@email.com",
			HashPassword: "mockHashPassword",
			Role:         core.Student,
			UpdatedAt:    time.Now(),
		}
		return user, nil
	default:
		return core.User{}, u.GetByEmail
	}
}

func (u *MockUserDB) GetUserById(id string) (core.User, error) {
	switch u.GetById {
	case nil:
		user := core.User{
			ID:           primitive.NewObjectID().Hex(),
			FirstName:    "mock First Name",
			LastName:     "mock Last Name",
			Email:        "mock@email.com",
			HashPassword: "mockHashPassword",
			Role:         core.Student,
			UpdatedAt:    time.Now(),
		}
		return user, nil
	default:
		return core.User{}, u.GetById
	}
}
