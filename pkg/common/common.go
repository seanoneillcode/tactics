package common

const (
	ScreenWidth      = 256
	ScreenHeight     = 240
	TileSize         = 16
	Scale            = 4
	HalfScreenWidth  = ScreenWidth / 2
	HalfScreenHeight = ScreenHeight / 2
	HalfTileSize     = TileSize / 2
)

func WorldToTile(x float64, y float64) (int, int) {
	return int(x / TileSize), int(y / TileSize)
}

func WorldToTileInt(x int, y int) (int, int) {
	return x / TileSize, y / TileSize
}

func TileToWorld(x int, y int) (float64, float64) {
	return float64(x * TileSize), float64(y * TileSize)
}
