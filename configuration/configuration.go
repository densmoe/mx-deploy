package configuration

type Configuration struct {
	DeployAPIBaseURL  string `env:"DEPLOY_API_BASE_URL" envDefault:"https://deploy.mendix.com/api/1"`
	DeployAPIUsername string `env:"DEPLOY_API_USERNAME"`
	DeployAPIKey      string `env:"DEPLOY_API_KEY"`
	DebugMode         bool   `env:"DEBUG_MODE" envDefault:true`
}

var CurrentConfig Configuration
