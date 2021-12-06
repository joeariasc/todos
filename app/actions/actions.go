package actions

import(
	"todos/app/render"
)

var (
	// r is a buffalo/render Engine that will be used by actions 
	// on this package to render render HTML or any other formats.
	r = render.Engine
)