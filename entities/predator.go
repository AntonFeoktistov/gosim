package entities

import "gosim/position"

type Predator struct {
	Hp       int
	Position position.Position
	Energy   int
}

func NewPredator(pos position.Position) *Predator {
	return &Predator{Hp: MaxHerbivoreHp, Position: pos, Energy: MaxPredatorEnergy}
}

func (p *Predator) Type() EntityType {
	return TypePredator
}

func (p *Predator) Pos() position.Position {
	return p.Position
}

func (p *Predator) SetPos(newPos position.Position) {
	p.Position = newPos
}

func (p *Predator) SetEnergy(newEnergy int) {
	p.Energy = newEnergy
}
