package db_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/TendonT52/e-learning-tendon/db"
	"github.com/TendonT52/e-learning-tendon/internal/core"
	"github.com/TendonT52/e-learning-tendon/internal/pkg/errs"
)

func TestCreateUser(t *testing.T) {
	_, err := db.UserDBInstance.InsertUser(
		"testFirstName",
		"testLastName",
		"testEmail",
		"testHashPassword",
		core.Student)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetUserByEmailFound(t *testing.T) {
	user, err := db.UserDBInstance.InsertUser(
		"testFirstName",
		"testLastName",
		"testEmailFound",
		"testHashPassword",
		core.Student)
	if err != nil {
		t.Error(err.Error())
	}
	result, err := db.UserDBInstance.GetUserByEmail("testEmailFound")
	if err != nil {
		t.Error(err.Error())
		return
	}
	if !reflect.DeepEqual(user, result) {
		t.Error("user and result should be equal")
		return
	}
}

func TestGetUserByEmailNotFound(t *testing.T) {
	_, err := db.UserDBInstance.InsertUser(
		"testFirstName",
		"testLastName",
		"testEmail",
		"testHashPassword",
		core.Student)
	if err != nil {
		t.Error(err.Error())
	}
	_, err = db.UserDBInstance.GetUserByEmail("testNotFoundEmaill@email.com")
	if err == nil {
		t.Error("err should not nil")
	}
	if !errors.Is(err, errs.ErrNotFound) {
		t.Error("errs should be error not found")
	}
}

func TestGetUserByIdFound(t *testing.T) {
	user, err := db.UserDBInstance.InsertUser(
		"testFirstName",
		"testLastName",
		"testEmail",
		"testHashPassword",
		core.Student)
	if err != nil {
		t.Error(err.Error())
	}
	result, err := db.UserDBInstance.GetUserById(user.ID)
	if err != nil {
		t.Error(err.Error())
	}
	user.ID = result.ID
	user.UpdatedAt = result.UpdatedAt
	if !reflect.DeepEqual(user, result) {
		t.Error("user and result should be equal")
	}
}

func TestGetUserByIdNotFound(t *testing.T) {
	_, err := db.UserDBInstance.InsertUser(
		"testFirstName",
		"testLastName",
		"testEmail",
		"testHashPassword",
		core.Student)
	if err != nil {
		t.Error(err.Error())
	}
	_, err = db.UserDBInstance.GetUserById("6391ba3a9d33a9f976ad899e")
	if err == nil {
		t.Error("error should not nil")
	}
	if !errors.Is(err, errs.ErrNotFound) {
		t.Error("errs should be error not found")
	}
}
