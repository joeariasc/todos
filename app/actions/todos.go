package actions

import (
	"net/http"
	"todos/app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/pkg/errors"
)

func Index(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	todos := models.Todos{}
	err := tx.All(&todos)
	if err != nil {
		return c.Error(http.StatusInternalServerError, errors.Wrap(err, "Index - Error while getting all todos"))
	}
	c.Set("todos", todos)
	return c.Render(http.StatusOK, r.HTML("todos/index.plush.html"))
}

func NewTodo(c buffalo.Context) error {
	c.Set("todo", models.Todo{})
	return c.Render(http.StatusOK, r.HTML("todos/new.plush.html"))
}

func SaveTodo(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	todo := models.Todo{}
	if err := c.Bind(&todo); err != nil {
		return c.Error(http.StatusInternalServerError, errors.Wrap(err, "Store - Error while bind a todo"))
	}
	if err := tx.Create(&todo); err != nil {
		return c.Error(http.StatusInternalServerError, errors.Wrap(err, "Store - Error while saving a todo"))
	}
	return c.Redirect(http.StatusSeeOther, "listTodoPath()")
}

func EditTodo(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	todoID := c.Param("todo_id")
	todo := models.Todo{}
	if err := tx.Find(&todo, todoID); err != nil {
		return c.Error(http.StatusNotFound, errors.Wrap(err, "Edit - todo not found"))
	}
	c.Set("todo", todo)
	return c.Render(http.StatusOK, r.HTML("todos/edit.plush.html"))
}

func UpdateTodo(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	todoID := c.Param("todo_id")
	todo := models.Todo{}
	if err := tx.Find(&todo, todoID); err != nil {
		return c.Error(http.StatusNotFound, errors.Wrap(err, "Update - todo not found"))
	}
	if err := c.Bind(&todo); err != nil {
		return c.Error(http.StatusInternalServerError, errors.Wrap(err, "Update - Error while bind a todo"))
	}
	if err := tx.Update(&todo); err != nil {
		return c.Error(http.StatusInternalServerError, errors.Wrap(err, "Update - Error while updating a todo"))
	}
	return c.Redirect(http.StatusSeeOther, "listTodoPath()")
}

func DeleteTodo(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	todo := models.Todo{}
	todoID := c.Param("todo_id")
	if err := tx.Find(&todo, todoID); err != nil {
		return c.Render(http.StatusNotFound, r.String("ToDo not Found"))
	}
	if err := tx.Destroy(&todo); err != nil {
		return c.Error(http.StatusInternalServerError, errors.Wrap(err, "Delete - Error while destroy a todo"))
	}
	return c.Redirect(http.StatusSeeOther, "listTodoPath()")
}
