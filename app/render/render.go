package render

import (
	"time"
	base "todos"

	"github.com/gobuffalo/buffalo/render"
)

// Engine for rendering across the app, it provides
// the base for rendering HTML, JSON, XML and other formats
// while also defining thing like the base layout.
var Engine = render.New(render.Options{
	HTMLLayout:   "application.plush.html",
	TemplatesBox: base.Templates,
	AssetsBox:    base.Assets,
	Helpers:      Helpers,
})

// Helpers available for the plush templates, there are
// some helpers that are injected by Buffalo but this is
// the list of custom Helpers.
var Helpers = map[string]interface{}{
	// partialFeeder is the helper used by the render engine
	// to find the partials that will be used, this is important
	"partialFeeder": base.Templates.FindString,
	"todayDate":     todayDate,
	"valueDateToDo": valueDateToDo,
	"minDateTodo":   minDateTodo,
}

func todayDate() string {
	now := time.Now().Format("Monday 02, January 2006")
	return now
}

func valueDateToDo(t time.Time) time.Time {
	if t.IsZero() {
		now := time.Now()
		return now
	}
	return t
}

func minDateTodo() time.Time {
	now := time.Now()
	return now
}
