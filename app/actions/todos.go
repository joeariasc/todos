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

func NewTodo(c buffalo.Context) error  {
	c.Set("todo", models.Todo{})
	return c.Render(http.StatusOK, r.HTML("todos/new.plush.html"))
}
