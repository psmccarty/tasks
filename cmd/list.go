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
	listAll bool
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
		if listAll {
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
		if listAll {
			fmt.Fprint(w, FullHeader)
		} else {
			fmt.Fprint(w, ReducedHeader)
		}
		for _, t := range tasks {
			dueOn, completedOn := dueAndCompletedStrings(t)
			if listAll {
				fmt.Fprintf(w, FullTemplate, t.ID, t.Description, timediff.TimeDiff(t.CreateTimestamp), dueOn, completedOn)
			} else {
				fmt.Fprintf(w, ReducedTemplate, t.ID, t.Description, timediff.TimeDiff(t.CreateTimestamp), dueOn)
			}
		}
		w.Flush()
	},
}

func dueAndCompletedStrings(t gen.Task) (string, string) {

	if !t.DueDateTimestamp.Valid && !t.CompletedTimestamp.Valid {
		return "-", "-"
	}

	if !t.DueDateTimestamp.Valid {
		return "-", timediff.TimeDiff(t.CompletedTimestamp.Time)
	}

	dueSince := time.Since(t.DueDateTimestamp.Time)
	if !t.CompletedTimestamp.Valid {
		return timediff.TimeDiff(t.DueDateTimestamp.Time), "-"
	}

	completedSince := time.Since(t.CompletedTimestamp.Time)
	diff := completedSince - dueSince

	rounding := time.Second
	if diff.Abs().Hours() >= 1 {
		rounding = time.Hour
	} else if diff.Abs().Minutes() >= 1 {
		rounding = time.Minute
	}

	if dueSince < completedSince {
		c := color.New(color.FgGreen)
		return timediff.TimeDiff(t.DueDateTimestamp.Time), c.Sprintf("%s early", diff.Abs().Round(rounding))
	} else {
		c := color.New(color.FgRed)
		return timediff.TimeDiff(t.DueDateTimestamp.Time), c.Sprintf("%s late", diff.Abs().Round(rounding))
	}

}

func init() {
	listCmd.Flags().BoolVarP(&listAll, "all", "a", false, "include completed tasks")
	rootCmd.AddCommand(listCmd)
}
