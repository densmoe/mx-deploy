package cmd

import (
	"encoding/json"

	buildapi "github.com/densmoe/mx-deploy/build_api"
	"github.com/densmoe/mx-deploy/configuration"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(packagesCmd)
	packagesCmd.AddCommand(packagesLsCmd)
	packagesCmd.AddCommand(packagesLatestCmd)
}

var packagesCmd = &cobra.Command{
	Use:   "packages",
	Short: "packages",
	Long:  `packages`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var packagesLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Retrieves list of packages",
	Long:  `Retrieves a list of all packages that can be accessed with the credentials.`,
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		b := buildapi.BuildAPI{
			BaseURL:  configuration.CurrentConfig.DeployAPIBaseURL,
			Username: configuration.CurrentConfig.DeployAPIUsername,
			APIKey:   configuration.CurrentConfig.DeployAPIKey,
		}
		packages := b.RetrievePackages(args[0])
		out, _ := json.MarshalIndent(packages, "", "  ")
		println(string(out))
	},
}

var packagesLatestCmd = &cobra.Command{
	Use:   "latest",
	Short: "Retrieves latest package",
	Long:  `Retrieves the latest package for an app`,
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		b := buildapi.BuildAPI{
			BaseURL:  configuration.CurrentConfig.DeployAPIBaseURL,
			Username: configuration.CurrentConfig.DeployAPIUsername,
			APIKey:   configuration.CurrentConfig.DeployAPIKey,
		}
		packages := b.GetLatestPackage(args[0])
		out, _ := json.MarshalIndent(packages, "", "  ")
		println(string(out))
	},
}
