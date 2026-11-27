package field

import (
	"gosim/entities"
	"gosim/position"
)

type Field struct {
	Rows      int
	Cols      int
	Cells     map[position.Position]entities.Occupier
	Creatures map[position.Position]entities.Creature
}

func NewField(rows, cols int) *Field {
	cells := make(map[position.Position]entities.Occupier, rows*cols)
	creatures := make(map[position.Position]entities.Creature)

	for y := 1; y <= rows; y++ {
		for x := 1; x <= cols; x++ {
			cells[position.Position{X: x, Y: y}] = nil
		}
	}
	return &Field{Rows: rows, Cols: cols, Cells: cells, Creatures: creatures}
}
