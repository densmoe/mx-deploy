package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(transportCmd)
}

var transportCmd = &cobra.Command{
	Use:   "transport",
	Short: "transport",
	Long:  `transport`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
