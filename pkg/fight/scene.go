package fight

import (
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"math/rand"
)

type Scene struct {
	Name      string
	tiledGrid *common.TiledGrid
}

func NewScene(name string, playerActors []*Actor, enemyActors []*Actor) *Scene {
	tiledGrid := common.NewTileGrid(name + ".json")
	objects := tiledGrid.GetObjectData()
	teamOnePos, teamTwoPos := loadStartPositions(objects)

	numberOnePositions := len(teamOnePos)
	offset := rand.Intn(numberOnePositions)
	for i, actor := range playerActors {
		index := (i + offset) % numberOnePositions
		actor.Pos = &common.Position{
			X: teamOnePos[index].X,
			Y: teamOnePos[index].Y,
		}
	}

	numberTwoPositions := len(teamTwoPos)
	offset = rand.Intn(numberTwoPositions)
	for i, actor := range enemyActors {
		index := (i + offset) % numberTwoPositions
		actor.Pos = &common.Position{
			X: teamTwoPos[index].X,
			Y: teamTwoPos[index].Y,
		}
	}

	s := &Scene{
		Name:      name,
		tiledGrid: tiledGrid,
	}

	return s
}

func (r *Scene) Update(delta int64, state *State) {
	// update objects
}

func (r *Scene) Draw(camera *Camera) {
	r.tiledGrid.Draw(camera)
	// draw objects
}

func loadStartPositions(objects []*common.ObjectData) ([]*common.Position, []*common.Position) {
	var teamOnePositions []*common.Position
	var teamTwoPositions []*common.Position
	for _, obj := range objects {
		if obj.ObjectType == "start" {
			npc := &common.Position{
				X: float64(obj.X),
				Y: float64(obj.Y - common.TileSize),
			}
			for _, p := range obj.Properties {
				if p.Name == "team" {
					switch p.Value {
					case "1":
						teamOnePositions = append(teamOnePositions, npc)
					case "2":
						teamTwoPositions = append(teamTwoPositions, npc)
					}
				}
			}
		}
	}
	return teamOnePositions, teamTwoPositions
}
