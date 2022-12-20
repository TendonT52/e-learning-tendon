package handlers

import (
	"net/http"

	"github.com/TendonT52/e-learning-tendon/internal/pkg/errs"
	"github.com/gin-gonic/gin"
)

func abortWithHttpError(ctx *gin.Context, err error) {
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
