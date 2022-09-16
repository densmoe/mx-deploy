package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of mx-deploy",
	Long:  `All software has versions. This is mx-deploy's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mx-deploy v0.1")
	},
}
