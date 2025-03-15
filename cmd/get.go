/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("To get instructions on how to proceed: \n type 1 for new to-do list")
		fmt.Println("type 2 for updating your current to-do list")
		fmt.Println("type 3 to display your current to-do lists.")
		var choice int32
		fmt.Scanf("%v", &choice)
		switch choice {
		case 1:
			fmt.Println("Enter the statemnent 'go run main.go new' in the terminal.")
		case 2:
			fmt.Println("Enter the statement 'go run main.go update' in the terminal.")
		case 3:
			fmt.Println("Enter the statement 'go run main.go curr' in the terminal.")
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
