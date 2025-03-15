/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Supposed to add new tasks to the to-do list.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hey! So you want to create a new To-Do List. Let's begin.")
		type Task struct {
			Description string `json:"description"`
		}
		var tasks []Task
		r := bufio.NewReader(os.Stdin)
		fmt.Println("How many tasks do you wish to enter? ")
		var no int
		fmt.Scanln(&no)
		r.ReadString('\n')
		for i := 0; i < no; i++ {
			fmt.Printf("Enter task %d: ", i+1)
			str, err := r.ReadString('\n')
			if err != nil {
				fmt.Printf("Error: %v", err)
				return
			}
			tasks = append(tasks, Task{Description: str})

		}

		// Marshal tasks to JSON
		jsonData, err := json.MarshalIndent(tasks, "", "  ")
		if err != nil {
			fmt.Printf("Error marshaling to JSON: %v", err)
			return
		}

		// Print or save the JSON data
		os.WriteFile("todo.json", jsonData, 0644)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
