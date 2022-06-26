package common

type Direction struct {
	X int
	Y int
}

type Position struct {
	X float64
	Y float64
}

func CopyPosition(from *Position) *Position {
	return &Position{
		X: from.X,
		Y: from.Y,
	}
}
func PositionFromInt(x int, y int) *Position {
	return &Position{
		X: float64(x),
		Y: float64(y),
	}
}

func PositionFromDirection(direction *Direction) *Position {
	return &Position{
		X: float64(direction.X),
		Y: float64(direction.Y),
	}
}

func (v *Position) Mul(val float64) *Position {
	return &Position{
		X: v.X * val,
		Y: v.Y * val,
	}
}

func (v *Position) Add(other *Position) *Position {
	return &Position{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func (v *Position) Sub(other *Position) *Position {
	return &Position{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}
