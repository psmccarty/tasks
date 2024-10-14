/*
Copyright © 2024 Patrick McCarty <patricksantos1234567@gmail.com>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete <taskid>",
	Short: "Marks a task as complete",
	Long:  `Mark a task complete given its taskid.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.OpenFile(Tasks, os.O_RDWR, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		defer f.Close()

		csvReader := csv.NewReader(f)
		records, err := csvReader.ReadAll()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if len(records) < 2 {
			fmt.Fprintln(os.Stderr, "task not found")
			return
		}

		for i := 1; i < len(records); i++ {
			if records[i][0] == args[0] {
				complete, err := strconv.ParseBool(records[i][0])
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					return
				}
				if complete {
					fmt.Println("task already complete")
					return
				}
				records[i][3] = strconv.FormatBool(true)
				f.Truncate(0)
				f.Seek(0, 0)
				csvWriter := csv.NewWriter(f)
				csvWriter.WriteAll(records)
				csvWriter.Flush()
				return
			}
		}
		fmt.Fprintln(os.Stderr, "task not found")
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
