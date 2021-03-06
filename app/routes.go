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

	root.Use(middleware.Authorize)
	root.Use(middleware.SetCurrentUser)
	root.Use(middleware.UncompletedTodos)

	root.GET("/", actions.Home).Name("homeTodoPath")

	root.GET("/users/new", actions.UsersNew)
	root.POST("/users", actions.UsersCreate).Name("saveUserPath")
	root.GET("/signin", actions.AuthNew)
	root.POST("/signin", actions.AuthCreate).Name("saveAuthPath")
	root.DELETE("/signout", actions.AuthDestroy)

	root.GET("/todos", actions.Index).Name("listTodoPath")
	root.GET("/todo/new", actions.NewTodo)
	root.POST("/todo", actions.SaveTodo).Name("saveTodoPath")
	root.GET("/todo/{todo_id}/edit", actions.EditTodo).Name("editTodoPath")
	root.PUT("/todo/{todo_id}", actions.UpdateTodo).Name("updateTodoPath")
	root.PUT("/todo/updateStatus/{todo_id}", actions.UpdateTodoStatus).Name("updateTodoStatusPath")
	root.DELETE("/todo/{todo_id}/", actions.DeleteTodo).Name("deleteTodoPath")

	root.Middleware.Skip(middleware.Authorize, actions.Home, actions.UsersNew, actions.UsersCreate, actions.AuthNew, actions.AuthCreate)
	root.Middleware.Skip(middleware.SetCurrentUser, actions.UsersNew, actions.UsersCreate, actions.AuthNew, actions.AuthCreate)
	root.Middleware.Skip(middleware.UncompletedTodos, actions.UsersNew, actions.UsersCreate, actions.AuthNew, actions.AuthCreate)
	root.ServeFiles("/", base.Assets)
}
