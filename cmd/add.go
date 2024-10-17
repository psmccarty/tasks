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

var dueDateString string

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

		var dueDate time.Time
		validDate := false
		if dueDateString != "" {
			duration, err := time.ParseDuration(dueDateString)
			if err != nil {
				fmt.Println(err)
				return
			}
			dueDate = time.Now().Add(duration)
			validDate = true
		}

		err = queries.CreateTask(ctx, gen.CreateTaskParams{
			Description:     args[0],
			CreateTimestamp: time.Now(),
			DueDateTimestamp: sql.NullTime{
				Time:  dueDate,
				Valid: validDate,
			},
		})
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&dueDateString, "due", "d", "", "Due date of task")
}
