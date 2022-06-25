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
	tiledGrid *common.TiledGrid
	enemies   []*Enemy
	// enemies ...
}

func NewLevel(name string) *Level {
	tiledGrid := common.NewTileGrid(name + ".json")
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
	tileData *common.TileData
	npc      *Npc
	link     *Link
	pickup   *Pickup
	action   *Action
	shop     *ShopData
	enemy    *Enemy

	// etc
}

func loadNpcs(objects []*common.ObjectData) []*Npc {
	var npcs []*Npc
	for _, obj := range objects {
		if obj.ObjectType == "npc" {
			npc := NewNpc(obj.Name)
			npc.SetPosition(obj.X, obj.Y-common.TileSize)
			npcs = append(npcs, npc)
		}
	}
	return npcs
}

func loadEnemy(objects []*common.ObjectData) []*Enemy {
	var enemies []*Enemy
	for _, obj := range objects {
		if obj.ObjectType == "enemy" {
			enemy := NewEnemy(obj.Name)
			enemy.SetPosition(obj.X, obj.Y-common.TileSize)
			enemies = append(enemies, enemy)
		}
	}
	return enemies
}

func loadPickups(objects []*common.ObjectData) []*Pickup {
	var pickups []*Pickup
	for _, obj := range objects {
		if obj.ObjectType == "pickup" {
			var itemName string
			var usedImageName string
			for _, p := range obj.Properties {
				if p.Name == "item" {
					itemName = (p.Value).(string)
				}
				if p.Name == "used-image" {
					usedImageName = (p.Value).(string)
				}
			}

			pickup := NewPickup(obj.Name, itemName, usedImageName)
			pickup.SetPosition(obj.X, obj.Y-common.TileSize)
			pickups = append(pickups, pickup)
		}
	}
	return pickups
}

func loadActions(objects []*common.ObjectData) []*Action {
	var actions []*Action
	for _, obj := range objects {
		if obj.ObjectType == "action" {
			action := NewAction(obj.Name, float64(obj.X), float64(obj.Y))
			actions = append(actions, action)
		}
	}
	return actions
}

func loadShops(objects []*common.ObjectData) []*ShopData {
	var shops []*ShopData
	for _, obj := range objects {
		if obj.ObjectType == "shop" {
			s := NewShopData(obj.Name, float64(obj.X), float64(obj.Y))
			for _, p := range obj.Properties {
				if p.Name == "merchantName" {
					s.MerchantName = (p.Value).(string)
				}
			}
			shops = append(shops, s)
		}
	}
	return shops
}

func loadLinks(objects []*common.ObjectData) []*Link {
	var links []*Link
	for _, obj := range objects {
		if obj.ObjectType == "link" {
			link := &Link{
				x:    obj.X,
				y:    obj.Y,
				name: obj.Name,
			}
			for _, p := range obj.Properties {
				if p.Name == "direction" {
					link.direction = (p.Value).(string)
				}
				if p.Name == "to-level" {
					link.toLevel = (p.Value).(string)
				}
			}
			links = append(links, link)
		}
	}
	return links
}
