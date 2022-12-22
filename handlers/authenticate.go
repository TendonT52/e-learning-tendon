package handlers

import (
	"net/http"

	"github.com/TendonT52/e-learning-tendon/handlers/validateService"
	"github.com/TendonT52/e-learning-tendon/internal/app"
	"github.com/TendonT52/e-learning-tendon/internal/core"
	"github.com/gin-gonic/gin"
)

func SignUpHandler(ctx *gin.Context) {
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

	user := core.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Role:      core.Student,
		Courses:   []string{},
	}
	token, err := app.SignUp(&user, req.Password)
	if err != nil {
		abortWithHttpError(ctx, err)
		return
	}

	ctx.SetCookie("refresh_token",
		token.Refresh,
		int(config.RefreshTokenDuration),
		"/refresh",
		config.Url,
		config.RefreshCookieSecure,
		config.RefreshCookieHttpOnly,
	)
	ctx.AbortWithStatusJSON(
		http.StatusCreated,
		gin.H{
			"id":          user.ID,
			"firstName":   user.FirstName,
			"lastName":    user.LastName,
			"email":       user.Email,
			"updateAt":    user.UpdatedAt,
			"accessToken": token.Access,
		},
	)
}

func SignInHandler(ctx *gin.Context) {
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

	user, token, err := app.SignIn(req.Email, req.Password)
	if err != nil {
		abortWithHttpError(ctx, err)
		return
	}

	ctx.SetCookie("refresh_token",
		token.Refresh,
		int(config.RefreshTokenDuration),
		"/refresh",
		config.Url,
		config.RefreshCookieSecure,
		config.RefreshCookieHttpOnly,
	)

	ctx.AbortWithStatusJSON(
		http.StatusOK,
		gin.H{
			"id":          user.ID,
			"firstName":   user.FirstName,
			"lastName":    user.LastName,
			"email":       user.Email,
			"updateAt":    user.UpdatedAt,
			"accessToken": token.Access,
		},
	)
}

func SignOutHandler(ctx *gin.Context) {
	type signInReq struct {
		AccessToken string `json:"accessToken" binding:"required"`
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
	refreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"message": "Unauthorized",
			},
		)
		return
	}

	app.SignOut(core.Token{
		Access:  req.AccessToken,
		Refresh: refreshToken,
	})
	ctx.AbortWithStatusJSON(
		http.StatusOK,
		gin.H{
			"message": "Sign out complete",
		},
	)
}

func RefreshTokenHandler(ctx *gin.Context) {
	refreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"message": "Unauthorized",
			},
		)
		return
	}
	token, err := app.RefreshToken(refreshToken)
	if err != nil {
		abortWithHttpError(ctx, err)
		return
	}

	ctx.SetCookie("refresh_token",
		token.Refresh,
		int(config.RefreshTokenDuration),
		"/refresh",
		config.Url,
		config.RefreshCookieSecure,
		config.RefreshCookieHttpOnly,
	)

	ctx.AbortWithStatusJSON(
		http.StatusOK,
		gin.H{
			"accessToken": token.Access,
		},
	)
}
