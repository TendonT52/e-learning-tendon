package main

import (
	"testing"

	"github.com/TendonT52/e-learning-tendon/db"
)

func TestMain(m *testing.M) {
	loadConfig()
	setupInstance()
	setupRouter()
	ClearDatabase()
	m.Run()
	ClearDatabase()
}

func ClearDatabase() {
	db.UserDBInstance.CleanUp()
	db.JwtDBInstance.CleanUp()
}
