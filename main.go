/*
Copyright © 2024 Patrick McCarty <patricksantos1234567@gmail.com>
*/
package main

import (
	"context"
	"database/sql"
	_ "embed"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/psmccarty/tasks/cmd"
)

//go:embed sql/schema.sql
var createTableIfNotExits string

func main() {
	db, err := sql.Open("sqlite3", "tasks.db")
	if err != nil {
		os.Exit(1)
	}

	ctx := context.Background()
	// create tables
	if _, err := db.ExecContext(ctx, createTableIfNotExits); err != nil {
		os.Exit(1)
	}
	db.Close()

	cmd.Execute()
}
