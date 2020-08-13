package main

import (
	"gopkg.in/go-playground/validator.v9"
)

// Validator validates the data before inserting and updationg database.
type Validator struct {
	validator *validator.Validate
}

// Validate validates the data.
func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
