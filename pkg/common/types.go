package common

type Direction struct {
	X int
	Y int
}

type Vector struct {
	X float64
	Y float64
}

func VectorFromInt(x int, y int) *Vector {
	return &Vector{
		X: float64(x),
		Y: float64(y),
	}
}

func (v *Vector) Mul(val float64) *Vector {
	return &Vector{
		X: v.X * val,
		Y: v.Y * val,
	}
}

func (v *Vector) Add(other *Vector) *Vector {
	return &Vector{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func (v *Vector) Sub(other *Vector) *Vector {
	return &Vector{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}
