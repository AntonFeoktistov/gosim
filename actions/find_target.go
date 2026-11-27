package actions

import (
	"gosim/entities"
	"gosim/field"
	"gosim/position"
)

func FindTarget(fld *field.Field, creature entities.Creature) (position.Position, bool) {
	targetType := creature.Target()
	start := creature.Pos()
	queue := []position.Position{start}
	visited := map[position.Position]bool{start: true}

	directions := []position.Position{
		{X: 0, Y: -1},
		{X: 0, Y: 1},
		{X: -1, Y: 0},
		{X: 1, Y: 0},
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, dir := range directions {
			next := position.Position{
				X: current.X + dir.X,
				Y: current.Y + dir.Y,
			}

			if next.X < 1 || next.X > fld.Cols || next.Y < 1 || next.Y > fld.Rows {
				continue
			}
			if visited[next] {
				continue
			}
			visited[next] = true

			occupier, hasCell := fld.Cells[next]
			creature, hasCreature := fld.Creatures[next]

			if hasCell && occupier != nil && occupier.Type() == targetType {
				return next, true
			}
			if hasCreature && creature != nil && creature.Type() == targetType {
				return next, true
			}

			queue = append(queue, next)
		}
	}

	return position.Position{}, false
}
