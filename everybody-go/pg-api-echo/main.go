package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"gopkg.in/go-playground/validator.v9"

	"github.com/go-gorp/gorp"
	"github.com/labstack/echo"
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

	e := echo.New()
	e.Validator = &Validator{validator: validator.New()}
	e.GET("/greeting", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello echo!")
	})
	e.Static("/", "static/")

	e.GET("/api/comments", func(c echo.Context) error {
		var comments []Comment
		_, err := dbmap.Select(&comments, "SELECT * FROM comments ORDER BY created desc LIMIT 10")
		if err != nil {
			c.Logger().Error("Select: ", err)
			return c.String(http.StatusBadRequest, "Select: "+err.Error())
		}
		return c.JSON(http.StatusOK, comments)
	})

	e.POST("/api/comments", func(c echo.Context) error {
		var comment Comment
		if err := c.Bind(&comment); err != nil {
			c.Logger().Error("Bind: ", err)
			return c.String(http.StatusBadRequest, "Bind: "+err.Error())
		}
		if err := c.Validate(&comment); err != nil {
			c.Logger().Error("Validate: ", err)
			return c.String(http.StatusBadRequest, "Validate: "+err.Error())
		}
		if err = dbmap.Insert(&comment); err != nil {
			c.Logger().Error("Insert: ", err)
			return c.String(http.StatusBadRequest, "Insert: "+err.Error())
		}
		c.Logger().Infof("ADDED: ID is %v", comment.ID)
		return c.JSON(http.StatusOK, "")
	})

	e.Logger.Fatal(e.Start(":8080"))

	return exitOK
}
