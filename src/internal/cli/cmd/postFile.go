/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

// postFileCmd represents the postFile command
var postFileCmd = &cobra.Command{
	Use:   "post-file",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		host, err := cmd.Flags().GetString("host")

		if err != nil {
			log.Fatal(err)
		}

		filePath, err := cmd.Flags().GetString("path")
		if err != nil {
			log.Fatal(err)
		}

		dirName, err := cmd.Flags().GetString("dir")
		if err != nil {
			log.Fatal(err)
		}

		AddFile(filePath, host, dirName)
	},
}

func init() {
	rootCmd.AddCommand(postFileCmd)

	postFileCmd.Flags().String("path", "", "FullPath of a file to send")
	postFileCmd.Flags().String("dir", "", "File will be saved in dir directory")
	postFileCmd.MarkFlagRequired("path")
	postFileCmd.MarkFlagRequired("dir")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// postFileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// postFileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
