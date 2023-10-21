/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"rss-cli/data"
	"github.com/spf13/cobra"
)

var addFeedCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a RSS feed to the known list of feeds",
	Long: `Add a RSS feed to the known list of feeds `,
	Run: func(cmd *cobra.Command, args []string) {
    url, err := cmd.Flags().GetString("url")
    if err != nil {
      return 
    }

    database, err := data.ConnectToDatabase()
    if err != nil {
      return
    }

    results := database.Model(data.Feed {Url: url}).First(&data.Feed{})
    if results.RowsAffected > 0 {
      return 
    }

    feed, err := data.FeedFromUrl(url)
    if err != nil {
      _ = fmt.Errorf("%v", err)
      return 
    }
    database.Create(&feed)
  },
}

func init() {
	feedCmd.AddCommand(addFeedCmd)
	feedCmd.AddCommand(listFeedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addFeedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	addFeedCmd.Flags().StringP("url", "u", "", "Url of the feed")
  addFeedCmd.MarkFlagRequired("url")
}
