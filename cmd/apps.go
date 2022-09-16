package cmd

import (
	"encoding/json"

	"github.com/densmoe/mx-deploy/configuration"
	deployapi "github.com/densmoe/mx-deploy/deploy_api"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(appsCmd)
	appsCmd.AddCommand(appsLsCmd)
	appsCmd.AddCommand(appsInfoCmd)
}

var appsCmd = &cobra.Command{
	Use:   "apps",
	Short: "apps",
	Long:  `apps`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var appsInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Retrieves information about an app",
	Long:  `app`,
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		d := deployapi.DeployAPI{
			BaseURL:  configuration.CurrentConfig.DeployAPIBaseURL,
			Username: configuration.CurrentConfig.DeployAPIUsername,
			APIKey:   configuration.CurrentConfig.DeployAPIKey,
		}
		apps := d.RetrieveApp(args[0])
		out, _ := json.MarshalIndent(apps, "", "  ")
		println(string(out))
	},
}

var appsLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Retrieves list of apps",
	Long:  `Retrieves a list of all apps that can be accessed with the credentials.`,
	Run: func(cmd *cobra.Command, args []string) {
		d := deployapi.DeployAPI{
			BaseURL:  configuration.CurrentConfig.DeployAPIBaseURL,
			Username: configuration.CurrentConfig.DeployAPIUsername,
			APIKey:   configuration.CurrentConfig.DeployAPIKey,
		}
		apps := d.RetrieveApps()
		out, _ := json.MarshalIndent(apps, "", "  ")
		println(string(out))
	},
}
