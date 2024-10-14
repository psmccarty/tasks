/*
Copyright © 2024 Patrick McCarty <patricksantos1234567@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete <taskid>",
	Short: "Deletes a task from the list",
	Long:  `Deletes a task given its taskid.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
