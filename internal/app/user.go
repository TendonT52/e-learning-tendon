package app

import (
	"errors"
	"net/http"

	"github.com/TendonT52/e-learning-tendon/internal/core"
	"github.com/TendonT52/e-learning-tendon/internal/pkg/auth"
	"github.com/TendonT52/e-learning-tendon/internal/pkg/errs"
	"github.com/TendonT52/e-learning-tendon/internal/pkg/util"
)

func SignUp(user *core.User, password string) (core.Token, error) {
	_, err := reposInstance.UserDB.FindUserByEmail(user.Email)
	if !errors.Is(err, errs.UserNotFound) {
		if err != nil {
			return core.Token{}, errs.NewHttpError(
				http.StatusInternalServerError,
				"can't find email in database",
			)
		}
		return core.Token{}, errs.NewHttpError(
			http.StatusConflict,
			"email already exists",
		)
	}
	user.HashedPassword = auth.HashPassword(password)
	err = reposInstance.UserDB.InsertUser(user)
	if err != nil {
		return core.Token{}, errs.NewHttpError(
			http.StatusInternalServerError,
			"can't insert user to database",
		)
	}
	token, err := createToken(user.ID)
	if err != nil {
		return core.Token{}, errs.NewHttpError(
			http.StatusInternalServerError,
			"can't create token",
		)
	}
	return token, nil
}

func SignIn(email, password string) (core.User, core.Token, error) {
	user, err := reposInstance.UserDB.FindUserByEmail(email)
	if err != nil {
		if errors.Is(err, errs.UserNotFound) {
			return core.User{}, core.Token{}, errs.NewHttpError(
				http.StatusUnauthorized,
				"invalid email or password",
			)
		}
		return core.User{}, core.Token{}, errs.NewHttpError(
			http.StatusInternalServerError,
			"can't find email in database",
		)
	}
	if !auth.ValidatePassword(password, user.HashedPassword) {
		return core.User{}, core.Token{}, errs.NewHttpError(
			http.StatusUnauthorized,
			"invalid email or password",
		)
	}
	token, err := createToken(user.ID)
	if err != nil {
		return core.User{}, core.Token{}, errs.NewHttpError(
			http.StatusInternalServerError,
			"can't create token",
		)
	}
	return user, token, nil
}

func SignOut(token core.Token) error {
	err := revokeToken(token.Refresh)
	if err != nil {
		return err
	}
	err = revokeToken(token.Access)
	if err != nil {
		return err
	}
	return nil
}

func GetUser(userID string) (core.User, error) {
	user, err := reposInstance.UserDB.FindUser(userID)
	if err != nil {
		if errors.Is(err, errs.UserNotFound) {
			return core.User{}, errs.NewHttpError(
				http.StatusNotFound,
				"user not found",
			)
		}
		if errors.Is(err, errs.InvalidUserID) {
			return core.User{}, errs.NewHttpError(
				http.StatusBadRequest,
				"invalid user id",
			)
		}
		return core.User{}, errs.NewHttpError(
			http.StatusInternalServerError,
			"can't find user in database",
		)
	}
	return user, nil
}

func UpdateUser(user *core.User, password string) error {
	prevUser, err := reposInstance.UserDB.FindUser(user.ID)
	if err != nil {
		if errors.Is(err, errs.UserNotFound) {
			return errs.NewHttpError(
				http.StatusNotFound,
				"user not found",
			)
		}
		if errors.Is(err, errs.InvalidUserID) {
			return errs.NewHttpError(
				http.StatusBadRequest,
				"invalid user id",
			)
		}
		return errs.NewHttpError(
			http.StatusInternalServerError,
			"can't find user in database",
		)
	}
	if password != "" {
		user.HashedPassword = auth.HashPassword(password)
	} else {
		user.HashedPassword = prevUser.HashedPassword
	}
	addedCourse := util.Difference(user.Courses, prevUser.Courses)
	course, err := reposInstance.CourseDB.FindManyCourse(addedCourse)
	if err != nil {
		return errs.NewHttpError(
			http.StatusInternalServerError,
			"can't find course in database",
		)
	}
	if len(course) != len(addedCourse) {
		return errs.NewHttpError(
			http.StatusBadRequest,
			"some course are not found",
		)
	}
	err = reposInstance.UserDB.UpdateUser(user)
	if err != nil {
		return errs.NewHttpError(
			http.StatusInternalServerError,
			"can't update user in database",
		)
	}
	return nil
}

func DeleteUser(userID string) error {
	err := reposInstance.UserDB.DeleteUser(userID)
	if err != nil {
		if errors.Is(err, errs.UserNotFound) {
			return errs.NewHttpError(
				http.StatusNotFound,
				"user not found",
			)
		}
		if errors.Is(err, errs.InvalidUserID) {
			return errs.NewHttpError(
				http.StatusBadRequest,
				"invalid user id",
			)
		}
		return errs.NewHttpError(
			http.StatusInternalServerError,
			"can't delete user in database",
		)
	}
	return nil
}