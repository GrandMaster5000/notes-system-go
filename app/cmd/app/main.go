package main

import (
	"log"

	"github.com/GrandMaster5000/notes-system-go/internal/app"
	"github.com/GrandMaster5000/notes-system-go/internal/config"
	"github.com/GrandMaster5000/notes-system-go/pkg/logging"
)

func main() {
	log.Print("config initialized")
	cfg := config.GetConfig()

	log.Print("logger initializing")
	logging.Init(cfg.AppConfig.LogLevel)
	logger := logging.GetLogger()

	app, err := app.NewApp(cfg, logger)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Println("Running Application")
	app.Run()
}
