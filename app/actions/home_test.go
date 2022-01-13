package actions_test

import (
	"net/http"
	"todos/app/models"
)

func (as ActionSuite) Test_Home_Todos_Page_User_Not_Logged_In() {
	res := as.HTML("/").Get()
	as.Equal(http.StatusOK, res.Code)

	body := res.Body.String()
	as.Contains(body, "Welcome to ToDo app")
	as.Contains(body, "Sign In")
	as.Contains(body, "Register")
}

func (as ActionSuite) Test_Home_Todos_Page_User_Logged_In() {
	currentUser := models.User{
		FirstName: "Joe",
		LastName:  "Arias",
		Email:     "jarias@testing.com",
	}

	tx := as.DB
	as.NoError(tx.Create(&currentUser))

	as.Session.Set("current_user", currentUser)
	as.Session.Set("current_user_id", currentUser.ID)
	as.Session.Save()
	res := as.HTML("/").Get()
	as.Equal(http.StatusOK, res.Code)

	body := res.Body.String()
	as.Contains(body, "Check your last ToDos!")
	as.NotContains(body, "Sign In")
	as.NotContains(body, "Register")
}
