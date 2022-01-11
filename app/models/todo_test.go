package models

import (
	"github.com/gobuffalo/nulls"
	"time"
)

func (ms *ModelSuite) Test_Todo_Create() {
	count, err := ms.DB.Count("todos")
	ms.NoError(err)
	ms.Equal(0, count)

	t := &Todo{
		Title:       "ToDo Test",
		Details:     nulls.NewString("It's really not much to add"),
		LimitDate:   time.Now().AddDate(0, 0, 15),
		IsCompleted: false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = ms.DB.Create(t)
	ms.NoError(err)

	count, err = ms.DB.Count("todos")
	ms.NoError(err)
	ms.Equal(1, count)
}
