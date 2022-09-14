package main

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	deployapi "github.com/densmoe/mx-deploy/deploy_api"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	DeployAPIBaseURL  string `env:"DEPLOY_API_BASE_URL" envDefault:"https://deploy.mendix.com/api/1"`
	DeployAPIUsername string `env:"DEPLOY_API_USERNAME"`
	DeployAPIKey      string `env:"DEPLOY_API_KEY"`
	DebugMode         bool   `env:"DEBUG_MODE" envDefault:true`
}

var config Config

func main() {
	fmt.Printf("Hello")
	log.Info("dd")

	// Load .env in current path if available
	if err := godotenv.Load(); err != nil {
		log.Info("No .env file found")
	}

	if err := env.Parse(&config); err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", config)

	deployapi := deployapi.DeployAPI{
		BaseURL:  config.DeployAPIBaseURL,
		Username: config.DeployAPIUsername,
		APIKey:   config.DeployAPIKey,
	}
	apps := deployapi.RetrieveApps()
	log.Info(apps)
}
