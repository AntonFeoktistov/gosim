package actions

import (
	"errors"
	"gosim/entities"
	"gosim/field"
	"gosim/position"
)

func MoveCreature(fld *field.Field, creature entities.Creature, newPos position.Position) (bool, error) {
	if !isCellFreeToMove(fld, newPos) {
		return false, errors.New("cannot move to this cell")
	}

	oldPos := creature.Pos()
	delete(fld.Creatures, oldPos)
	fld.Creatures[newPos] = creature
	creature.SetPos(newPos)

	dist := position.GetDistance(oldPos, newPos)
	newSpeed := creature.GetEnergy() - dist
	creature.SetEnergy(newSpeed)

	return true, nil
}

func FindAllAvailableCells(fld *field.Field, creature entities.Creature) []position.Position {
	availableCells := make([]position.Position, 0)
	for cell := range fld.Cells {
		if isCellFreeToMove(fld, cell) && isEnergyEnoughToMove(creature, cell) {
			availableCells = append(availableCells, cell)
		}
	}
	return availableCells
}

func isCellFreeToMove(fld *field.Field, newPos position.Position) bool {
	_, ok := fld.Cells[newPos]
	if !ok {
		return false
	}
	if _, ok := fld.Creatures[newPos]; ok {
		return false
	}
	return true
}

func isEnergyEnoughToMove(creature entities.Creature, targetCell position.Position) bool {
	dist := position.GetDistance(creature.Pos(), targetCell)
	return dist <= creature.GetEnergy()
}
