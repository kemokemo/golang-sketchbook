package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/mattn/go-nulltype"
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

// User is the user data.
type User struct {
	ID   int64
	Name nulltype.NullString `json:"name"`
	Age  int64
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

	var user User
	err = db.QueryRow(`SELECT name FROM users`).Scan(&user.Name)
	if err != nil {
		log.Printf("failed to select name from the users table: %v", err)
	}

	fmt.Printf("user.Name.Valid(): %v\n", user.Name.Valid())
	fmt.Printf("user.Name: %v\n", user.Name)

	return exitOK
}
