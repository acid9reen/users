package main

import (
	"fmt"

	"users/app/http"
	"users/config"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		fmt.Printf("Failed to load config: %s", err)
		return
	}

	app := http.New(config)
	app.Run()
}
