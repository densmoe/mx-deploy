package cmd

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/briandowns/spinner"

	buildapi "github.com/densmoe/mx-deploy/build_api"
	"github.com/densmoe/mx-deploy/configuration"
	deployapiv2 "github.com/densmoe/mx-deploy/deploy_api_v2"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(packagesCmd)
	packagesCmd.AddCommand(packagesLsCmd)
	packagesCmd.AddCommand(packagesLatestCmd)
	rootCmd.AddCommand(uploadCmd)
	packagesCmd.AddCommand(uploadCmd)
	uploadCmd.PersistentFlags().BoolP("autobuild", "a", false, "toggle autobuild")
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
			Username: configuration.CurrentConfig.DeployAPIUsername,
			APIKey:   configuration.CurrentConfig.DeployAPIKey,
		}
		packages := b.GetLatestPackage(args[0])
		out, _ := json.MarshalIndent(packages, "", "  ")
		println(string(out))
	},
}

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "upload",
	Long:  `upload`,
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		appId := args[0]
		if cmd.Flags().Changed("autobuild") {
			fmt.Println("autobuild enabled, but not implemented yet")
		}
		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		s.Prefix = "Uploading package... "
		s.Start()
		d := deployapiv2.DeployAPIv2{
			Username: configuration.CurrentConfig.DeployAPIUsername,
			APIKey:   configuration.CurrentConfig.DeployAPIKey,
		}
		uploadResponse := d.UploadPackage(appId, args[1])
		// out, _ := json.MarshalIndent(uploadResponse, "", "  ")
		// println(string(out))
		s.Stop()
		fmt.Println("Uploading package... done.")
		fmt.Printf("JobId: %s", uploadResponse.JobId)
		s.Prefix = "Processing uploaded package... "
		s.Start()
		statusResponse := d.PollForUploadStatus(appId, uploadResponse.JobId, 300*time.Second)
		s.Stop()
		fmt.Println("Processing uploaded package... done.")
		fmt.Printf("Upload complete with status: %s", statusResponse.Status)
	},
}
