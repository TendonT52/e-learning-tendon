package handlers

import (
	"net/http"

	"github.com/TendonT52/e-learning-tendon/handlers/validateService"
	"github.com/TendonT52/e-learning-tendon/internal/app"
	"github.com/TendonT52/e-learning-tendon/internal/core"
	"github.com/gin-gonic/gin"
)

func GetLessonHandler(ctx *gin.Context) {
	lessonID := ctx.Param("id")
	lesson, err := app.GetLesson(lessonID)
	if err != nil {
		abortWithHttpError(ctx, err)
		return
	}
	ctx.AbortWithStatusJSON(
		http.StatusOK,
		gin.H{
			"id":          lesson.ID,
			"name":        lesson.Name,
			"description": lesson.Description,
			"access":      lesson.Access,
			"createBy":    lesson.CreateBy,
			"updateAt":    lesson.UpdatedAt,
			"nodes":       lesson.Nodes,
			"nextLesson":  lesson.NextLessons,
			"prevLesson":  lesson.PrevLessons,
		},
	)
}

func PatchLessonHandler(ctx *gin.Context) {
	lessonID := ctx.Param("id")
	lesson, err := app.GetLesson(lessonID)
	if err != nil {
		abortWithHttpError(ctx, err)
		return
	}
	var reqBody struct {
		Name        string   `json:"name" binding:"omitempty,min=1,max=40"`
		Description string   `json:"description" binding:"omitempty"`
		Access      string   `json:"access" binding:"omitempty"`
		CreateBy    string   `json:"createBy" binding:"omitempty"`
		Nodes       []string `json:"nodes" binding:"omitempty"`
		NextLessons []string `json:"nextLesson" binding:"omitempty"`
		PrevLessons []string `json:"prevLesson" binding:"omitempty"`
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
		lesson.Name = reqBody.Name
	}
	if reqBody.Description != "" {
		lesson.Description = reqBody.Description
	}
	if reqBody.Access != "" {
		lesson.Access = reqBody.Access
	}
	if reqBody.CreateBy != "" {
		lesson.CreateBy = reqBody.CreateBy
	}
	if reqBody.Nodes != nil {
		lesson.Nodes = reqBody.Nodes
	}
	if reqBody.NextLessons != nil {
		lesson.NextLessons = reqBody.NextLessons
	}
	if reqBody.PrevLessons != nil {
		lesson.PrevLessons = reqBody.PrevLessons
	}

	err = app.UpdateLesson(&lesson)
	if err != nil {
		abortWithHttpError(ctx, err)
		return
	}
	ctx.AbortWithStatusJSON(
		http.StatusOK,
		gin.H{
			"id":          lesson.ID,
			"name":        lesson.Name,
			"description": lesson.Description,
			"access":      lesson.Access,
			"createBy":    lesson.CreateBy,
			"updateAt":    lesson.UpdatedAt,
			"nodes":       lesson.Nodes,
			"nextLesson":  lesson.NextLessons,
			"prevLesson":  lesson.PrevLessons,
		},
	)
}

func DeleteLessonHandler(ctx *gin.Context) {
	lessonID := ctx.Param("id")
	err := app.DeleteLesson(lessonID)
	if err != nil {
		abortWithHttpError(ctx, err)
		return
	}
	ctx.AbortWithStatus(http.StatusOK)
}
func PostLessonHandler(ctx *gin.Context) {
	var reqBody struct {
		Name        string   `json:"name" binding:"required,min=1,max=40"`
		Description string   `json:"description" binding:"required"`
		Access      string   `json:"access" binding:"required"`
		Nodes       []string `json:"nodes" binding:"required"`
		NextLessons []string `json:"nextLesson" binding:"required"`
		PrevLessons []string `json:"prevLesson" binding:"required"`
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

	lesson := core.Lesson{
		Name:        reqBody.Name,
		Description: reqBody.Description,
		Access:      reqBody.Access,
		CreateBy:    ctx.GetString("userID"),
		Nodes:       reqBody.Nodes,
		NextLessons: reqBody.NextLessons,
		PrevLessons: reqBody.PrevLessons,
	}
	err = app.CreateLesson(&lesson)
	if err != nil {
		abortWithHttpError(ctx, err)
		return
	}
	ctx.AbortWithStatusJSON(
		http.StatusCreated,
		gin.H{
			"id":          lesson.ID,
			"name":        lesson.Name,
			"description": lesson.Description,
			"access":      lesson.Access,
			"createBy":    lesson.CreateBy,
			"updateAt":    lesson.UpdatedAt,
			"nodes":       lesson.Nodes,
			"nextLesson":  lesson.NextLessons,
			"prevLesson":  lesson.PrevLessons,
		},
	)
}
