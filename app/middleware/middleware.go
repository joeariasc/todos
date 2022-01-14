// middleware package is intended to host the middlewares used
// across the app.
package middleware

import (
	"net/http"
	"todos/app/models"

	tx "github.com/gobuffalo/buffalo-pop/v2/pop/popmw"
	csrf "github.com/gobuffalo/mw-csrf"
	paramlogger "github.com/gobuffalo/mw-paramlogger"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/pkg/errors"
)

var (
	// Transaction middleware wraps the request with a pop
	// transaction that is committed on success and rolled
	// back when errors happen.
	Transaction = tx.Transaction(models.DB())

	// ParameterLogger logs out parameters that the app received
	// taking care of sensitive data.
	ParameterLogger = paramlogger.ParameterLogger

	// CSRF middleware protects from CSRF attacks.
	CSRF = csrf.New
)

// SetCurrentUser attempts to find a user based on the current_user_id
// in the session. If one is found it is set on the context.
func SetCurrentUser(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid != nil {
			u := &models.User{}
			tx := c.Value("tx").(*pop.Connection)
			err := tx.Find(u, uid)
			if err != nil {
				return errors.WithStack(err)
			}
			c.Set("current_user", u)
		}
		return next(c)
	}
}

// Authorize require a user be logged in before accessing a route
func Authorize(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid == nil {
			c.Flash().Add("danger", "You must be authorized to see that page")
			return c.Redirect(302, "/")
		}
		return next(c)
	}
}

func UncompletedTodos(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		currentUserID := c.Session().Get("current_user_id")
		tx := c.Value("tx").(*pop.Connection)
		count, err := tx.Where("user_id = ?", currentUserID).Where("is_completed = ?", false).Count(&models.Todos{})

		if err != nil {
			return c.Error(http.StatusInternalServerError, errors.Wrap(err, "UncompletedTodos MW"))
		}

		c.Set("uncompleted_todos", count)
		return next(c)
	}
}
