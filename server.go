package main

import (
	"fmt"

	"github.com/TendonT52/e-learning-tendon/config"
	"github.com/TendonT52/e-learning-tendon/handlers"
)

func main() {
	fmt.Println("server is running")
	config.LoadConfig()
	config.SetupInstance()
	handlers.InitUser()
	handlers.SetupRouter()
	handlers.StartServer()
}
