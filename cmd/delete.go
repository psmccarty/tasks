/*
Copyright © 2024 Patrick McCarty <patricksantos1234567@gmail.com>
*/
package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/psmccarty/tasks/sql/gen"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete <taskid>",
	Short: "Deletes a task from the list",
	Long:  `Deletes a task given its taskid.`,
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
		res, err := queries.DeleteTask(ctx, int64(id))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Deleted task number %d, %q\n", res.ID, res.Description)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
