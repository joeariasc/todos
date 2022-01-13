package actions_test

import (
	"net/http"
	"time"
	"todos/app/models"

	"github.com/gobuffalo/nulls"
)

func (as ActionSuite) Test_User_Todos_Blank_State() {
	// 1. Arrange
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

	// 2. Act
	res := as.HTML("/todos").Get()

	// 3. Assert
	as.Equal(http.StatusOK, res.Code)

	body := res.Body.String()
	as.Contains(body, "Joe Arias")
	as.Contains(body, "No ToDos to show!")
	as.Contains(body, "Sign Out")
}

func (as ActionSuite) Test_User_Todos() {
	// 1. Arrange
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

	todos := models.Todos{
		{
			Title:     "Nothing much to do 1",
			Details:   nulls.NewString("Oh yes, it's a lovely day to make tests"),
			LimitDate: time.Now(),
			UserID:    currentUser.ID,
		},
		{
			Title:     "Nothing much to do 2",
			Details:   nulls.NewString("Oh yes, it's a lovely day to make tests"),
			LimitDate: time.Now(),
			UserID:    currentUser.ID,
		},
	}

	as.NoError(tx.Create(&todos))

	// 2. Act
	res := as.HTML("/todos").Get()

	// 3. Assert
	as.Equal(http.StatusOK, res.Code)

	body := res.Body.String()
	as.Contains(body, "Joe Arias")
	as.Contains(body, "Nothing much to do 1")
	as.Contains(body, "Nothing much to do 2")
}

func (as ActionSuite) Test_Create_Todos_Validation() {
	// 1. Arrange
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

	todo := models.Todo{
		Title:     "",
		Details:   nulls.NewString("Little"),
		LimitDate: time.Now(),
		UserID:    currentUser.ID,
	}

	// 2. Act
	res := as.HTML("/todo").Post(todo)

	// 3. Assert
	as.Equal(http.StatusUnprocessableEntity, res.Code)
}

func (as ActionSuite) Test_Create_Todos() {
	// 1. Arrange
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

	todos := models.Todos{
		{
			Title:     "first Todo",
			Details:   nulls.NewString("first description"),
			LimitDate: time.Now(),
			UserID:    currentUser.ID,
		},
		{
			Title:     "second Todo",
			Details:   nulls.NewString("second description"),
			LimitDate: time.Now().AddDate(0, 0, 10),
			UserID:    currentUser.ID,
		},
	}

	// 2. Act
	for _, todo := range todos {
		res := as.HTML("/todo").Post(todo)

		// 3. Assert
		as.Equal(http.StatusOK, res.Code)
	}
}

func (as ActionSuite) Test_Update_Todos() {
	// 1. Arrange
	currentUser := models.User{
		FirstName: "Joe",
		LastName:  "Arias",
		Email:     "jarias@testing.com",
	}

	tx := as.DB

	as.NoError(tx.Create(&currentUser))
	as.Session.Set("current_user", currentUser)
	as.Session.Set("current_user_id", currentUser.ID)
	err := as.Session.Save()
	as.NoError(err)

	todo := models.Todo{
		Title:     "Hello!",
		Details:   nulls.NewString("Really nothing!"),
		LimitDate: time.Now(),
		UserID:    currentUser.ID,
	}

	as.NoError(tx.Create(&todo))

	form := &models.Todo{
		ID:          todo.ID,
		Title:       "Do nothing!",
		Details:     nulls.NewString("Updated! but maybe tomorrow I'll do it"),
		IsCompleted: false,
		LimitDate:   time.Now().AddDate(0, 0, 15),
		UserID:      currentUser.ID,
	}

	// 2. Act
	res := as.HTML("/todo/{%s}", todo.ID).Put(form)

	// 3. Assert
	as.Equal(http.StatusSeeOther, res.Code)

	err = tx.Reload(&todo)
	as.NoError(err)
	as.Equal("Do nothing!", todo.Title)
}

func (as ActionSuite) Test_Delete_Todos() {
	// 1. Arrange
	currentUser := models.User{
		FirstName: "Joe",
		LastName:  "Arias",
		Email:     "jarias@testing.com",
	}

	tx := as.DB

	as.NoError(tx.Create(&currentUser))
	as.Session.Set("current_user", currentUser)
	as.Session.Set("current_user_id", currentUser.ID)
	err := as.Session.Save()
	as.NoError(err)

	todo := models.Todo{
		Title:     "No one to do!",
		Details:   nulls.NewString("I don't whats going on"),
		LimitDate: time.Now(),
		UserID:    currentUser.ID,
	}

	as.NoError(tx.Create(&todo))

	// 2. Act
	res := as.HTML("/todo/{%s}", todo.ID).Delete()

	// 3. Assert
	as.Equal(http.StatusSeeOther, res.Code)
}
