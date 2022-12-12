package application

import (
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/TendonT52/e-learning-tendon/internal/core"
	"github.com/TendonT52/e-learning-tendon/internal/pkg/errs"
	"github.com/TendonT52/e-learning-tendon/mock"
)

func TestGenerateAccessTokenJwt(t *testing.T) {
	config := core.JwtConfig{
		AppName:              "test app name",
		AccessSecret:         "7D6E",
		AccesstokenDuration:  time.Minute,
		RefreshtokenDuration: time.Minute,
	}
	db := &mock.MockJwtDB{}
	NewJwtService(db, config)
	_, err := JwtServiceInstance.generateAccessToken("userID")
	if err != nil {
		t.Error("error while generate jwt")
		return
	}
}

func TestGenerateAccessTokenJwtErrDatabase(t *testing.T) {
	config := core.JwtConfig{
		AppName:              "test app name",
		AccessSecret:         "7D6E",
		AccesstokenDuration:  time.Minute,
		RefreshtokenDuration: time.Minute,
	}
	db := &mock.MockJwtDB{Insert: errs.ErrDatabase}
	NewJwtService(db, config)
	_, err := JwtServiceInstance.generateAccessToken("userID")
	httpErr, ok := err.(errs.HttpError)
	if !ok {
		t.Error("Wrong error type")
		return
	}
	if !httpErr.Is(errs.NewHttpError(
		http.StatusInternalServerError,
		"Error while create token",
	)) {
		t.Error("Wrong error")
	}
}

func TestGenerateRefreshTokenJwt(t *testing.T) {
	config := core.JwtConfig{
		AppName:              "test app name",
		AccessSecret:         "7D6E",
		AccesstokenDuration:  time.Minute,
		RefreshtokenDuration: time.Minute,
	}
	db := &mock.MockJwtDB{}
	NewJwtService(db, config)
	_, err := JwtServiceInstance.generateRefreshToken("userID")
	if err != nil {
		t.Error("error while generate jwt")
		return
	}
}

func TestGenerateRefreshTokenJwtErrDatabase(t *testing.T) {
	config := core.JwtConfig{
		AppName:              "test app name",
		AccessSecret:         "7D6E",
		AccesstokenDuration:  time.Minute,
		RefreshtokenDuration: time.Minute,
	}
	db := &mock.MockJwtDB{Insert: errs.ErrDatabase}
	NewJwtService(db, config)
	_, err := JwtServiceInstance.generateRefreshToken("userID")
	if !errors.Is(err, errs.NewHttpError(
		http.StatusInternalServerError,
		"Error while create token",
	)) {
		t.Error("error in wrong type")
	}
}

func TestValidateToken(t *testing.T) {
	config := core.JwtConfig{
		AppName:              "test app name",
		AccessSecret:         "7D6E",
		AccesstokenDuration:  time.Minute,
		RefreshtokenDuration: time.Minute,
	}
	db := &mock.MockJwtDB{}
	NewJwtService(db, config)
	encodedToken, err := JwtServiceInstance.generateAccessToken("userID")
	if err != nil {
		t.Error("error while generate access token")
	}
	claim, err := JwtServiceInstance.validateToken(encodedToken, "userID")
	if err != nil {
		t.Error("error while validate token")
	}
	if claim.Subject != "userID" {
		t.Error("Wrong claim")
	}
}

func TestValidateTokenErrTokenExpired(t *testing.T) {
	config := core.JwtConfig{
		AppName:              "test app name",
		AccessSecret:         "7D6E",
		AccesstokenDuration:  0,
		RefreshtokenDuration: 0,
	}
	db := &mock.MockJwtDB{}
	NewJwtService(db, config)
	encodedToken, err := JwtServiceInstance.generateAccessToken("userID")
	if err != nil {
		t.Error("error while generate access token")
	}
	_, err = JwtServiceInstance.validateToken(encodedToken, "userID")
	if !errors.Is(err, errs.NewHttpError(
		http.StatusConflict,
		"token is expired",
	)) {
		t.Error("error in wrong type")
	}
}

func TestValidateTokenErrNotFound(t *testing.T) {
	config := core.JwtConfig{
		AppName:              "test app name",
		AccessSecret:         "7D6E",
		AccesstokenDuration:  time.Minute,
		RefreshtokenDuration: time.Minute,
	}
	db := &mock.MockJwtDB{
		Check: errs.ErrNotFound,
	}
	NewJwtService(db, config)
	encodedToken, err := JwtServiceInstance.generateAccessToken("userID")
	if err != nil {
		t.Error("error while generate access token")

	}
	NewJwtService(db, config)
	_, err = JwtServiceInstance.validateToken(encodedToken, "userID")
	if !errors.Is(err, errs.NewHttpError(
		http.StatusConflict,
		"token already revoke",
	)) {
		t.Error("error in wrong type")
	}
}
