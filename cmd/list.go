/*
Copyright © 2024 Patrick McCarty <patricksantos1234567@gmail.com>
*/
package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/psmccarty/tasks/sql/gen"
	"github.com/spf13/cobra"
)

var (
	all bool
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all current tasks",
	Long:  `Lists all current tasks in the todo list.`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := sql.Open("sqlite3", "tasks.db")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer db.Close()

		ctx := context.Background()
		queries := gen.New(db)

		var tasks []gen.Task
		if all {
			tasks, err = queries.ListAllTasks(ctx)
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			tasks, err = queries.ListUncompletedTasks(ctx)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
		fmt.Fprintln(w, "ID\tTask\tCreated")
		for _, t := range tasks {
			fmt.Fprintf(w, "%d\t%s\t%s\n", t.ID, t.Description, t.CreateTimestamp.Format(time.RFC3339))
		}
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&all, "all", "a", false, "include completed tasks")
}
