package cmd

import (
	"encoding/json"

	"github.com/densmoe/mx-deploy/configuration"
	deployapi "github.com/densmoe/mx-deploy/deploy_api"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(environmentsCmd)
	environmentsCmd.AddCommand(environmentsLsCmd)
}

var environmentsCmd = &cobra.Command{
	Use:   "environments",
	Short: "environments",
	Long:  `environments`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var environmentsLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List environments for app",
	Long:  `Lists the environments of a app`,
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		d := deployapi.DeployAPI{
			BaseURL:  configuration.CurrentConfig.DeployAPIBaseURL,
			Username: configuration.CurrentConfig.DeployAPIUsername,
			APIKey:   configuration.CurrentConfig.DeployAPIKey,
		}
		environments := d.RetrieveEnvironments(args[0])
		out, _ := json.MarshalIndent(environments, "", "  ")
		println(string(out))
	},
}
