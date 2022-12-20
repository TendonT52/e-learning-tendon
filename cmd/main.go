package main

import (
	"fmt"
	"log"

	"github.com/TendonT52/e-learning-tendon/db"
	"github.com/TendonT52/e-learning-tendon/handlers"
	"github.com/TendonT52/e-learning-tendon/internal/app"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var router *gin.Engine

func loadConfig() {
	log.Println("Loading config...")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("..")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	log.Println("Load config success")
}

func setupInstance() {

	db.NewClient(viper.GetString("mongo.connection"), db.MongoConfig{
		InsertTimeOut: viper.GetDuration("mongo.insertTimeOut"),
		FindTimeOut:   viper.GetDuration("mongo.findTimeOut"),
		UpdateTimeOut: viper.GetDuration("mongo.updateTimeOut"),
		DeleteTimeOut: viper.GetDuration("mongo.deleteTimeOut"),
	})
	db.NewDB(viper.GetString("mongo.name"))
	db.NewUserDB(viper.GetString("mongo.collection.user.name"))
	db.NewJwtDB(viper.GetString("mongo.collection.jwt.name"))
	db.NewCourseDB(viper.GetString("mongo.collection.curriculum.name"))
	db.NewLessonDB(viper.GetString("mongo.collection.lesson.name"))
	db.NewNodeDB(viper.GetString("mongo.collection.node.name"))

	appConfig := app.AppConfig{
		AppName:              viper.GetString("token.issuer"),
		AccessSecret:         viper.GetString("token.jwtAccessSecret"),
		RefreshSecret:        viper.GetString("token.jwtRefreshSecret"),
		AccesstokenDuration:  viper.GetDuration("token.accessTokenExpire"),
		RefreshtokenDuration: viper.GetDuration("token.refreshTokenExpire"),
	}

	reposConfig := app.ReposInstance{
		UserDB:   db.UserDBInstance,
		JwtDB:    db.JwtDBInstance,
		CourseDB: db.CourseDBInstance,
		LessonDB: db.LessonDBInstance,
		NodeDB:   db.NodeDBInstance,
	}

	app.NewApp(appConfig, reposConfig)
}

func setupRouter() {
	router = gin.New()
	router.Use(gin.Logger())
	handlers.NewHandlerConfig(
		application.UserServiceInstance,
		application.JwtServiceInstance,
		handlers.Config{
			Url:            viper.GetString("app.url"),
			CookieSecure:   viper.GetBool("cookie.secure"),
			CookieHttpOnly: viper.GetBool("cookire.httpOnly"),
		})
	v1 := router.Group("/api/v1")
	{
		v1.POST("/user/sign-up", handlers.HandlerConfigInstance.SignUpHandler)
		v1.POST("/user/sign-in", handlers.HandlerConfigInstance.SignInHandler)
		auth := v1.Group("/auth", handlers.HandlerConfigInstance.Auth())
		{
			auth.POST("/user/sign-out", handlers.HandlerConfigInstance.SignOutHandler)
		}
	}
}

func main() {
	loadConfig()
	setupInstance()
	setupRouter()
	router.Run(":8080")
}
