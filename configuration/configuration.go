package configuration

type Configuration struct {
	DeployAPIBaseURL  string `env:"DEPLOY_API_BASE_URL" envDefault:"https://deploy.mendix.com/api/1"`
	DeployAPIUsername string `env:"DEPLOY_API_USERNAME"`
	DeployAPIKey      string `env:"DEPLOY_API_KEY"`
	AppId             string `env:"DEPLOY_API_APP_ID"`
	ProjectId         string `env:"DEPLOY_API_PROJECT_ID"`
	DebugMode         bool   `env:"DEBUG_MODE" envDefault:true`
}

var CurrentConfig Configuration
