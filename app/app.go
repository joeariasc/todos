package app

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
)

var (
	root  *buffalo.App
)

// App creates a new application with default settings and reading 
// GO_ENV. It calls setRoutes to setup the routes for the app that's being
// created before returning it
func New() *buffalo.App {
	if root != nil {
		return root
	}

	root = buffalo.New(buffalo.Options{
		Env:         envy.Get("GO_ENV", "development"),
		SessionName: "_todos_session",
	})

	// Setting the routes for the app
	setRoutes(root)

	return root
}