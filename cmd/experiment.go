package cmd

import (
	"encoding/json"

	"github.com/densmoe/mx-deploy/configuration"
	deployapiv4 "github.com/densmoe/mx-deploy/deploy_api_v4"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(experimentCmd)
	experimentCmd.AddCommand(experimentSetServiceAccountPermissionsCmd)
}

var experimentCmd = &cobra.Command{
	Use:   "experiment",
	Short: "experiment",
	Long:  `experiment`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var experimentSetServiceAccountPermissionsCmd = &cobra.Command{
	Use:   "setup-service-account",
	Short: "Set up service account",
	Long:  `Sets the correct permissions for a service account in all apps`,
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {

		userEmail := args[0]
		dUser := deployapiv4.DeployAPIv4{
			PAT: configuration.CurrentConfig.ExpPAT,
		}
		dAdmin := deployapiv4.DeployAPIv4{
			PAT: configuration.CurrentConfig.PAT,
		}

		apps := dUser.GetLicensedApps()
		for _, app := range apps {
			envs := dUser.GetEnvironments(app.ID)
			for _, environment := range envs {
				permissions := dAdmin.SetUserPermissionsForEnvironment(
					app.ID,
					environment.ID,
					userEmail,
					true,
					false,
					false,
					true,
					false,
					false,
				)
				out, _ := json.MarshalIndent(permissions, "", "  ")
				println(string(out))
			}
		}

	},
}
