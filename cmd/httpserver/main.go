package main

import (
	"fmt"

	"github.com/acid9reen/users/app/http"
	"github.com/acid9reen/users/config"
	"github.com/acid9reen/users/pkg/logging"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		fmt.Printf("Failed to load config: %s", err)
		return
	}

	logger := logging.NewLogger(config.Logger.Level)

	logger.Info("Initializing app")
	app := http.New(config, logger)
	logger.Info("App created")
	app.Run()
}
