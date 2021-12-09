package actions

import (
	"fmt"
	"net/http"
	"todos/app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
)

func Index(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	todos := models.Todos{}
	err := tx.All(&todos)
	if err != nil {
		return err
	}
	todosType := fmt.Sprintf("%T", todos)
	fmt.Println(todosType)
	for _, todo := range todos {
		fmt.Println("===========TODO")
		fmt.Println(todo)
	}

	return c.Render(http.StatusOK, r.HTML("todos/index.plush.html"))
}
