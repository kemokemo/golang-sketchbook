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

	rows, err := db.Query(`SELECT id, name, age FROM users ORDER BY name`)
	if err != nil {
		log.Printf("failed to select rows from the users table: %v", err)
		return exitFailQueryDB
	}

	for rows.Next() {
		var id int64
		var name string
		var age int64
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			log.Printf("failed to scan rows: %v", err)
			return exitFailScanRow
		}
		fmt.Println(id, name, age)
	}

	return exitOK
}
