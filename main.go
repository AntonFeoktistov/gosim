package main

import (
	"gosim/actions"
	"gosim/field"
	"gosim/render"
)

func main() {
	fld := field.NewField(10, 10)
	actions.InitField(fld)
	render.PrintField(*fld)
}
