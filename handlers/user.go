package handlers

import (
	"net/http"

	"github.com/TendonT52/e-learning-tendon/handlers/validateService"
	"github.com/TendonT52/e-learning-tendon/internal/application"
	"github.com/TendonT52/e-learning-tendon/internal/pkg/errs"
	"github.com/gin-gonic/gin"
)

func AbortWithHttpError(ctx *gin.Context, err error) {
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
	user, token, err := application.
		UserServiceInstance.
		SignUp(req.FirstName, req.LastName, req.Email, req.Password)
	if err != nil {
		AbortWithHttpError(ctx, err)
		return
	}
	ctx.JSON(
		http.StatusCreated,
		gin.H{
			"id":        user.ID,
			"firstName": user.FirstName,
			"lastName":  user.LastName,
			"email":     user.Email,
			"updateAt":  user.UpdatedAt,
			"token": gin.H{
				"accessToken":  token.Access,
				"refreshToken": token.Refresh,
			},
		},
	)
}
