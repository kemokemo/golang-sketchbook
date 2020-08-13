package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-gorp/gorp"

	_ "github.com/lib/pq"
)

type exitcode int

const (
	exitOK exitcode = iota
	exitFailOpenDB
	exitFailCreateDB
	exitFailPingDB
	exitFailQuery
)

func main() {
	os.Exit(int(run()))
}

func run() exitcode {
	dsn := os.Getenv("DSN")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Printf("failed to open db: %v", err)
		return exitFailOpenDB
	}
	defer db.Close()

	// Create table
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbmap.AddTableWithName(Comment{}, "comments").SetKeys(true, "id")
	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		log.Printf("failed to create commnet db: %v", err)
		return exitFailCreateDB
	}

	// Pre-check with ping
	err = db.Ping()
	if err != nil {
		log.Printf("failed to ping db: %v", err)
		return exitFailPingDB
	}

	// Query
	err = dbmap.Insert(&Comment{Name: "bob", Text: "こんにちは"})
	if err != nil {
		log.Printf("failed to insert a Comment to db: %v", err)
		return exitFailQuery
	}

	var comment Comment
	err = dbmap.SelectOne(&comment, "SELECT * FROM comments WHERE id = 1")
	if err != nil {
		log.Printf("failed to select one Comment from db: %v", err)
		return exitFailQuery
	}

	comment.Text = "こんばんは"
	n, err := dbmap.Update(&comment)
	fmt.Printf("The number of the updated comments: %v\n", n)

	var comments []Comment
	_, err = dbmap.Select(&comments, "SELECT * FROM comments WHERE name = $1", "bob")
	if err != nil {
		log.Printf("failed to select Comments from db: %v", err)
		return exitFailQuery
	}

	return exitOK
}
