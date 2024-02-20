/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

// getFileCmd represents the getFile command
var getFileCmd = &cobra.Command{
	Use:   "get-file",
	Short: "Get file by id",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fileId := args[0]

		host, err := cmd.Flags().GetString("host")

		if err != nil {
			log.Fatal(err)
		}

		Get(fileId, host)
	},
}

func init() {
	rootCmd.AddCommand(getFileCmd)
	getFileCmd.PersistentFlags().String("host", "", "Host to make request to")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getFileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getFileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
