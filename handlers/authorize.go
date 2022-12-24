package handlers

import (
	"fmt"
	"net/http"

	"github.com/TendonT52/e-learning-tendon/internal/app"
	"github.com/TendonT52/e-learning-tendon/internal/core"
	"github.com/gin-gonic/gin"
)

func HaveAccessToken(ctx *gin.Context) {
	accessToken := ctx.GetHeader("Authorization")
	if len(accessToken) <= 7 {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"message": "unauthorized",
			},
		)
		return
	}
	accessToken = fmt.Sprint(accessToken[7:])
	claim, err := app.ValidateToken(accessToken)
	if err != nil {
		abortWithHttpError(ctx, err)
		return
	}
	ctx.Set("userID", claim.Subject)
}

func IsStudent(ctx *gin.Context) {
}

func IsAdmin(ctx *gin.Context) {
	userID := ctx.GetString("userID")
	user, err := app.GetUser(userID)
	if err != nil {
		abortWithHttpError(ctx, err)
		return
	}
	if user.Role != core.Admin {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"message": "unauthorized",
			},
		)
		return
	}
}

func IsTeacher(ctx *gin.Context) {
	userID := ctx.GetString("userID")
	user, err := app.GetUser(userID)
	if err != nil {
		abortWithHttpError(ctx, err)
		return
	}
	if user.Role == core.Student {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"message": "unauthorized",
			},
		)
		return
	}
}
