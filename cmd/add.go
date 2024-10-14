/*
Copyright © 2024 Patrick McCarty <patricksantos1234567@gmail.com>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <description>",
	Short: "Adds a task to the list",
	Long:  `Adds a task to the todo list application with a short description.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.OpenFile(Tasks, os.O_RDWR, 0644)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			return
		}
		defer f.Close()

		csvReader := csv.NewReader(f)
		records, err := csvReader.ReadAll()
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			return
		}

		id := 1
		if len(records) > 1 {
			lastId, err := strconv.Atoi(records[len(records)-1][0])
			if err != nil {
				fmt.Fprint(os.Stderr, err)
				return
			}
			id = lastId + 1
		}
		csvWriter := csv.NewWriter(f)
		err = csvWriter.Write([]string{strconv.Itoa(id), args[0], time.Now().Format(time.RFC3339), strconv.FormatBool(false)})
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			return
		}
		csvWriter.Flush()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
