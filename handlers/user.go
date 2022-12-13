package handlers

import (
	"net/http"

	"github.com/TendonT52/e-learning-tendon/handlers/validateService"
	"github.com/TendonT52/e-learning-tendon/internal/application"
	"github.com/TendonT52/e-learning-tendon/internal/pkg/errs"
	"github.com/gin-gonic/gin"
)

func (hc *handlerConfig) AbortWithHttpError(ctx *gin.Context, err error) {
	if httpErr, ok := err.(errs.HttpError); ok {
		ctx.AbortWithStatusJSON(
			httpErr.Code,
			gin.H{
				"message": httpErr.Message,
			},
		)
	} else {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "something went wrong",
			},
		)
	}
}

func (hc *handlerConfig) SignUpHandler(ctx *gin.Context) {
	type signUpReq struct {
		FirstName string `json:"firstName" binding:"required,min=1,max=40"`
		LastName  string `json:"lastName" binding:"required,min=1,max=40"`
		Email     string `json:"email" binding:"required,email"`
		Password  string `json:"password" binding:"required,min=8,max=99"`
	}
	req := signUpReq{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		message := validateService.TranslateError(err)
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": "validation error",
				"error":   message,
			},
		)
		return
	}
	user, token, err := application.
		UserServiceInstance.
		SignUp(req.FirstName, req.LastName, req.Email, req.Password)
	if err != nil {
		hc.AbortWithHttpError(ctx, err)
		return
	}
	ctx.SetCookie("token",
		token.Access,
		int(application.JwtServiceInstance.GetCookieDuration().Seconds()),
		"/auth",
		hc.config.Url,
		hc.config.CookieSecure,
		hc.config.CookieHttpOnly)
	ctx.AbortWithStatusJSON(
		http.StatusCreated,
		gin.H{
			"id":           user.ID,
			"firstName":    user.FirstName,
			"lastName":     user.LastName,
			"email":        user.Email,
			"updateAt":     user.UpdatedAt,
			"refreshToken": token.Refresh,
		},
	)
}

func (hc *handlerConfig) SignInHandler(ctx *gin.Context) {
	type signInReq struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8,max=99"`
	}
	req := signInReq{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		message := validateService.TranslateError(err)
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": "validation error",
				"error":   message,
			},
		)
		return
	}
	user, token, err := application.UserServiceInstance.SignIn(req.Email, req.Password)
	if err != nil {
		hc.AbortWithHttpError(ctx, err)
		return
	}
	ctx.SetCookie("token",
		token.Access,
		int(application.JwtServiceInstance.GetCookieDuration().Seconds()),
		"/auth",
		hc.config.Url,
		hc.config.CookieSecure,
		hc.config.CookieHttpOnly)
	ctx.AbortWithStatusJSON(
		http.StatusOK,
		gin.H{
			"id":           user.ID,
			"firstName":    user.FirstName,
			"lastName":     user.LastName,
			"email":        user.Email,
			"updateAt":     user.UpdatedAt,
			"refreshToken": token.Refresh,
		},
	)
}

func (hc *handlerConfig) SignOutHandler(ctx *gin.Context) {
	type signOutReq struct {
		RefreshToken string `json:"refreshToken" binding:"required"`
	}
	req := signOutReq{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		message := validateService.TranslateError(err)
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": "validation error",
				"error":   message,
			},
		)
		return
	}
	accessToken, err := ctx.Cookie("token")
	application.UserServiceInstance.SignOut(accessToken, req.RefreshToken)
	ctx.AbortWithStatusJSON(
		http.StatusOK,
		gin.H{
			"message": "Sign out complete",
		},
	)
}
