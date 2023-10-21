package cmd

import "github.com/spf13/cobra"

var listFeedCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a RSS feed to the known list of feeds",
	Long:  `Add a RSS feed to the known list of feeds `,
	Run: func(cmd *cobra.Command, args []string) {
		
	},
}
