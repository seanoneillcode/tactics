package common

const (
	ScreenWidth      = 384
	ScreenHeight     = 240
	TileSize         = 16
	TileSizeF        = 16.0
	Scale            = 4
	ScaleF           = 4.0
	HalfScreenWidth  = ScreenWidth / 2
	HalfScreenHeight = ScreenHeight / 2
	HalfTileSize     = TileSize / 2
)

func WorldToTile(pos *Position) (int, int) {
	return int(pos.X / TileSize), int(pos.Y / TileSize)
}

func WorldToTileInt(x int, y int) (int, int) {
	return x / TileSize, y / TileSize
}
