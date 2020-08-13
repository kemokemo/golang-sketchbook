package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type exitcode int

const (
	exitOK exitcode = iota
	exitFailOpenDB
	exitFailPingDB
	exitFailExecDB
	exitFailGetAffectedRow
	exitFailQueryDB
	exitFailScanRow
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

	err = db.Ping()
	if err != nil {
		log.Printf("failed to ping db: %v", err)
		return exitFailPingDB
	}

	row := db.QueryRow(`SELECT name, age FROM users WHERE id=$1`, 1)
	if row == nil {
		log.Printf("failed to select a row from the users table: %v", err)
		return exitFailQueryDB
	}

	var name string
	var age int64
	err = row.Scan(&name, &age)
	if err != nil {
		log.Printf("failed to scan a row: %v", err)
		return exitFailScanRow
	}
	fmt.Println(name, age)

	return exitOK
}
