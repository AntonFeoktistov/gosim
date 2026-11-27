package entities

import "gosim/position"

type Herbivore struct {
	Position position.Position
	Hp       int
	Energy   int
}

func NewHerbivore(pos position.Position) *Herbivore {
	return &Herbivore{Position: pos, Hp: MaxHerbivoreHp, Energy: MaxHerbivoreEnergy}
}

func (h *Herbivore) Type() EntityType {
	return TypeHerbivore
}

func (h *Herbivore) Pos() position.Position {
	return h.Position
}

func (h *Herbivore) SetPos(newPos position.Position) {
	h.Position = newPos
}

func (h *Herbivore) SetEnergy(newEnergy int) {
	h.Energy = newEnergy
}

func (h *Herbivore) SetHp(newHp int) {
	h.Hp = newHp
}

func (h *Herbivore) Target() EntityType {
	return TypeGrass
}
