package actions_test

import (
	"net/http"
	"net/url"
	"time"
	"todos/app/models"

	"github.com/gobuffalo/nulls"
)

func (as ActionSuite) Test_User_Todos_Blank_State() {
	// 1. Arrange
	as.Login()

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
	as.Login()

	limitDate := time.Now().AddDate(0, 0, 3)

	form := url.Values{
		"Title":     []string{"Test Title"},
		"Details":   []string{"Some long description"},
		"LimitDate": []string{limitDate.Format("2006-01-02")},
	}

	// 2. Act
	res := as.HTML("/todo").Post(form)

	// 3. Assert
	as.Equal(http.StatusSeeOther, res.Code)
	as.Equal("/todos/", res.Location())

	res = as.HTML("/todos").Get()
	body := res.Body.String()

	as.Contains(body, "Test Title")
	as.Contains(body, limitDate.Format("Monday 02, January 2006"))
	as.Contains(body, "1 Uncompleted ToDo(s)")
}

func (as ActionSuite) Test_Create_Todos_Validation() {
	// 1. Arrange
	as.Login()

	tcases := []struct { // table-driven tests
		URLValues        url.Values
		ExpectedHTTPCode int
		ExpectedLocation string
	}{
		{
			URLValues: url.Values{
				"Title":     []string{"First Todo"},
				"Details":   []string{"first description"},
				"LimitDate": []string{time.Now().AddDate(0, 0, 5).Format("2006-01-02")},
			},
			ExpectedHTTPCode: http.StatusSeeOther,
			ExpectedLocation: "/todos/",
		},
		{
			URLValues: url.Values{
				"Title":     []string{"Second Todo"},
				"Details":   []string{"second description"},
				"LimitDate": []string{time.Now().AddDate(0, 0, -1).Format("2006-01-02")},
			},
			ExpectedHTTPCode: http.StatusUnprocessableEntity,
			ExpectedLocation: "",
		},
	}

	// 2. Act
	for i, tc := range tcases {
		res := as.HTML("/todo").Post(tc.URLValues)

		// 3. Assert
		as.Equal(tc.ExpectedHTTPCode, res.Code, "Case #: %v", i)

		if tc.ExpectedLocation != "" {
			as.Equal(tc.ExpectedLocation, res.Location())
		}
	}
}

func (as ActionSuite) Test_Update_Todos() {
	// 1. Arrange
	currentUser := as.Login()

	todo := models.Todo{
		Title:     "Hello!",
		Details:   nulls.NewString("Really nothing!"),
		LimitDate: time.Now(),
		UserID:    currentUser.ID,
	}

	tx := as.DB

	as.NoError(tx.Create(&todo))

	limitDate := time.Now().AddDate(0, 0, 3)

	form := url.Values{
		"Title":     []string{"Do nothing!"},
		"Details":   []string{"Updated! but maybe tomorrow I'll do it"},
		"LimitDate": []string{limitDate.Format("2006-01-02")},
	}

	// 2. Act
	res := as.HTML("/todo/{%s}", todo.ID).Put(form)

	// 3. Assert
	as.Equal(http.StatusSeeOther, res.Code)

	err := tx.Reload(&todo)
	as.NoError(err)
	as.Equal("Do nothing!", todo.Title)
}

func (as ActionSuite) Test_Delete_Todos() {
	// 1. Arrange
	currentUser := as.Login()

	tx := as.DB

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

func (as ActionSuite) Test_Update_Todo_Status() {
	currentUser := as.Login()
	tx := as.DB

	limitDate := time.Now().AddDate(0, 0, 3)

	todo := models.Todo{
		Title:     "Hello!",
		Details:   nulls.NewString("Really nothing!"),
		LimitDate: limitDate,
		UserID:    currentUser.ID,
	}

	as.NoError(tx.Create(&todo))

	form := url.Values{
		"Title":       []string{"Do nothing!"},
		"Details":     []string{"Updated! but maybe tomorrow I'll do it"},
		"LimitDate":   []string{limitDate.Format("2006-01-02")},
		"IsCompleted": []string{"true"},
	}

	res := as.HTML("/todo/updateStatus/{%s}", todo.ID).Put(form)
	as.Equal(http.StatusSeeOther, res.Code)

	err := tx.Reload(&todo)
	as.NoError(err)
	as.True(todo.IsCompleted)
	as.Equal("/todos/", res.Location())
}
