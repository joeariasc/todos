package actions

import (
	"github.com/gobuffalo/buffalo"
	"net/http"
)

func Home(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("home/index.plush.html"))
}