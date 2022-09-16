package main

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/densmoe/mx-deploy/cmd"
	configuration "github.com/densmoe/mx-deploy/configuration"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	// log.SetFormatter(&log.JSONFormatter{PrettyPrint: true})
	// Load .env in current path if available
	if err := godotenv.Load(); err != nil {
		log.Info("No .env file found")
	}

	if err := env.Parse(&configuration.CurrentConfig); err != nil {
		fmt.Printf("%+v\n", err)
	}
	// fmt.Printf("%+v\n", config)
	cmd.Execute()
}
