package app

import (
	base "todos"
	"todos/app/actions"
	"todos/app/middleware"

	"github.com/gobuffalo/buffalo"
)

// SetRoutes for the application
func setRoutes(root *buffalo.App) {
	root.Use(middleware.Transaction)
	root.Use(middleware.ParameterLogger)
	root.Use(middleware.CSRF)

	root.GET("/", actions.Index)
	root.GET("/todo/new", actions.NewTodo)
	root.POST("/todo", actions.SaveTodo)
	root.DELETE("/todo/{todo_id}/", actions.DeleteTodo)
	root.ServeFiles("/", base.Assets)
}
