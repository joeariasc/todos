package actions_test

import (
	"testing"
	"todos/app"
	"todos/app/models"

	"github.com/gobuffalo/suite/v3"
)

type ActionSuite struct {
	*suite.Action
}

func Test_ActionSuite(t *testing.T) {
	as := &ActionSuite{suite.NewAction(app.New())}
	suite.Run(t, as)
}

func (as *ActionSuite) CreateUser() *models.User {
	user := &models.User{
		FirstName:            "Joe",
		LastName:             "Arias",
		Email:                "jarias@testing.com",
		Password:             "joe123",
		PasswordConfirmation: "joe123",
	}
	verrs, err := user.Create(as.DB)

	as.NoError(err)
	as.False(verrs.HasAny())

	return user
}

func (as *ActionSuite) Login() *models.User {
	user := as.CreateUser()
	as.Session.Set("current_user", user)
	as.Session.Set("current_user_id", user.ID)
	err := as.Session.Save()
	as.NoError(err)
	return user
}
