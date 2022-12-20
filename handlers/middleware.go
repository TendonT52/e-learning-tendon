package handlers

import (
	"net/http"

	"github.com/TendonT52/e-learning-tendon/internal/app"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken, err := ctx.Cookie("token")
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"message": "can't get access token error",
				},
			)
			return
		}
		claim, err := app.ValidateToken(accessToken)
		if err != nil {
			abortWithHttpError(ctx, err)
			return
		}
		ctx.Set("userID", claim.ID)
		ctx.Next()
	}
}
