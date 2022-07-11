package fight

import "github.com/seanoneillcode/go-tactics/pkg/common"

type Pattern interface {
	GetPattern(pos common.Tile, state *State) []common.Tile
}

type BasicPattern struct{}

func (r *BasicPattern) GetPattern(pos common.Tile, state *State) []common.Tile {
	return []common.Tile{
		{
			X: pos.X + 1,
			Y: pos.Y,
		},
		{
			X: pos.X - 1,
			Y: pos.Y,
		},
		{
			X: pos.X,
			Y: pos.Y + 1,
		},
		{
			X: pos.X,
			Y: pos.Y - 1,
		},
	}
}

type SinglePattern struct{}

func (r *SinglePattern) GetPattern(pos common.Tile, state *State) []common.Tile {
	return []common.Tile{
		{
			X: pos.X,
			Y: pos.Y,
		},
	}
}

type BallPattern struct{}

func (r *BallPattern) GetPattern(pos common.Tile, state *State) []common.Tile {
	var tiles []common.Tile
	i := 0
	for i < 9 {
		x := i % 3
		y := i / 3
		tiles = append(tiles, common.Tile{
			X: pos.X - 1 + x,
			Y: pos.Y - 1 + y,
		})
		i += 1
	}
	return tiles
}

type CrossDistancePattern struct {
	Distance int
}

func (r *CrossDistancePattern) GetPattern(pos common.Tile, state *State) []common.Tile {
	return []common.Tile{
		{
			X: pos.X + r.Distance,
			Y: pos.Y,
		},
		{
			X: pos.X - r.Distance,
			Y: pos.Y,
		},
		{
			X: pos.X,
			Y: pos.Y + r.Distance,
		},
		{
			X: pos.X,
			Y: pos.Y - r.Distance,
		},
	}
}
