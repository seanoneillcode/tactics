package common

const (
	ScreenWidth      = 256
	ScreenHeight     = 240
	TileSize         = 16
	TileSizeF        = 16.0
	Scale            = 4
	ScaleF           = 4.0
	HalfScreenWidth  = ScreenWidth / 2
	HalfScreenHeight = ScreenHeight / 2
	HalfTileSize     = TileSize / 2
)

func WorldToTile(pos *VectorF) (int, int) {
	return int(pos.X / TileSize), int(pos.Y / TileSize)
}

func WorldToTileInt(x int, y int) (int, int) {
	return x / TileSize, y / TileSize
}

func TileToWorld(vector Vector) *VectorF {
	return &VectorF{
		X: float64(TileSizeF * vector.X),
		Y: float64(TileSizeF * vector.Y),
	}
}
