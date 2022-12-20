package main

import (
	"github.com/TendonT52/e-learning-tendon/config"
	"github.com/TendonT52/e-learning-tendon/handlers"
)

func main() {
	config.LoadConfig()
	config.SetupInstance()
	handlers.SetupRouter()
}
