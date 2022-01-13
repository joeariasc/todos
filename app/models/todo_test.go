package models_test

import (
	"github.com/gobuffalo/nulls"
	"time"
	"todos/app/models"
)

func (ms *ModelSuite) Test_Todo_Create() {
	u := ms.CreateUser()

	count, err := ms.DB.Count("todos")
	ms.NoError(err)
	ms.Equal(0, count)

	t := &models.Todo{
		Title:       "ToDo Test",
		Details:     nulls.NewString("It's really not much to add"),
		LimitDate:   time.Now().AddDate(0, 0, 15),
		IsCompleted: false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		UserID:      u.ID,
	}

	err = ms.DB.Create(t)
	ms.NoError(err)

	count, err = ms.DB.Count("todos")
	ms.NoError(err)
	ms.Equal(1, count)
}
