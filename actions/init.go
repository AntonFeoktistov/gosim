package actions

import (
	"gosim/entities"
	"gosim/field"
	"gosim/position"
	"math/rand"
)

func InitField(fld *field.Field) {
	for y := 1; y <= fld.Rows; y++ {
		for x := 1; x <= fld.Cols; x++ {
			pos := position.Position{X: x, Y: y}
			ent := chooseRandomNotAliveEntity(pos)
			fld.Cells[pos] = ent
			creature, ok := chooseRandomCreatureOrNil(pos)
			if ok {
				fld.Creatures[pos] = creature
			}
		}
	}
}

func chooseRandomNotAliveEntity(pos position.Position) entities.Occupier {
	n := rand.Intn(100) + 1 // [1, 100]

	switch {
	case n <= 60:
		return entities.NewGrass(pos)
	case n <= 80:
		return entities.NewRock(pos)
	default:
		return entities.NewTree(pos)
	}
}

func chooseRandomCreatureOrNil(pos position.Position) (entities.Creature, bool) {
	n := rand.Intn(100) + 1 // [1, 100]

	switch {
	case n <= 20:
		return entities.NewHerbivore(pos), true
	case n <= 30:
		return entities.NewPredator(pos), true
	default:
		return nil, false
	}
}
