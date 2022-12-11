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
	accessToken, err := JwtServiceInstance.generateAccessToken(user.ID)
	if err != nil {
		return core.User{}, core.Token{}, err
	}
	refreshToken, err := JwtServiceInstance.generateRefreshToken(user.ID)
	if err != nil {
		return core.User{}, core.Token{}, err
	}
	token := core.Token{
		Access:  accessToken,
		Refresh: refreshToken,
	}
	return user, token, err
}

// func (us UserService) checkEmailPassword(req core.SignIn) (core.UserResponse, error) {
// 	user, err := us.userDB.GetUserByEmail(req.Email)
// 	if err != nil {
// 		if errors.Is(err, &errs.ErrNotFound) {
// 			return core.UserResponse{},
// 				errs.NewHttpError(
// 					http.StatusBadRequest,
// 					"incorrect email or password",
// 				)
// 		}
// 	}
// 	correct := auth.ValidatePassword(req.Password, user.HashPassword)
// 	if !correct {
// 		return core.UserResponse{},
// 			errs.NewHttpError(
// 				http.StatusForbidden,
// 				"incorrect email or password",
// 			)
// 	}
// 	return core.UserResponse{
// 		Id:        user.Id,
// 		FirstName: user.FirstName,
// 		LastName:  user.LastName,
// 		Email:     user.Email,
// 	}, nil
// }
