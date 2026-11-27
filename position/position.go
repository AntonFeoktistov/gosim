package position

type Position struct {
	X int
	Y int
}

func GetDistance(pos1, pos2 Position) int {
	dx := pos2.X - pos1.X
	dy := pos2.Y - pos1.Y

	if dx < 0 {
		dx = -dx
	}
	if dy < 0 {
		dy = -dy
	}

	return dx + dy
}

func FindClosestCell(cells []Position, targetCell Position) (Position, bool) {
	if len(cells) == 0 {
		return Position{}, false
	}

	closest := cells[0]
	minDist := GetDistance(closest, targetCell)

	for _, cell := range cells[1:] {
		dist := GetDistance(cell, targetCell)
		if dist < minDist {
			minDist = dist
			closest = cell
		}
	}

	return closest, true
}
