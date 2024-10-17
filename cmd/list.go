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

	"github.com/fatih/color"
	"github.com/mergestat/timediff"
	"github.com/psmccarty/tasks/sql/gen"
	"github.com/spf13/cobra"
)

const (
	FullHeader   = "ID\tTask\tCreated\tDue\tCompleted\n"
	FullTemplate = "%d\t%s\t%s\t%s\t%v\n"

	ReducedHeader   = "ID\tTask\tCreated\tDue\n"
	ReducedTemplate = "%d\t%s\t%s\t%s\n"
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
		c := color.New(color.FgWhite)
		if all {
			c.Fprintf(w, FullHeader)
		} else {
			c.Fprintf(w, ReducedHeader)
		}
		for _, t := range tasks {

			var dueDateString string
			c := color.New(color.FgWhite)
			if t.DueDateTimestamp.Valid {
				dueDateString = timediff.TimeDiff(t.DueDateTimestamp.Time)
				if time.Since(t.DueDateTimestamp.Time).Hours() > -1 && time.Since(t.DueDateTimestamp.Time).Hours() < 0 {
					c = color.New(color.FgYellow)
				}
			} else {
				dueDateString = "-"
			}

			var completedOnString string
			if t.CompletedTimestamp.Valid {
				completedOnString = timediff.TimeDiff(t.CompletedTimestamp.Time)
			} else {
				completedOnString = "-"
			}

			if t.DueDateTimestamp.Valid && t.CompletedTimestamp.Valid {
				completedSince := time.Since(t.DueDateTimestamp.Time)
				dueSince := time.Since(t.CompletedTimestamp.Time)
				lateSince := completedSince - dueSince

				rounding := time.Second

				if lateSince.Abs().Hours() >= 1 {
					rounding = time.Hour
				} else if lateSince.Abs().Minutes() >= 1 {
					rounding = time.Minute
				}

				if lateSince < 0 {
					completedOnString = fmt.Sprintf("%s early", lateSince.Abs().Round(rounding))
					c = color.New(color.FgGreen)
				} else {
					completedOnString = fmt.Sprintf("%s late", lateSince.Round(rounding))
					c = color.New(color.FgRed)
				}
			}

			if all {
				c.Fprintf(w, FullTemplate, t.ID, t.Description, timediff.TimeDiff(t.CreateTimestamp), dueDateString, completedOnString)
			} else {
				c.Fprintf(w, ReducedTemplate, t.ID, t.Description, timediff.TimeDiff(t.CreateTimestamp), dueDateString)
			}
		}
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&all, "all", "a", false, "include completed tasks")
}
