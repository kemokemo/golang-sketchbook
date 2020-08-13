package main

import (
	"time"

	"github.com/go-gorp/gorp"
)

// Comment is a comment for the db.
type Comment struct {
	ID      int64     `db:"id,primarykey,autoincrement"`
	Name    string    `db:"name,notnull,default:'名無し',size:200"`
	Text    string    `db:"text,notnull,size:400"`
	Created time.Time `db:"created,notnull"`
	Updated time.Time `db:"updated,notnull"`
}

// PreInsert is a function to be executed before insert.
func (c *Comment) PreInsert(s gorp.SqlExecutor) error {
	c.Created = time.Now()
	c.Updated = c.Created
	return nil
}

// PreUpdate is a function to be executed before update.
func (c *Comment) PreUpdate(s gorp.SqlExecutor) error {
	c.Updated = time.Now()
	return nil
}
