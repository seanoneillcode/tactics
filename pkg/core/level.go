package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
)

type Level struct {
	npcs      []*Npc
	tiledGrid *TiledGrid
	// pickups, dialogues, enemies ...
}

func NewLevel(fileName string) *Level {
	tiledGrid := NewTileGrid(fileName)
	objects := tiledGrid.GetObjectData()

	return &Level{
		npcs:      loadNpcs(objects),
		tiledGrid: tiledGrid,
	}
}

func (l *Level) Update(delta int64, state *State) {
	for _, npc := range l.npcs {
		npc.Update(delta, state)
	}
}

func (l *Level) Draw(screen *ebiten.Image) {
	l.tiledGrid.Draw(screen)
	for _, npc := range l.npcs {
		npc.Draw(screen)
	}
}

func (l *Level) GetTileInfo(x int, y int) *TileInfo {
	ti := &TileInfo{
		tileData: l.tiledGrid.GetTileData(x, y),
		npcs:     []*Npc{},
	}
	// todo this could be a performance bottleneck, consider making a level state and updating it
	for _, npc := range l.npcs {
		nx, ny := common.WorldToTile(npc.GetPosition())
		if nx == x && ny == y {
			ti.npcs = append(ti.npcs, npc)
		}
	}
	return ti
}

// TileInfo is a read only struct of references to handle tile based queries
type TileInfo struct {
	tileData *TileData
	npcs     []*Npc
	// etc
}

func loadNpcs(objects []*ObjectData) []*Npc {
	var npcs []*Npc

	for _, obj := range objects {
		if obj.objectType == "npc" {
			npc := NewNpc(obj.name)
			npc.SetPosition(obj.x, obj.y)
			npcs = append(npcs, npc)
		}
	}

	return npcs
}
