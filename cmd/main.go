package main

import (
	"fmt"
	"log"

	"github.com/TendonT52/e-learning-tendon/db"
	"github.com/TendonT52/e-learning-tendon/handlers"
	"github.com/TendonT52/e-learning-tendon/internal/application"
	"github.com/TendonT52/e-learning-tendon/internal/core"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func loadConfig() {
	log.Println("Loading config...")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("..")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	log.Println("Load config success")
}

func setupRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	v1 := router.Group("/api/v1")
	{
		v1.POST("/user/sign-up", handlers.SignUpHandler)
	}
	return router
}

func main() {
	loadConfig()
	db.NewClient(viper.GetString("mongo.connection"), db.MongoConfig{
		CreateTimeOut: viper.GetDuration("mongo.insertTimeOut"),
		FindTimeout:   viper.GetDuration("mongo.findTimeOut"),
		UpdateTimeout: viper.GetDuration("mongo.updateTimeOut"),
		DeleteTimeout: viper.GetDuration("mongo.deleteTimeOut"),
	})
	db.NewDB(viper.GetString("mongo.name"))
	db.NewUserDB(viper.GetString("mongo.collection.user.name"))
	db.NewJwtTokenDB(viper.GetString("mongo.collection.jwt.name"))
	application.NewUserService(db.UserDBInstance)
	application.NewJwtService(db.JwtDBInstance, core.JwtConfig{
		AppName:              viper.GetString("token.issuer"),
		AccessSecret:         viper.GetString("token.jwtAccessSecret"),
		RefreshSecret:        viper.GetString("token.jwtRefreshSecret"),
		AccesstokenDuration:  viper.GetDuration("token.accessTokenExpire"),
		RefreshtokenDuration: viper.GetDuration("token.refreshTokenExpire"),
	})
	router := setupRouter()
	router.Run(":8080")
}
