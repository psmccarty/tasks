/*
Copyright © 2024 Patrick McCarty <patricksantos1234567@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <description>",
	Short: "Adds a task to the list",
	Long:  `Adds a task to the todo list application with a short description.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
