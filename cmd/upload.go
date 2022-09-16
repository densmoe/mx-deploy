package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(uploadCmd)
}

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "upload",
	Long:  `upload`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
