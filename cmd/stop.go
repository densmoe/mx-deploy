package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(stopCmd)
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop",
	Long:  `stop`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
