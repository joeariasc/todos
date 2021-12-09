package models

import (
	"fmt"
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/gofrs/uuid"
)

// Todo model struct.
type Todo struct {
	ID          uuid.UUID    `json:"id" db:"id"`
	Title       string       `json:"title" db:"title"`
	Details     nulls.String `json:"details" db:"details"`
	LimitDate time.Time    `json:"limit_date" db:"limit_date"`
	IsCompleted bool         `json:"is_completed" db:"is_completed"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at" db:"updated_at"`
}

// TableName overrides the table name used by Pop.
func (t Todo) TableName() string {
	return "todos"
}

// Todoes array model struct of Todo.
type Todos []Todo

// String converts the struct into a string value.
func (t Todo) String() string {
	return fmt.Sprintf("%#v", t)
}
