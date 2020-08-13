package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type exitcode int

const (
	exitOK exitcode = iota
	exitFailOpenDB
	exitFailPingDB
)

func main() {
	os.Exit(int(run()))
}

func run() exitcode {
	dsn := os.Getenv("DSN")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Printf("failed to open db: %v¥n", err)
		return exitFailOpenDB
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Printf("failed to ping db: %v", err)
		return exitFailPingDB
	}

	return exitOK
}
