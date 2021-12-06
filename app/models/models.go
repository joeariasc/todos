package models

import (
	"bytes"
	"log"
	base "todos"
	
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/pop/v5"
)

// Loading connections from database.yml in the pop.Connections
// variable for later usage.
func init() {
	bf, err := base.Config.Find("database.yml")
	if err != nil {
		log.Fatal(err)
	}

	err = pop.LoadFrom(bytes.NewReader(bf))
	if err != nil {
		log.Fatal(err)
	}
}

// DB returns the DB connection for the current environment.
func DB() *pop.Connection {
	c, err := pop.Connect(envy.Get("GO_ENV", "development"))
	if err != nil {
		log.Fatal(err)
	}

	return c
}