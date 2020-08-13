package main

import (
	"time"

	"github.com/go-gorp/gorp"
)

// Comment is a comment for the db.
// Text is a required field to POST.
type Comment struct {
	ID      int64     `json:"id" db:"id,primarykey,autoincrement"`
	Name    string    `json:"name" db:"name,notnull,size:200"`
	Text    string    `validate:"required" json:"text" db:"text,notnull,size:400"`
	Created time.Time `json:"created" db:"created,notnull"`
	Updated time.Time `json:"updated" db:"updated,notnull"`
}

// PreInsert is a function to be executed before insert.
func (c *Comment) PreInsert(s gorp.SqlExecutor) error {
	if c.Name == "" {
		c.Name = "名無し"
	}
	c.Created = time.Now()
	c.Updated = c.Created
	return nil
}

// PreUpdate is a function to be executed before update.
func (c *Comment) PreUpdate(s gorp.SqlExecutor) error {
	c.Updated = time.Now()
	return nil
}
