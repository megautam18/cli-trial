/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// currCmd represents the curr command
var currCmd = &cobra.Command{
	Use:   "curr",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Retrieving available tasks..")
		type Task struct {
			Description string
		}
		var tasks []Task
		data, err := os.ReadFile("todo.json")
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		err = json.Unmarshal(data, &tasks)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			return
		}

		// Display tasks
		if len(tasks) == 0 {
			fmt.Println("No tasks available")
		} else {
			for i, task := range tasks {
				fmt.Printf("%d. %s\n", i+1, task.Description)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(currCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// currCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// currCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
