package fight

import "github.com/seanoneillcode/go-tactics/pkg/common"

type Pattern interface {
	GetPattern(pos common.Tile, state *State) []common.Tile
}

type CrossPattern struct{}

func (r *CrossPattern) GetPattern(pos common.Tile, state *State) []common.Tile {
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

type LinePattern struct{}

func (r *LinePattern) GetPattern(pos common.Tile, state *State) []common.Tile {
	var tiles []common.Tile

	tiles = append(tiles, getLineTiles(pos, common.Tile{
		X: 1,
		Y: 0,
	}, state)...)
	tiles = append(tiles, getLineTiles(pos, common.Tile{
		X: -1,
		Y: 0,
	}, state)...)
	tiles = append(tiles, getLineTiles(pos, common.Tile{
		X: 0,
		Y: 1,
	}, state)...)
	tiles = append(tiles, getLineTiles(pos, common.Tile{
		X: 0,
		Y: -1,
	}, state)...)

	return tiles
}

func getLineTiles(startPos common.Tile, dir common.Tile, state *State) []common.Tile {
	var tiles = []common.Tile{}
	indexPos := common.Tile{
		X: startPos.X,
		Y: startPos.Y,
	}
	for indexPos.X < 100 && indexPos.X >= 0 && indexPos.Y < 100 && indexPos.Y >= 0 {
		indexPos.X += dir.X
		indexPos.Y += dir.Y
		data := state.Scene.tiledGrid.GetTileData(indexPos.X, indexPos.Y)
		if data.IsBlock {
			break
		}
		tiles = append(tiles, common.Tile{
			X: indexPos.X,
			Y: indexPos.Y,
		})
		actor := GetActorAtPos(state, indexPos)
		if actor != nil {
			// is blocked!
			break
		}
	}
	return tiles
}
