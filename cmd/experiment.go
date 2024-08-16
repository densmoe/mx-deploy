package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/densmoe/mx-deploy/configuration"
	deployapi "github.com/densmoe/mx-deploy/deploy_api"
	deployapiv4 "github.com/densmoe/mx-deploy/deploy_api_v4"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(experimentCmd)
	experimentCmd.AddCommand(experimentSetServiceAccountPermissionsCmd)
	experimentCmd.AddCommand(experimentGetMxMessageConstantsCmd)
	experimentCmd.AddCommand(experimentUpdateConstantCmd)
	experimentCmd.AddCommand(experimentDetectDomainConstantsCmd)
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

var experimentGetMxMessageConstantsCmd = &cobra.Command{
	Use:   "get-mx-message-constants",
	Short: "Get Mendix message constants",
	Long:  `Get Mendix message constants`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		d1 := deployapi.DeployAPI{
			Username: configuration.CurrentConfig.DeployAPIUsername,
			APIKey:   configuration.CurrentConfig.DeployAPIKey,
		}
		// d4 := deployapiv4.DeployAPIv4{
		// 	PAT: configuration.CurrentConfig.PAT,
		// }
		println("Getting constants")
		apps := d1.RetrieveApps()
		for _, app := range apps {
			// println(app.Name, app.AppId)
			if app.Name != "Sprintr" {
				continue
			}
			envs := d1.RetrieveEnvironments(app.AppId)
			for _, environment := range envs {
				if environment.Mode != "Acceptance" {
					continue
				}
				// println(app.Name, environment.Mode, environment.Status)
				constants, _, _, err := d1.GetEnvironmentSettings(app.AppId, environment.Mode)
				// println(constants)
				if err != nil {
					println(err)
					return
				}
				if err != nil {
					println(err)
					return
				}
				for _, constant := range constants {
					if constant.Name == "MxMessage.MxMessageLocation" {
						println(app.Name, environment.Mode, environment.Status)
						fmt.Printf("MxMessage.MxMessageLocation: %s \n", constant.Value)
					}
				}
			}
		}

		// constants := d.GetMxMessageConstants()
		// out, _ := json.MarshalIndent(constants, "", "  ")
		// println(string(out))
	},
}

var experimentDetectDomainConstantsCmd = &cobra.Command{
	Use:   "detect-domain-constants",
	Short: "Detect domain constants",
	Long:  `Detect domain constants`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		d1 := deployapi.DeployAPI{
			Username: configuration.CurrentConfig.DeployAPIUsername,
			APIKey:   configuration.CurrentConfig.DeployAPIKey,
		}
		domainValue := args[0]
		fmt.Printf("Searching for domain %s in constants... \n\n", domainValue)
		apps := d1.RetrieveApps()
		for _, app := range apps {
			// println(app.Name, app.AppId)
			// if app.Name != "Sprintr" {
			// 	continue
			// }
			envs := d1.RetrieveEnvironments(app.AppId)
			for _, environment := range envs {
				// if environment.Mode != "Acceptance" {
				// 	continue
				// }
				// println(app.Name, environment.Mode, environment.Status)
				constants, _, _, err := d1.GetEnvironmentSettings(app.AppId, environment.Mode)
				// println(constants)
				if err != nil {
					println(err)
					return
				}
				if err != nil {
					println(err)
					return
				}
				for _, constant := range constants {
					if strings.Contains(constant.Value, domainValue) {
						fmt.Printf("Detected constant in app: %s environment: %s \n", app.Name, environment.Mode)
						fmt.Printf("Name: %s \n", constant.Name)
						fmt.Printf("Value: %s \nDeployedValue: %s\n\n", constant.Value, constant.DeployedValue)
					}
				}
			}
		}

		// constants := d.GetMxMessageConstants()
		// out, _ := json.MarshalIndent(constants, "", "  ")
		// println(string(out))
	},
}

var experimentUpdateConstantCmd = &cobra.Command{
	Use:   "update-constant",
	Short: "Update environment constants",
	Long:  `Update environment constants`,
	Args:  cobra.MatchAll(cobra.ExactArgs(4)),
	Run: func(cmd *cobra.Command, args []string) {
		d := deployapi.DeployAPI{
			Username: configuration.CurrentConfig.DeployAPIUsername,
			APIKey:   configuration.CurrentConfig.DeployAPIKey,
		}
		subdomain := args[0]
		environmentMode := args[1]
		constantName := args[2]
		constantValue := args[3]

		// println(app.Name, environment.Mode, environment.Status)
		constants, scheduledEvents, customSettings, err := d.GetEnvironmentSettings(subdomain, environmentMode)
		// println(constants)
		if err != nil {
			println(err)
			return
		}
		if err != nil {
			println(err)
			return
		}
		println(subdomain, environmentMode)
		newConstants := []deployapi.Constant{}
		for _, constant := range constants {
			if constant.Name == constantName {
				fmt.Printf("Constant found: %s \n", constant.Name)
				fmt.Printf("Current value: %s \n", constant.Value)
				fmt.Printf("New Value: %s \n", constantValue)
				constant.Value = constantValue
			}
			newConstants = append(newConstants, constant)
		}
		d.SetEnvironmentSettings(subdomain, environmentMode, newConstants, scheduledEvents, customSettings)

		// out, _ := json.MarshalIndent(constants, "", "  ")
		// println(string(out))
	},
}

// https://mxmessage-accp.mendixcloud.com/rest/v1
// https://mxmessage-api-accp.mendix.com/rest/v1
