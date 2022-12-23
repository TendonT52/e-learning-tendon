package config

import (
	"fmt"
	"log"

	"github.com/TendonT52/e-learning-tendon/db"
	"github.com/TendonT52/e-learning-tendon/handlers"
	"github.com/TendonT52/e-learning-tendon/internal/app"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func LoadConfig() {
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
	gin.SetMode(gin.ReleaseMode)
}

func LoadConfigTest() {
	log.Println("Loading config...")
	viper.SetConfigName("config_test")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("..")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	log.Println("Load config success")
}

func SetupInstance() {

	db.NewClient(viper.GetString("mongo.connection"), db.MongoConfig{
		InsertTimeOut: viper.GetDuration("mongo.insertTimeOut"),
		FindTimeOut:   viper.GetDuration("mongo.findTimeOut"),
		UpdateTimeOut: viper.GetDuration("mongo.updateTimeOut"),
		DeleteTimeOut: viper.GetDuration("mongo.deleteTimeOut"),
	})
	db.NewDB(viper.GetString("mongo.name"))
	db.NewUserDB(viper.GetString("mongo.collection.user.name"))
	db.NewJwtDB(viper.GetString("mongo.collection.jwt.name"))
	db.NewCourseDB(viper.GetString("mongo.collection.course.name"))
	db.NewLessonDB(viper.GetString("mongo.collection.lesson.name"))
	db.NewNodeDB(viper.GetString("mongo.collection.node.name"))

	appConfig := app.AppConfig{
		AppName:              viper.GetString("token.issuer"),
		AccessSecret:         viper.GetString("token.access.secret"),
		RefreshSecret:        viper.GetString("token.refresh.secret"),
		AccesstokenDuration:  viper.GetDuration("token.access.expire"),
		RefreshtokenDuration: viper.GetDuration("token.refresh.expire"),
	}

	reposConfig := app.ReposInstance{
		UserDB:   db.UserDBInstance,
		JwtDB:    db.JwtDBInstance,
		CourseDB: db.CourseDBInstance,
		LessonDB: db.LessonDBInstance,
		NodeDB:   db.NodeDBInstance,
	}

	app.NewApp(appConfig, reposConfig)

	handlers.SetConfig(
		handlers.Config{
			Port:                  viper.GetString("app.port"),
			Url:                   viper.GetString("app.url"),
			AccessCookieSecure:    viper.GetBool("token.access.cookie.secure"),
			AccessCookieHttpOnly:  viper.GetBool("token.access.cookie.httpOnly"),
			AccessTokenDuration:   viper.GetDuration("token.access.expire"),
			RefreshCookieSecure:   viper.GetBool("token.refresh.cookie.secure"),
			RefreshCookieHttpOnly: viper.GetBool("token.refresh.cookie.httpOnly"),
			RefreshTokenDuration:  viper.GetDuration("token.refresh.expire"),
		},
	)

	
}
