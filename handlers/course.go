package handlers

import (
	"net/http"

	"github.com/TendonT52/e-learning-tendon/handlers/validateService"
	"github.com/TendonT52/e-learning-tendon/internal/app"
	"github.com/TendonT52/e-learning-tendon/internal/core"
	"github.com/gin-gonic/gin"
)

func GetCourseHandler(ctx *gin.Context) {
	courseID := ctx.Param("id")
	course, err := app.GetCourse(courseID)
	if err != nil {
		abortWithHttpError(ctx, err)
		return
	}
	ctx.AbortWithStatusJSON(
		http.StatusOK,
		gin.H{
			"id":          course.ID,
			"name":        course.Name,
			"description": course.Description,
			"access":      course.Access,
			"createBy":    course.CreateBy,
			"updateAt":    course.UpdatedAt,
			"lessons":     course.Lessons,
		},
	)
}
func PatchCourseHandler(ctx *gin.Context) {
	courseID := ctx.Param("id")
	course, err := app.GetCourse(courseID)
	if err != nil {
		abortWithHttpError(ctx, err)
		return
	}
	var reqBody struct {
		Name        string   `json:"name" binding:"omitempty,min=1,max=40"`
		Description string   `json:"description" binding:"omitempty"`
		Access      string   `json:"access" binding:"omitempty"`
		CreateBy    string   `json:"createBy" binding:"omitempty"`
		Lessons     []string `json:"lessons" binding:"omitempty"`
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
	if reqBody.Name != "" {
		course.Name = reqBody.Name
	}
	if reqBody.Description != "" {
		course.Description = reqBody.Description
	}
	if reqBody.Access != "" {
		course.Access = reqBody.Access
	}
	if reqBody.CreateBy != "" {
		course.CreateBy = reqBody.CreateBy
	}
	if reqBody.Lessons != nil {
		reqBody.Lessons = course.Lessons
	}
	err = app.UpdateCourse(&course)
	if err != nil {
		abortWithHttpError(ctx, err)
		return
	}
	ctx.AbortWithStatusJSON(
		http.StatusOK,
		gin.H{
			"id":          course.ID,
			"name":        course.Name,
			"description": course.Description,
			"access":      course.Access,
			"createBy":    course.CreateBy,
			"updateAt":    course.UpdatedAt,
			"lessons":     course.Lessons,
		},
	)
}

func DeleteCourseHandler(ctx *gin.Context) {
	courseID := ctx.Param("id")
	err := app.DeleteCourse(courseID)
	if err != nil {
		abortWithHttpError(ctx, err)
		return
	}
	ctx.AbortWithStatus(http.StatusOK)
}
func PostCourseHandler(ctx *gin.Context) {
	var reqBody struct {
		Name        string   `json:"name" binding:"required,min=1,max=40"`
		Description string   `json:"description" binding:"omitempty"`
		Access      string   `json:"access" binding:"required"`
		Lessons     []string `json:"lessons" binding:"omitempty"`
	}
	err := ctx.ShouldBindJSON(&reqBody)
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

	course := core.Course{
		Name:        reqBody.Name,
		Description: reqBody.Description,
		Access:      reqBody.Access,
		CreateBy:    ctx.GetString("userID"),
		Lessons:     reqBody.Lessons,
	}
	err = app.CreateCourse(&course)
	if err != nil {
		abortWithHttpError(ctx, err)
		return
	}
	ctx.AbortWithStatusJSON(
		http.StatusCreated,
		gin.H{
			"id":          course.ID,
			"name":        course.Name,
			"description": course.Description,
			"access":      course.Access,
			"createBy":    course.CreateBy,
			"updateAt":    course.UpdatedAt,
			"lessons":     course.Lessons,
		},
	)
}
