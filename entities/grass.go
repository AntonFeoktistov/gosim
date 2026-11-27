package entities

import (
	"gosim/position"
)

type Grass struct {
	Position position.Position
	Hp       int
}

func NewGrass(pos position.Position) *Grass {
	return &Grass{Position: pos, Hp: MaxGrassHp}
}

func (g *Grass) Type() EntityType {
	return TypeGrass
}

func (g *Grass) Pos() position.Position {
	return g.Position
}

func (g *Grass) SetHp(newHp int) {
	g.Hp = newHp
}
