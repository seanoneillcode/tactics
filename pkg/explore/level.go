package explore

import (
	"github.com/seanoneillcode/go-tactics/pkg/common"
)

type Level struct {
	Name      string
	npcs      []*Npc
	links     []*Link
	pickups   []*Pickup
	actions   []*Action
	shops     []*ShopData
	tiledGrid *TiledGrid
	enemies   []*Enemy
	// enemies ...
}

func NewLevel(name string) *Level {
	tiledGrid := NewTileGrid(name + ".json")
	objects := tiledGrid.GetObjectData()

	return &Level{
		Name:      name,
		npcs:      loadNpcs(objects),
		links:     loadLinks(objects),
		pickups:   loadPickups(objects),
		actions:   loadActions(objects),
		shops:     loadShops(objects),
		enemies:   loadEnemy(objects),
		tiledGrid: tiledGrid,
	}
}

func (l *Level) Update(delta int64, state *State) {
	for _, npc := range l.npcs {
		npc.Update(delta, state)
	}
	for _, enemy := range l.enemies {
		enemy.Update(delta, state)
	}
}

func (l *Level) Draw(camera *Camera) {
	l.tiledGrid.Draw(camera)
	for _, npc := range l.npcs {
		npc.Draw(camera)
	}
	for _, enemy := range l.enemies {
		enemy.Draw(camera)
	}
	for _, pickup := range l.pickups {
		pickup.Draw(camera)
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
	for _, enemy := range l.enemies {
		nx, ny := common.WorldToTile(enemy.GetPosition())
		if nx == x && ny == y {
			ti.enemy = enemy
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
	for _, pickup := range l.pickups {
		nx, ny := common.WorldToTile(pickup.GetPosition())
		if nx == x && ny == y {
			ti.pickup = pickup
			break
		}
	}
	for _, action := range l.actions {
		nx, ny := common.WorldToTile(action.GetPosition())
		if nx == x && ny == y {
			ti.action = action
			break
		}
	}
	for _, s := range l.shops {
		nx, ny := common.WorldToTile(s.GetPosition())
		if nx == x && ny == y {
			ti.shop = s
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
	pickup   *Pickup
	action   *Action
	shop     *ShopData
	enemy    *Enemy

	// etc
}

func loadNpcs(objects []*ObjectData) []*Npc {
	var npcs []*Npc
	for _, obj := range objects {
		if obj.objectType == "npc" {
			npc := NewNpc(obj.name)
			npc.SetPosition(obj.x, obj.y-common.TileSize)
			npcs = append(npcs, npc)
		}
	}
	return npcs
}

func loadEnemy(objects []*ObjectData) []*Enemy {
	var enemies []*Enemy
	for _, obj := range objects {
		if obj.objectType == "enemy" {
			enemy := NewEnemy(obj.name)
			enemy.SetPosition(obj.x, obj.y-common.TileSize)
			enemies = append(enemies, enemy)
		}
	}
	return enemies
}

func loadPickups(objects []*ObjectData) []*Pickup {
	var pickups []*Pickup
	for _, obj := range objects {
		if obj.objectType == "pickup" {
			var itemName string
			var usedImageName string
			for _, p := range obj.properties {
				if p.name == "item" {
					itemName = (p.value).(string)
				}
				if p.name == "used-image" {
					usedImageName = (p.value).(string)
				}
			}

			pickup := NewPickup(obj.name, itemName, usedImageName)
			pickup.SetPosition(obj.x, obj.y-common.TileSize)
			pickups = append(pickups, pickup)
		}
	}
	return pickups
}

func loadActions(objects []*ObjectData) []*Action {
	var actions []*Action
	for _, obj := range objects {
		if obj.objectType == "action" {
			action := NewAction(obj.name, float64(obj.x), float64(obj.y))
			actions = append(actions, action)
		}
	}
	return actions
}

func loadShops(objects []*ObjectData) []*ShopData {
	var shops []*ShopData
	for _, obj := range objects {
		if obj.objectType == "shop" {
			s := NewShopData(obj.name, float64(obj.x), float64(obj.y))
			for _, p := range obj.properties {
				if p.name == "merchantName" {
					s.MerchantName = (p.value).(string)
				}
			}
			shops = append(shops, s)
		}
	}
	return shops
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
