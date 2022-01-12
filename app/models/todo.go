package models

import (
	"fmt"
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

// Todo model struct.
type Todo struct {
	ID          uuid.UUID    `json:"id" db:"id"`
	Title       string       `json:"title" db:"title"`
	Details     nulls.String `json:"details" db:"details"`
	LimitDate   time.Time    `json:"limit_date" db:"limit_date"`
	IsCompleted bool         `json:"is_completed" db:"is_completed"`
	UserID      uuid.UUID    `json:"user_id" db:"user_id"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at" db:"updated_at"`
}

// TableName overrides the table name used by Pop.
func (t Todo) TableName() string {
	return "todos"
}

// Todos array model struct of Todo.
type Todos []Todo

// String converts the struct into a string value.
func (t Todo) String() string {
	return fmt.Sprintf("%#v", t)
}

func (t *Todo) Validate() *validate.Errors {
	return validate.Validate(
		&validators.StringIsPresent{
			Field:   t.Title,
			Name:    "Title",
			Message: "Title cannot be empty",
		},
		&validators.StringLengthInRange{
			Field:   t.Details.String,
			Name:    "Details",
			Min:     10,
			Max:     255,
			Message: "Detail must be at least 10 characters",
		},
		&validators.TimeIsPresent{
			Name:    "Limit Date",
			Field:   t.LimitDate,
			Message: "Limit Date can not be blank",
		},
		&ValidLimitDateAfterToday{
			FirstName:  "Limit Date",
			LimitDate:  t.LimitDate,
			SecondName: "Today",
			Today:      time.Now(),
		},
	)
}

type ValidLimitDateAfterToday struct {
	FirstName  string
	LimitDate  time.Time
	SecondName string
	Today      time.Time
	Message    string
}

func (v *ValidLimitDateAfterToday) IsValid(errors *validate.Errors) {
	if v.LimitDate.After(v.Today) {
		return
	}
	errors.Add(validators.GenerateKey(v.FirstName), fmt.Sprintf("%s must be after %s.", v.FirstName, v.SecondName))
}
