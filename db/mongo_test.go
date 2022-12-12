package db_test

import (
	"log"
	"testing"
	"time"

	"github.com/TendonT52/e-learning-tendon/db"
)

func TestMain(m *testing.M) {
	db.NewClient("mongodb://admin:password@localhost:27017",
		db.MongoConfig{
			CreateTimeOut: time.Minute,
			FindTimeout:   time.Minute,
			UpdateTimeout: time.Minute,
			DeleteTimeout: time.Minute,
		})
	defer db.DisconnectMongo()
	db.NewDB("tendon")
	db.NewUserDB("user_test")
	db.NewJwtTokenDB("jwt_test")
	m.Run()
	log.Println(db.UserDBInstance.CleanUp())
	log.Println(db.JwtDBInstance.CleanUp())
}
