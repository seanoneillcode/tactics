package common

type TileVector struct {
	X int
	Y int
}

type VectorF struct {
	X float64
	Y float64
}

func VectorFromInt(x int, y int) *VectorF {
	return &VectorF{
		X: float64(x),
		Y: float64(y),
	}
}

func (v *VectorF) Mul(val float64) *VectorF {
	return &VectorF{
		X: v.X * val,
		Y: v.Y * val,
	}
}

func (v *VectorF) Add(other *VectorF) *VectorF {
	return &VectorF{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func (v *VectorF) Sub(other *VectorF) *VectorF {
	return &VectorF{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}
