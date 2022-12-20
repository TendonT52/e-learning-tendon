package handlers

import "github.com/gin-gonic/gin"


var Router *gin.Engine

func SetupRouter() {
	Router = gin.New()
	Router.Use(gin.Logger())
	v1 := Router.Group("/api/v1")
	{
		v1.POST("/user/sign-up", SignUpHandler)
		// v1.POST("/user/sign-in", handlers.HandlerConfi)
		// auth := v1.Group("/auth", handlers.HandlerConfigInstance.Auth())
		// {
		// 	auth.POST("/user/sign-out", handlers.HandlerConfigInstance.SignOutHandler)
		// }
	}
}

func StartServer(){
	Router.Run(config.Port)
}