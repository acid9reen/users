package http

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"users/config"

	"users/internal/controller/http"
)

type App struct {
	config *config.Config
}

func New(config *config.Config) *App {
	return &App{config}
}

func (a *App) Run() {
	app := v1.Setup(a.config)
	err := app.Listen(":" + a.config.HTTP.Port)
	if err != nil {
		fmt.Printf("Can't start fiber server %v\n", err)
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		fmt.Println("app - Run - signal: " + s.String())
	}

	err = app.Shutdown()
	if err != nil {
		fmt.Println(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
