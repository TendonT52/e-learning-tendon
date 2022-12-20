package app

import (
	"errors"
	"net/http"
	"time"

	"github.com/TendonT52/e-learning-tendon/internal/core"
	"github.com/TendonT52/e-learning-tendon/internal/pkg/errs"
	"github.com/golang-jwt/jwt/v4"
)

func GetCookieDuration() time.Duration {
	return appConfig.AccesstokenDuration
}

func createToken(userID string) (core.Token, error) {
	accessToken, err := generateAccessToken(userID)
	if err != nil {
		return core.Token{}, err
	}
	refreshToken, err := generateRefreshToken(userID)
	if err != nil {
		return core.Token{}, err
	}
	token := core.Token{
		Access:  accessToken,
		Refresh: refreshToken,
	}
	return token, nil
}

func generateAccessToken(userID string) (string, error) {
	exp := time.Now().Add(appConfig.AccesstokenDuration)
	claim := jwt.RegisteredClaims{
		Subject:   userID,
		Issuer:    appConfig.AppName,
		ExpiresAt: jwt.NewNumericDate(exp),
		NotBefore: jwt.NewNumericDate(time.Now()),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	id, err := reposInstance.JwtDB.InsertJwt(time.Now().Add(appConfig.AccesstokenDuration))
	if err != nil {
		return id, errs.NewHttpError(
			http.StatusInternalServerError,
			"Error while create token",
		)
	}
	claim.ID = id
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	ss, err := token.SignedString([]byte(appConfig.AccessSecret))
	if err != nil {
		return id, errs.NewHttpError(
			http.StatusInternalServerError,
			"Error while signing token",
		)
	}
	return ss, nil
}

func generateRefreshToken(userID string) (string, error) {
	exp := time.Now().Add(appConfig.RefreshtokenDuration)
	claim := jwt.RegisteredClaims{
		Subject:   userID,
		Issuer:    appConfig.AppName,
		ExpiresAt: jwt.NewNumericDate(exp),
		NotBefore: jwt.NewNumericDate(time.Now()),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	id, err := reposInstance.JwtDB.InsertJwt(time.Now().Add(appConfig.AccesstokenDuration))
	if err != nil {
		return id, errs.NewHttpError(
			http.StatusInternalServerError,
			"Error while create token",
		)
	}
	claim.ID = id
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	ss, err := token.SignedString([]byte(appConfig.AccessSecret))
	if err != nil {
		return id, errs.NewHttpError(
			http.StatusInternalServerError,
			"Error while signing token",
		)
	}
	return ss, nil
}

func ValidateToken(encodedToken string) (jwt.RegisteredClaims, error) {
	claim := jwt.RegisteredClaims{}
	_, err := jwt.ParseWithClaims(
		encodedToken,
		&claim,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(appConfig.AccessSecret), nil
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
	err = reposInstance.JwtDB.CheckJwt(claim.ID)
	if err != nil {
		if errors.Is(err, errs.TokenNotfound) {
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

func revokeToken(encodedToken string) error {
	claim, err := ValidateToken(encodedToken)
	if err != nil {
		return err
	}
	err = reposInstance.JwtDB.DeleteJwt(claim.ID)
	if err != nil {
		if errors.Is(err, errs.TokenNotfound) {
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

func RefreshToken(encodedToken string) (core.Token, error) {
	claim, err := ValidateToken(encodedToken)
	if err != nil {
		return core.Token{}, err
	}
	err = revokeToken(encodedToken)
	if err != nil {
		return core.Token{}, err
	}
	token, err := createToken(claim.Subject)
	if err != nil {
		return core.Token{}, err
	}
	return token, nil
}