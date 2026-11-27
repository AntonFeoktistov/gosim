package entities

import "gosim/position"

type Tree struct {
	Position position.Position
}

func NewTree(pos position.Position) *Tree {
	return &Tree{Position: pos}
}

func (t *Tree) Type() EntityType {
	return TypeTree
}

func (t *Tree) Pos() position.Position {
	return t.Position
}
