package handlers

import (
	"net/http"

	"github.com/TendonT52/e-learning-tendon/handlers/validateService"
	"github.com/TendonT52/e-learning-tendon/internal/app"
	"github.com/gin-gonic/gin"
)

func GetUserHandler(ctx *gin.Context) {
	userID := ctx.Param("id")
	user, err := app.GetUser(userID)
	if err != nil {
		abortWithHttpError(ctx, err)
		return
	}
	ctx.AbortWithStatusJSON(
		http.StatusOK,
		gin.H{
			"id":        user.ID,
			"firstName": user.FirstName,
			"lastName":  user.LastName,
			"email":     user.Email,
			"updateAt":  user.UpdatedAt,
		},
	)
}

func PatchUserHandler(ctx *gin.Context) {
	userID := ctx.Param("id")
	user, err := app.GetUser(userID)
	if err != nil {
		abortWithHttpError(ctx, err)
		return
	}
	var reqBody struct {
		FirstName string   `json:"firstName" binding:"omitempty,min=1,max=40"`
		LastName  string   `json:"lastName" binding:"omitempty,min=1,max=40"`
		Email     string   `json:"email" binding:"omitempty,email"`
		Password  string   `json:"password" binding:"omitempty,min=8,max=99"`
		Role      string   `json:"role"`
		Courses   []string `json:"courses"`
	}
	err = ctx.ShouldBindJSON(&reqBody)
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
	if reqBody.FirstName != "" {
		user.FirstName = reqBody.FirstName
	}
	if reqBody.LastName != "" {
		user.LastName = reqBody.LastName
	}
	if reqBody.Email != "" {
		user.Email = reqBody.Email
	}
	if reqBody.Role != "" {
		user.Role = reqBody.Role
	}
	if reqBody.Courses != nil {
		user.Courses = reqBody.Courses
	}

	err = app.UpdateUser(&user, reqBody.Password)
	if err != nil {
		abortWithHttpError(ctx, err)
		return
	}
	ctx.AbortWithStatusJSON(
		http.StatusOK,
		gin.H{
			"id":        user.ID,
			"firstName": user.FirstName,
			"lastName":  user.LastName,
			"email":     user.Email,
			"updateAt":  user.UpdatedAt,
		},
	)
}

func DeleteUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	err := app.DeleteUser(userID)
	if err != nil {
		abortWithHttpError(ctx, err)
		return
	}
	ctx.AbortWithStatus(http.StatusOK)
}
