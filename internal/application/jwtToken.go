package application

import (
	"errors"
	"net/http"
	"time"

	"github.com/TendonT52/e-learning-tendon/internal/core"
	"github.com/TendonT52/e-learning-tendon/internal/pkg/errs"
	"github.com/TendonT52/e-learning-tendon/internal/ports/repos"
	"github.com/golang-jwt/jwt/v4"
)

var JwtServiceInstance *jwtService

type jwtService struct {
	jwtDB  repos.JwtDB
	config core.JwtConfig
}

func NewJwtService(jwtDB repos.JwtDB, config core.JwtConfig) {
	JwtServiceInstance = &jwtService{
		jwtDB:  jwtDB,
		config: config,
	}
}

func (jw *jwtService) GetCookieDuration() time.Duration {
	return jw.config.AccesstokenDuration
}

func (jw *jwtService) createToken(userID string) (core.Token, error) {
	accessToken, err := JwtServiceInstance.generateAccessToken(userID)
	if err != nil {
		return core.Token{}, err
	}
	refreshToken, err := JwtServiceInstance.generateRefreshToken(userID)
	if err != nil {
		return core.Token{}, err
	}
	token := core.Token{
		Access:  accessToken,
		Refresh: refreshToken,
	}
	return token, nil
}

func (js *jwtService) generateAccessToken(userID string) (string, error) {
	exp := time.Now().Add(js.config.AccesstokenDuration)
	claim := jwt.RegisteredClaims{
		Subject:   userID,
		Issuer:    js.config.AppName,
		ExpiresAt: jwt.NewNumericDate(exp),
		NotBefore: jwt.NewNumericDate(time.Now()),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	id, err := js.jwtDB.InsertJwtToken(time.Now().Add(js.config.AccesstokenDuration))
	if err != nil {
		return id, errs.NewHttpError(
			http.StatusInternalServerError,
			"Error while create token",
		)
	}
	claim.ID = id
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	ss, err := token.SignedString([]byte(js.config.AccessSecret))
	if err != nil {
		return id, errs.NewHttpError(
			http.StatusInternalServerError,
			"Error while signing token",
		)
	}
	return ss, nil
}

func (js *jwtService) generateRefreshToken(userID string) (string, error) {
	exp := time.Now().Add(js.config.RefreshtokenDuration)
	claim := jwt.RegisteredClaims{
		Subject:   userID,
		Issuer:    js.config.AppName,
		ExpiresAt: jwt.NewNumericDate(exp),
		NotBefore: jwt.NewNumericDate(time.Now()),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	id, err := js.jwtDB.InsertJwtToken(time.Now().Add(js.config.AccesstokenDuration))
	if err != nil {
		return id, errs.NewHttpError(
			http.StatusInternalServerError,
			"Error while create token",
		)
	}
	claim.ID = id
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	ss, err := token.SignedString([]byte(js.config.AccessSecret))
	if err != nil {
		return id, errs.NewHttpError(
			http.StatusInternalServerError,
			"Error while signing token",
		)
	}
	return ss, nil
}

func (js *jwtService) ValidateToken(encodedToken string) (jwt.RegisteredClaims, error) {
	claim := jwt.RegisteredClaims{}
	_, err := jwt.ParseWithClaims(
		encodedToken,
		&claim,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(js.config.AccessSecret), nil
		},
	)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return claim, errs.NewHttpError(
				http.StatusConflict,
				"token is expired",
			)
		} else {
			return claim, errs.NewHttpError(
				http.StatusConflict,
				"token not valid",
			)
		}
	}
	err = js.jwtDB.CheckJwtToken(claim.ID)
	if err != nil {
		if errors.Is(err, errs.NotFound) {
			return claim, errs.NewHttpError(
				http.StatusConflict,
				"token already revoke",
			)
		} else {
			return claim, errs.NewHttpError(
				http.StatusInternalServerError,
				"can't find token in database",
			)
		}
	}
	return claim, nil
}

func (js *jwtService) revokeToken(encodedToken string) error {
	claim, err := js.ValidateToken(encodedToken)
	if err != nil {
		return err
	}
	err = js.jwtDB.DeleteJwtToken(claim.ID)
	if err != nil {
		if errors.Is(err, errs.NotFound) {
			return errs.NewHttpError(
				http.StatusConflict,
				"token already revoke",
			)
		} else {
			return errs.NewHttpError(
				http.StatusInternalServerError,
				"can't find token in database",
			)
		}
	}
	return nil
}
