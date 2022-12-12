package mock

import (
	"time"

	"github.com/TendonT52/e-learning-tendon/internal/core"
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
			ID:           "6396adedc0dc98e37661989f",
			FirstName:    "mock First Name",
			LastName:     "mock Last Name",
			Email:        "mock@email.com",
			HashPassword: "$2a$14$wgpnGTeGygi6imo4VCPk4.B45P/KZMqqQsD65R5xGTrH9Kbr4DTze",
			Role:         core.Student,
			UpdatedAt:    time.Unix(1122777834, 0),
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
			ID:           "6396adedc0dc98e37661989f",
			FirstName:    "mock First Name",
			LastName:     "mock Last Name",
			Email:        "mock@email.com",
			HashPassword: "$2a$14$wgpnGTeGygi6imo4VCPk4.B45P/KZMqqQsD65R5xGTrH9Kbr4DTze",
			Role:         core.Student,
			UpdatedAt:    time.Unix(1122777834, 0),
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
			ID:           "6396adedc0dc98e37661989f",
			FirstName:    "mock First Name",
			LastName:     "mock Last Name",
			Email:        "mock@email.com",
			HashPassword: "mockHashPassword",
			Role:         core.Student,
			UpdatedAt:    time.Unix(1122777834, 0),
		}
		return user, nil
	default:
		return core.User{}, u.GetById
	}
}
