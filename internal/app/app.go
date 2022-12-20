package app

import (
	"time"

	"github.com/TendonT52/e-learning-tendon/internal/ports/repos"
)

type AppConfig struct {
	AppName              string
	AccessSecret         string
	RefreshSecret        string
	AccesstokenDuration  time.Duration
	RefreshtokenDuration time.Duration
}

type ReposInstance struct {
	UserDB   repos.UserDB
	JwtDB    repos.JwtDB
	CourseDB repos.CurriculumDB
	LessonDB repos.LessonDB
	NodeDB   repos.NodeDB
}

var appConfig AppConfig
var reposInstance ReposInstance


func NewApp(app AppConfig, repos ReposInstance) {
	appConfig = app
	reposInstance = repos
}
