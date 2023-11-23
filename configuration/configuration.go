package configuration

type Configuration struct {
	DeployAPIUsername string `env:"DEPLOY_API_USERNAME"`
	DeployAPIKey      string `env:"DEPLOY_API_KEY"`
	PAT               string `env:"DEPLOY_API_PAT"`
	AppId             string `env:"DEPLOY_API_APP_ID"`
	ProjectId         string `env:"DEPLOY_API_PROJECT_ID"`
	DebugMode         bool   `env:"DEBUG_MODE" envDefault:"true"`
	ExpPAT            string `env:"EXP_PAT"`
}

var CurrentConfig Configuration
