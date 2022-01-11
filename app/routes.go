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

	root.Use(middleware.SetCurrentUser)
	//root.Use(middleware.Authorize)

	root.GET("/users/new", actions.UsersNew)
	root.POST("/users", actions.UsersCreate)
	root.GET("/signin", actions.AuthNew)
	root.POST("/signin", actions.AuthCreate)
	root.DELETE("/signout", actions.AuthDestroy)

	root.GET("/", actions.Index).Name("listTodoPath")
	root.GET("/todo/new", actions.NewTodo)
	root.POST("/todo", actions.SaveTodo).Name("saveTodoPath")
	root.GET("/todo/{todo_id}/edit", actions.EditTodo).Name("editTodoPath")
	root.PUT("/todo/{todo_id}", actions.UpdateTodo).Name("updateTodoPath")
	root.DELETE("/todo/{todo_id}/", actions.DeleteTodo).Name("deleteTodoPath")
	root.ServeFiles("/", base.Assets)
}
