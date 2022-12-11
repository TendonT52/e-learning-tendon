package db_test

import (
	"errors"
	"testing"
	"time"

	"github.com/TendonT52/e-learning-tendon/db"
	"github.com/TendonT52/e-learning-tendon/internal/pkg/errs"
)

func TestCreateJwtToken(t *testing.T) {
	_, err := db.JwtDBInstance.InsertJwtToken(time.Now().Add(time.Minute))
	if err != nil {
		t.Log(err)
		t.Error("Error while insert jwt to database")
	}
}

func TestGetJwtToken(t *testing.T) {
	id, err := db.JwtDBInstance.InsertJwtToken(time.Now().Add(time.Minute))
	if err != nil {
		t.Log(err)
		t.Error("Error while insert jwt to database")
	}
	err = db.JwtDBInstance.CheckJwtToken(id)
	if err != nil {
		t.Error("Error can't find jwt")
	}
}

func TestGetJwtTokenErrorInvalid(t *testing.T) {
	err := db.JwtDBInstance.CheckJwtToken("fakeId")
	if !errors.Is(err, errs.ErrInvalidToken) {
		t.Error("Wrong error type")
	}
}

func TestGetJwtTokenErrorNotFound(t *testing.T) {
	err := db.JwtDBInstance.CheckJwtToken("63955f26131a9f4401f6dd1f")
	if !errors.Is(err, errs.ErrNotFound) {
		t.Error("Wrong error type")
	}
}

func TestDeleteJwtToken(t *testing.T) {
	id, err := db.JwtDBInstance.InsertJwtToken(time.Now().Add(time.Minute))
	if err != nil {
		t.Log(err)
		t.Error("Error while insert jwt to database")
	}
	err = db.JwtDBInstance.DeleteJwtToken(id)
	if err != nil {
		t.Error("Error while delete jwt token")
	}
}

func TestDeleteJwtTokenErrorNotFound(t *testing.T) {
	err := db.JwtDBInstance.DeleteJwtToken("63955f26131a9f4401f6dd1f")
	if !errors.Is(err, errs.ErrNotFound) {
		t.Error("Wrong error type")
	}
}
