package main

import (
	"fmt"
	"gosim/actions"
	"gosim/field"
	"gosim/render"
)

func main() {
	fld := field.NewField(10, 10)
	actions.InitField(fld)
	render.PrintField(*fld)
	for _, cr := range fld.Creatures {
		target, _ := actions.FindTarget(fld, cr)
		fmt.Print(cr.Pos())
		fmt.Print(":")
		fmt.Println(target)
	}
}
