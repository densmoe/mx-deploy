package main

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	buildapi "github.com/densmoe/mx-deploy/build_api"
	"github.com/densmoe/mx-deploy/cmd"
	configuration "github.com/densmoe/mx-deploy/configuration"
	deployapi "github.com/densmoe/mx-deploy/deploy_api"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	DeployAPIBaseURL  string `env:"DEPLOY_API_BASE_URL" envDefault:"https://deploy.mendix.com/api/1"`
	DeployAPIUsername string `env:"DEPLOY_API_USERNAME"`
	DeployAPIKey      string `env:"DEPLOY_API_KEY"`
	AppId             string `env:"DEPLOY_API_APP_ID"`
	ProjectId         string `env:"DEPLOY_API_PROJECT_ID"`
	DebugMode         bool   `env:"DEBUG_MODE" envDefault:true`
}

var config Config
var Client deployapi.DeployAPI

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

	deployapi := deployapi.DeployAPI{
		BaseURL:   config.DeployAPIBaseURL,
		Username:  config.DeployAPIUsername,
		APIKey:    config.DeployAPIKey,
		AppID:     config.AppId,
		ProjectID: config.ProjectId,
	}

	apps := deployapi.RetrieveApps()
	log.Info(apps)
	app := deployapi.RetrieveApp()
	log.Info(app)
	log.Info(deployapi.AppID)
	appId := deployapi.GetAppIdForProjectId(config.ProjectId)
	log.Info(appId)
	log.Info(deployapi.AppID)
	deployapi.SetAppIdForProjectId(config.ProjectId)
	log.Info(deployapi.AppID)

	buildapi := buildapi.BuildAPI{
		BaseURL:  config.DeployAPIBaseURL,
		Username: config.DeployAPIUsername,
		APIKey:   config.DeployAPIKey,
	}

	buildapi.AppID = deployapi.GetAppIdForProjectId(config.ProjectId)
	packages := buildapi.RetrievePackages()
	log.Info(packages)

	lastPackage := buildapi.GetLatestPackage()
	log.Info(lastPackage)

	rev := buildapi.GetRevisionFromPackage(lastPackage)
	log.Info(rev)
	cmd.Execute()
}
