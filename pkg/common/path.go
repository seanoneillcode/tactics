package common

import (
	"github.com/beefsack/go-astar"
)

type PathGrid map[int]map[int]*PathTile

func (w PathGrid) GetMapTile(x, y int) *PathTile {
	if w[x] == nil {
		return nil
	}
	return w[x][y]
}

func (w PathGrid) SetTile(t *PathTile, x, y int) {
	if w[x] == nil {
		w[x] = map[int]*PathTile{}
	}
	w[x][y] = t
	t.X = x
	t.Y = y
	t.W = w
}

type TileKind int

const (
	NormalTileKind  TileKind = 0
	BlockerTileKind TileKind = 1
)

type PathTile struct {
	TileKind TileKind
	X        int
	Y        int
	W        PathGrid
}

func (t *PathTile) PathNeighbors() []astar.Pather {
	neighbors := []astar.Pather{}
	for _, offset := range [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	} {
		n := t.W.GetMapTile(t.X+offset[0], t.Y+offset[1])
		if n != nil && n.TileKind != BlockerTileKind {
			neighbors = append(neighbors, n)
		}
	}
	return neighbors
}

// PathNeighborCost returns the movement cost of the directly neighboring tile.
func (t *PathTile) PathNeighborCost(to astar.Pather) float64 {
	return 1
}

// PathEstimatedCost uses Manhattan distance to estimate orthogonal distance
// between non-adjacent nodes.
func (t *PathTile) PathEstimatedCost(to astar.Pather) float64 {
	toT := to.(*PathTile)
	absX := toT.X - t.X
	if absX < 0 {
		absX = -absX
	}
	absY := toT.Y - t.Y
	if absY < 0 {
		absY = -absY
	}
	return float64(absX + absY)
}
