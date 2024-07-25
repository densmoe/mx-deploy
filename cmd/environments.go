package cmd

import (
	"encoding/json"

	"github.com/densmoe/mx-deploy/configuration"
	deployapi "github.com/densmoe/mx-deploy/deploy_api"
	deployapiv4 "github.com/densmoe/mx-deploy/deploy_api_v4"
	"github.com/spf13/cobra"
)

var canDeploy bool
var canManageBackups bool
var canViewAlerts bool
var canAccessAPI bool
var canViewLogs bool
var canManagePrivileges bool

func init() {
	rootCmd.AddCommand(environmentsCmd)
	environmentsCmd.AddCommand(environmentsLsCmd)
	environmentsCmd.AddCommand(environmentsInfoCmd)
	environmentsCmd.AddCommand(environmentsGetUserPermissionsCmd)
	environmentsCmd.AddCommand(environmentsSetUserPermissionsCmd)
	environmentsSetUserPermissionsCmd.Flags().BoolVar(&canDeploy, "canDeploy", false, "Set the canDeploy permission")
	environmentsSetUserPermissionsCmd.Flags().BoolVar(&canManageBackups, "canManageBackups", false, "Set the canManageBackups permission")
	environmentsSetUserPermissionsCmd.Flags().BoolVar(&canViewAlerts, "canViewAlerts", false, "Set the canViewAlerts permission")
	environmentsSetUserPermissionsCmd.Flags().BoolVar(&canAccessAPI, "canAccessAPI", false, "Set the canAccessAPI permission")
	environmentsSetUserPermissionsCmd.Flags().BoolVar(&canViewLogs, "canViewLogs", false, "Set the canViewLogs permission")
	environmentsSetUserPermissionsCmd.Flags().BoolVar(&canManagePrivileges, "canManagePrivileges", false, "Set the canManagePrivileges permission")
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
			Username: configuration.CurrentConfig.DeployAPIUsername,
			APIKey:   configuration.CurrentConfig.DeployAPIKey,
		}
		environments := d.RetrieveEnvironments(args[0])
		out, _ := json.MarshalIndent(environments, "", "  ")
		println(string(out))
	},
}

var environmentsInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Retrieves information about an environment",
	Long:  `Retrieves information about an environment`,
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		d := deployapi.DeployAPI{
			Username: configuration.CurrentConfig.DeployAPIUsername,
			APIKey:   configuration.CurrentConfig.DeployAPIKey,
		}
		environment := d.RetrieveEnvironment(args[0], args[1])
		out, _ := json.MarshalIndent(environment, "", "  ")
		println(string(out))
	},
}

var environmentsGetConstantsCmd = &cobra.Command{
	Use:   "get-constants",
	Short: "Retrieves constants of an environment",
	Long:  `Retrieves constants of an environment`,
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		d := deployapi.DeployAPI{
			Username: configuration.CurrentConfig.DeployAPIUsername,
			APIKey:   configuration.CurrentConfig.DeployAPIKey,
		}
		constants, _, _, _ := d.GetEnvironmentSettings(args[0], args[1])
		out, _ := json.MarshalIndent(constants, "", "  ")
		println(string(out))
	},
}

var environmentsGetUserPermissionsCmd = &cobra.Command{
	Use:   "get-permissions",
	Short: "Retrieves permisions of a sincle user in an environment",
	Long:  `Retrieves permisions of a sincle user in an environment`,
	Args:  cobra.MatchAll(cobra.ExactArgs(3)),
	Run: func(cmd *cobra.Command, args []string) {
		d := deployapiv4.DeployAPIv4{
			PAT: configuration.CurrentConfig.PAT,
		}
		permissions := d.GetUserPermissionsForEnvironment(args[0], args[1], args[2])
		out, _ := json.MarshalIndent(permissions, "", "  ")
		println(string(out))
	},
}

var environmentsSetUserPermissionsCmd = &cobra.Command{
	Use:   "set-permissions",
	Short: "Sets permisions of a sincle user in an environment",
	Long:  `Sets permisions of a sincle user in an environment`,
	Args:  cobra.MatchAll(cobra.ExactArgs(3)),
	Run: func(cmd *cobra.Command, args []string) {
		d := deployapiv4.DeployAPIv4{
			PAT: configuration.CurrentConfig.PAT,
		}
		permissions := d.SetUserPermissionsForEnvironment(
			args[0],
			args[1],
			args[2],
			canDeploy,
			canManageBackups,
			canViewAlerts,
			canAccessAPI,
			canViewLogs,
			canManagePrivileges,
		)
		out, _ := json.MarshalIndent(permissions, "", "  ")
		println(string(out))
	},
}
