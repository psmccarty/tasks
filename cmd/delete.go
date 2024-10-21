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

var (
	deleteAll bool
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete <taskid>",
	Short: "Deletes a task from the list",
	Long:  `Deletes a task given its taskid.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if !deleteAll {
			return cobra.ExactArgs(1)(cmd, args)
		}
		return cobra.ExactArgs(0)(cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {

		db, err := sql.Open("sqlite3", "tasks.db")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer db.Close()

		ctx := context.Background()
		queries := gen.New(db)

		if deleteAll {
			err = queries.DeleteList(ctx)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Deleted all tasks")
			return
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		description, err := queries.DeleteTask(ctx, int64(id))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Deleted task number %d, %q\n", id, description)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().BoolVarP(&deleteAll, "all", "A", false, "delete all tasks")
}
