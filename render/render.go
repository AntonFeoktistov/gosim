package render

import (
	"fmt"
	"gosim/entities"
	"gosim/field"
	"gosim/position"
)

func PrintField(fld field.Field) {
	for j := 0; j <= fld.Rows+1; j++ {
		for i := 0; i <= fld.Cols+1; i++ {
			if isWall(i, j, fld.Cols, fld.Rows) {
				fmt.Print("#  ")
				continue
			}
			pos := position.Position{X: i, Y: j}
			creature, ok := fld.Creatures[pos]
			if ok {
				printEntity(creature)
				continue
			}
			ent, _ := fld.Cells[pos]
			printEntity(ent)
		}
		fmt.Println()
	}
}

func isWall(i, j int, maxX, maxY int) bool {
	if i == 0 || j == 0 || i == (maxX+1) || j == (maxY+1) {
		return true
	}
	return false
}

func printEntity(ent entities.Occupier) {
	switch ent.Type() {
	case entities.TypeGrass:
		fmt.Print("🌿")
	case entities.TypeTree:
		fmt.Print("🌲")
	case entities.TypeRock:
		fmt.Print("🪨")
	case entities.TypeHerbivore:
		fmt.Print("🐑")
	case entities.TypePredator:
		fmt.Print("🦊")
	}
	fmt.Print(" ")
}
