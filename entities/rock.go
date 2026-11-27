package entities

import "gosim/position"

type Rock struct {
	Position position.Position
}

func NewRock(pos position.Position) *Rock {
	return &Rock{Position: pos}
}

func (r *Rock) Type() EntityType {
	return TypeRock
}

func (r *Rock) Pos() position.Position {
	return r.Position
}
