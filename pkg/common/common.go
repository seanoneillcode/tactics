package common

const (
	ScreenWidth      = 320
	ScreenHeight     = 180
	TileSize         = 16
	TileSizeF        = 16.0
	Scale            = 4
	ScaleF           = 4.0
	HalfScreenWidth  = ScreenWidth / 2
	HalfScreenHeight = ScreenHeight / 2
	HalfTileSize     = TileSize / 2
)

func WorldToTile(pos *Position) Tile {
	return Tile{
		X: int(pos.X / TileSize),
		Y: int(pos.Y / TileSize),
	}
}

func WorldToTileInt(x int, y int) Tile {
	return Tile{
		X: x / TileSize,
		Y: y / TileSize,
	}
}
