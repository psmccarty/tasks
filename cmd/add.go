/*
Copyright © 2024 Patrick McCarty <patricksantos1234567@gmail.com>
*/
package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/psmccarty/tasks/sql/gen"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <description>",
	Short: "Adds a task to the list",
	Long:  `Adds a task to the todo list application with a short description.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		db, err := sql.Open("sqlite3", "tasks.db")
		if err != nil {
			os.Exit(1)
		}
		defer db.Close()

		ctx := context.Background()
		queries := gen.New(db)
		task, err := queries.CreateTask(ctx, gen.CreateTaskParams{
			Description:     args[0],
			CreateTimestamp: time.Now(),
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
