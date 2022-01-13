package models_test

import (
	"testing"
	"todos/app/models"

	"github.com/gobuffalo/suite/v3"
)

type ModelSuite struct {
	*suite.Model
}

func Test_ModelSuite(t *testing.T) {
	suite.Run(t, &ModelSuite{
		Model: suite.NewModel(),
	})
}

func (ms *ModelSuite) CreateUser() *models.User {
	user := &models.User{
		FirstName:            "Joe",
		LastName:             "Arias",
		Email:                "jarias@testing.com",
		Password:             "joe123",
		PasswordConfirmation: "joe123",
	}
	verrs, err := user.Create(ms.DB)

	ms.NoError(err)
	ms.False(verrs.HasAny())

	return user
}
