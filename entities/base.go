package entities

import (
	"gosim/position"
)

type EntityType string

const (
	TypeGrass          EntityType = "grass"
	TypeRock           EntityType = "rock"
	TypeTree           EntityType = "tree"
	TypeHerbivore      EntityType = "herbivore"
	TypePredator       EntityType = "predator"
	MaxHerbivoreHp     int        = 5
	MaxGrassHp         int        = 6
	MaxHerbivoreEnergy int        = 4
	MaxPredatorEnergy  int        = 2
)

type Occupier interface {
	Type() EntityType
	Pos() position.Position
}

type Creature interface {
	Occupier
	SetPos(newPos position.Position)
	SetEnergy(newSpeed int)
}

type Eatable interface {
	SetHp(newHp int)
}
