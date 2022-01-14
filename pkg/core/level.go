package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
)

type Level struct {
	npcs      []*Npc
	links     []*Link
	tiledGrid *TiledGrid
	// pickups, dialogues, enemies ...
}

func NewLevel(fileName string) *Level {
	tiledGrid := NewTileGrid(fileName)
	objects := tiledGrid.GetObjectData()

	return &Level{
		npcs:      loadNpcs(objects),
		links:     loadLinks(objects),
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
	}
	for _, npc := range l.npcs {
		nx, ny := common.WorldToTile(npc.GetPosition())
		if nx == x && ny == y {
			ti.npc = npc
			break
		}
	}
	for _, link := range l.links {
		nx, ny := common.WorldToTileInt(link.GetPosition())
		if nx == x && ny == y {
			ti.link = link
			break
		}
	}
	return ti
}

// TileInfo is a read only struct of references to handle tile based queries
type TileInfo struct {
	tileData *TileData
	npc      *Npc
	link     *Link
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

func loadLinks(objects []*ObjectData) []*Link {
	var links []*Link
	for _, obj := range objects {
		if obj.objectType == "link" {
			link := &Link{
				x:    obj.x,
				y:    obj.y,
				name: obj.name,
			}
			for _, p := range obj.properties {
				if p.name == "direction" {
					link.direction = (p.value).(string)
				}
				if p.name == "to-level" {
					link.toLevel = (p.value).(string)
				}
			}
			links = append(links, link)
		}
	}
	return links
}
