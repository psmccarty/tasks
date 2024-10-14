/*
Copyright © 2024 Patrick McCarty <patricksantos1234567@gmail.com>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete <taskid>",
	Short: "Deletes a task from the list",
	Long:  `Deletes a task given its taskid.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.OpenFile(Tasks, os.O_RDWR, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		csvReader := csv.NewReader(f)
		records, err := csvReader.ReadAll()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		newRecords := make([][]string, 0, len(records))
		newRecords = append(newRecords, records[0])

		for i := 1; i < len(records); i++ {
			if records[i][0] != args[0] {
				newRecords = append(newRecords, records[i])
			}
		}

		if len(newRecords) == len(records) {
			fmt.Fprintln(os.Stderr, "task not found")
			return
		}

		err = f.Truncate(0)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		_, err = f.Seek(0, 0)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		csvWriter := csv.NewWriter(f)
		err = csvWriter.WriteAll(newRecords)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
