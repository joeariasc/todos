package actions_test

import (
	"todos/app/models"
)

func (as *ActionSuite) Test_Auth_New() {
	res := as.HTML("/signin").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Sign In")
}

func (as *ActionSuite) Test_Auth_Create() {
	u := &models.User{
		FirstName:            "Joe",
		LastName:             "Arias",
		Email:                "joe@example.com",
		Password:             "password",
		PasswordConfirmation: "password",
	}
	verrs, err := u.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())

	res := as.HTML("/signin").Post(u)
	as.Equal(302, res.Code)
	as.Equal("/todos/", res.Location())
}

func (as *ActionSuite) Test_Auth_Create_UnknownUser() {
	u := &models.User{
		Email:    "joe@example.com",
		Password: "password",
	}
	res := as.HTML("/signin").Post(u)
	as.Equal(422, res.Code)
	as.Contains(res.Body.String(), "invalid email/password")
}

func (as *ActionSuite) Test_Auth_Create_BadPassword() {
	u := &models.User{
		FirstName:            "Joe",
		LastName:             "Arias",
		Email:                "joe@example.com",
		Password:             "password",
		PasswordConfirmation: "password",
	}
	verrs, err := u.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())

	u.Password = "bad"
	res := as.HTML("/signin").Post(u)
	as.Equal(422, res.Code)
	as.Contains(res.Body.String(), "invalid email/password")
}
