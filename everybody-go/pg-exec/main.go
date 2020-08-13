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
	eixtFailExecDB
	eixtFailGetAffectedRow
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

	result, err := db.Exec(`INSERT INTO users(name,age) VALUES($1, $2)`, "Bob", 18)
	if err != nil {
		log.Printf("failed to insert row to the users table: %v", err)
		return eixtFailExecDB
	}

	affected, err := result.RowsAffected()
	if err != nil {
		log.Printf("failed to get the affected row number: %v", err)
		return eixtFailGetAffectedRow
	}

	fmt.Printf("Affected row number: %v\n", affected)

	return exitOK
}
