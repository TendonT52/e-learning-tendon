package handlers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func SetupRouter() {
	Router = gin.New()
	Router.Use(cors.Default())
	Router.Use(gin.Logger())
	v1 := Router.Group("/api/v1")
	{
		v1.POST("/user/sign-up", SignUpHandler)
		v1.POST("/user/sign-in", SignInHandler)
		v1.POST("/user/sign-out", SignOutHandler)
		v1.POST("/user/refresh-token", RefreshTokenHandler)

		auth := v1.Group("/auth", HaveAccessToken)
		{
			authAdmin := auth.Group("", IsAdmin)
			{
				authAdmin.GET("/users/:id", GetUserHandler)
				authAdmin.PATCH("/users/:id", PatchUserHandler)
				authAdmin.DELETE("/users/:id", DeleteUser)
			}
			authTeacher := auth.Group("", IsTeacher)
			{
				authTeacher.POST("/courses", PostCourseHandler)

				authTeacher.POST("/lessons", PostLessonHandler)

				authTeacher.POST("/nodes", PostNodeHandler)
			}
			authStudent := auth.Group("", IsStudent)
			{
				authStudent.GET("/courses/:id", GetCourseHandler)
				authStudent.PATCH("/courses/:id", PatchCourseHandler)
				authStudent.DELETE("/courses/:id", DeleteCourseHandler)

				authStudent.GET("/lessons/:id", GetLessonHandler)
				authStudent.PATCH("/lessons/:id", PatchLessonHandler)
				authStudent.DELETE("/lessons/:id", DeleteLessonHandler)

				authStudent.GET("/nodes/:id", GetNodeHandler)
				authStudent.PATCH("/nodes/:id", PatchNodeHandler)
				authStudent.DELETE("/nodes/:id", DeleteNodeHandler)
			}
		}

	}
}

func StartServer() {
	Router.Run(config.Port)
}
