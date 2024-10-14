/*
Copyright © 2024 Patrick McCarty <patricksantos1234567@gmail.com>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

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

		f, err := os.Open(Tasks)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		defer f.Close()

		csvReader := csv.NewReader(f)
		header, err := csvReader.Read()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		fmt.Println(header)
		for {
			record, err := csvReader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			fmt.Println(record)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&all, "all", "a", false, "include completed tasks")
}
