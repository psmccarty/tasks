/*
Copyright © 2024 Patrick McCarty <patricksantos1234567@gmail.com>
*/
package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/psmccarty/tasks/sql/gen"
	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete <taskid>",
	Short: "Marks a task as complete",
	Long:  `Mark a task complete given its taskid.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		db, err := sql.Open("sqlite3", "tasks.db")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer db.Close()

		ctx := context.Background()
		queries := gen.New(db)
		description, err := queries.UpdateComplete(ctx, gen.UpdateCompleteParams{
			ID:                 int64(id),
			CompletedTimestamp: sql.NullTime{Time: time.Now(), Valid: true},
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Completed task number %d, %q\n", id, description)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
