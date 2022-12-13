package application

import (
	"errors"
	"net/http"

	"github.com/TendonT52/e-learning-tendon/internal/core"
	"github.com/TendonT52/e-learning-tendon/internal/pkg/auth"
	"github.com/TendonT52/e-learning-tendon/internal/pkg/errs"
	"github.com/TendonT52/e-learning-tendon/internal/ports/repos"
)

var UserServiceInstance *UserService

type UserService struct {
	userDB repos.UserDB
}

func NewUserService(userDB repos.UserDB) {
	UserServiceInstance = &UserService{
		userDB: userDB,
	}
}

func (us *UserService) SignUp(firstName, lastName, email, password string) (core.User, core.Token, error) {
	_, err := us.userDB.GetUserByEmail(email)
	if !errors.Is(err, errs.ErrNotFound) {
		return core.User{},
			core.Token{},
			errs.NewHttpError(
				http.StatusConflict,
				"email already exists")
	}
	hashPassword := auth.HashPassword(password)
	user, err := us.userDB.InsertUser(firstName, lastName, email, hashPassword, core.Student)
	if err != nil {
		return core.User{}, core.Token{}, errs.NewHttpError(
			http.StatusInternalServerError,
			"error while add user to database",
		)
	}
	token, err := JwtServiceInstance.createToken(user.ID)
	if err != nil {
		return core.User{}, core.Token{}, err
	}
	return user, token, nil
}

func (us *UserService) SignIn(email, password string) (core.User, core.Token, error) {
	user, err := us.userDB.GetUserByEmail(email)
	if err != nil {
		return core.User{}, core.Token{}, errs.NewHttpError(
			http.StatusConflict,
			"wrong email or password",
		)
	}
	ok := auth.ValidatePassword(password, user.HashPassword)
	if !ok {
		return core.User{}, core.Token{}, errs.NewHttpError(
			http.StatusConflict,
			"wrong email or password",
		)
	}
	token, err := JwtServiceInstance.createToken(user.ID)
	if err != nil {
		return core.User{}, core.Token{}, err
	}
	return user, token, nil
}

func (us *UserService) SignOut(accessToken, refreshToken string) error {
	err := JwtServiceInstance.revokeToken(refreshToken)
	if err != nil {
		return err
	}
	err = JwtServiceInstance.revokeToken(accessToken)
	if err != nil {
		return err
	}
	return nil
}
