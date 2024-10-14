/*
Copyright © 2024 Patrick McCarty <patricksantos1234567@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete <taskid>",
	Short: "Marks a task as complete",
	Long:  `Mark a task complete given its taskid.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("complete called")
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
