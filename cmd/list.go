/*
Copyright © 2024 Patrick McCarty <patricksantos1234567@gmail.com>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"

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

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
		defer w.Flush()

		if all {
			fmt.Fprintln(w, strings.Join(header[:], "\t"))
		} else {
			fmt.Fprintln(w, strings.Join(header[:len(header)-1], "\t"))
		}

		for {
			record, err := csvReader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			complete, err := strconv.ParseBool(record[len(record)-1])
			if err != nil {
				fmt.Fprint(os.Stderr, err)
				return
			}
			if all {
				fmt.Fprintln(w, strings.Join(record[:], "\t"))
			} else if !complete {
				fmt.Fprintln(w, strings.Join(record[:len(record)-1], "\t"))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&all, "all", "a", false, "include completed tasks")
}
