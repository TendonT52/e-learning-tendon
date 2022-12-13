package handlers

import (
	"net/http"

	"github.com/TendonT52/e-learning-tendon/internal/application"
	"github.com/gin-gonic/gin"
)

func (hc *handlerConfig) Auth() gin.HandlerFunc {
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
		claim, err := application.JwtServiceInstance.ValidateToken(accessToken)
		if err != nil {
			hc.AbortWithHttpError(ctx, err)
			return
		}
		ctx.Set("userID", claim.ID)
		ctx.Next()
	}
}
