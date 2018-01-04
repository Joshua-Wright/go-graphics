package dragon_turtle

import (
	m "github.com/joshua-wright/image_stuff/go"
)

type Symbol int

const (
	X       Symbol = iota
	Y
	FORWARD
	LEFT     // -
	RIGHT    // +
)

type Direction int

const (
	EAST  Direction = iota
	NORTH
	WEST
	SOUTH
)

func (d Direction) left() Direction {
	return (d + 1) % 4
}
func (d Direction) right() Direction {
	return (d + 3) % 4
}

func dragonCurveSingleIteration(instructions []Symbol) []Symbol {
	out := []Symbol{}
	for _, move := range instructions {
		switch move {
		case X:
			out = append(out, X, RIGHT, Y, FORWARD, RIGHT)
		case Y:
			out = append(out, LEFT, FORWARD, X, LEFT, Y)
		case FORWARD, LEFT, RIGHT:
			out = append(out, move)
		}
	}
	return out
}

func DragonCurveSymbols(depth int) []Symbol {
	moves := []Symbol{FORWARD, X}
	for i := 0; i < depth; i++ {
		moves = dragonCurveSingleIteration(moves)
	}
	return moves
}

// axes: east, north, west, south
func DragonCurve(moves []Symbol, start m.Vec2, dir Direction, axes [4]m.Vec2, gridSize m.Float) []m.Vec2 {
	out := []m.Vec2{}
	current := start
	for _, move := range moves {
		switch move {
		case X, Y:

		case LEFT:
			dir = dir.left()

		case RIGHT:
			dir = dir.right()

		case FORWARD:
			out = append(out, current)
			current = current.AddV(axes[dir].MulS(gridSize))
			//switch dir {
			//case NORTH:
			//	current.Y += gridSize
			//case SOUTH:
			//	current.Y -= gridSize
			//case EAST:
			//	current.X += gridSize
			//case WEST:
			//	current.X -= gridSize
			//}
		}
	}
	return out
}

// east, north, west, south
var DefaultAxes = [4]m.Vec2{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

var DiagonalAxes = [4]m.Vec2{
	{1, 1},
	{-1, 1},
	{-1, -1},
	{1, -1},
}
