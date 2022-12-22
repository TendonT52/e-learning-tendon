package handlers

import (
	"net/http"

	"github.com/TendonT52/e-learning-tendon/handlers/validateService"
	"github.com/TendonT52/e-learning-tendon/internal/app"
	"github.com/TendonT52/e-learning-tendon/internal/core"
	"github.com/gin-gonic/gin"
)

func GetNodeHandler(ctx *gin.Context) {
	nodeID := ctx.Param("id")
	lesson, err := app.GetNode(nodeID)
	if err != nil {
		abortWithHttpError(ctx, err)
		return
	}
	ctx.AbortWithStatusJSON(
		http.StatusOK,
		gin.H{
			"id":       lesson.ID,
			"type":     lesson.Type,
			"data":     lesson.Data,
			"createBy": lesson.CreateBy,
			"updateAt": lesson.UpdatedAt,
		},
	)
}

func PatchNodeHandler(ctx *gin.Context) {
	nodeID := ctx.Param("id")
	node, err := app.GetNode(nodeID)
	if err != nil {
		abortWithHttpError(ctx, err)
		return
	}
	var reqBody struct {
		Data string `json:"data" binding:"omitempty"`
		Type string `json:"description" binding:"omitempty"`
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
	err = app.UpdateNode(&node)
	if err != nil {
		abortWithHttpError(ctx, err)
		return
	}
	ctx.AbortWithStatusJSON(
		http.StatusOK,
		gin.H{
			"id":       node.ID,
			"data":     node.Data,
			"type":     node.Type,
			"createBy": node.CreateBy,
			"updateAt": node.UpdatedAt,
		},
	)
}

func DeleteNodeHandler(ctx *gin.Context) {
	nodeID := ctx.Param("id")
	err := app.DeleteNode(nodeID)
	if err != nil {
		abortWithHttpError(ctx, err)
		return
	}
	ctx.AbortWithStatus(http.StatusOK)
}
func PostNodeHandler(ctx *gin.Context) {
	var reqBody struct {
		Data string `json:"data" binding:"omitempty"`
		Type string `json:"description" binding:"omitempty"`
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

	node := core.Node{
		Type: reqBody.Type,
		Data: reqBody.Data,
	}
	err = app.CreateNode(&node)
	if err != nil {
		abortWithHttpError(ctx, err)
		return
	}
	ctx.AbortWithStatusJSON(
		http.StatusOK,
		gin.H{
			"id":       node.ID,
			"data":     node.Data,
			"type":     node.Type,
			"createBy": node.CreateBy,
			"updateAt": node.UpdatedAt,
		},
	)
}
