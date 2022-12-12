package application_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/TendonT52/e-learning-tendon/internal/application"
	"github.com/TendonT52/e-learning-tendon/internal/core"
	"github.com/TendonT52/e-learning-tendon/internal/pkg/errs"
	"github.com/TendonT52/e-learning-tendon/mock"
)

func TestSignUp(t *testing.T) {
	userDB := &mock.MockUserDB{GetByEmail: errs.ErrNotFound}
	jwtDB := &mock.MockJwtDB{}
	config := core.JwtConfig{
		AppName:              "test app name",
		AccessSecret:         "7D6E",
		AccesstokenDuration:  time.Minute,
		RefreshtokenDuration: time.Minute,
	}
	application.NewUserService(userDB)
	application.NewJwtService(jwtDB, config)
	user, token, err := application.UserServiceInstance.SignUp(
		"mock First Name",
		"testLastName",
		"testEmail",
		"testpassword",
	)
	if err != nil {
		t.Error("Error while sign up")
	}
	if user.FirstName != "mock First Name" {
		t.Error("invalid first name")
	}
	if token.Access == "" {
		t.Error("invalid access token")
	}
}

func TestSignIn(t *testing.T) {
	userDB := &mock.MockUserDB{GetByEmail: errs.ErrNotFound}
	jwtDB := &mock.MockJwtDB{}
	config := core.JwtConfig{
		AppName:              "test app name",
		AccessSecret:         "7D6E",
		AccesstokenDuration:  time.Minute,
		RefreshtokenDuration: time.Minute,
	}
	application.NewUserService(userDB)
	application.NewJwtService(jwtDB, config)
	userCreate, _, err := application.UserServiceInstance.SignUp(
		"mock First Name",
		"testLastName",
		"testEmail",
		"testpassword",
	)
	if err != nil {
		t.Error("Error while sign up")
	}
	userDB = &mock.MockUserDB{}
	application.NewUserService(userDB)
	user, _, err := application.UserServiceInstance.SignIn("testEmail", "testpassword")
	if err != nil {
		t.Error("Error while sign in user")
		return
	}
	if !reflect.DeepEqual(userCreate, user) {
		t.Error("user should be equal")
		return
	}
}
