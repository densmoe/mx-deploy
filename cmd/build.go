package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(buildCmd)
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "build",
	Long:  `build`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
