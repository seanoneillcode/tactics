package fight

import (
	"github.com/beefsack/go-astar"
	"github.com/seanoneillcode/go-tactics/pkg/common"
)

const AiTurnTime = 500
const numberEnemyMoveTiles = 3

type AiController struct {
	CurrentActor *Actor
	PathGrid     common.PathGrid

	StepTimer     int64
	StepPositions []*common.Position
	StepIndex     int
}

func (r *AiController) StartTurn(state *State) {
	r.CurrentActor = state.AiTeam.Actors[0]
	state.ActiveTeam = state.AiTeam
	state.AiTeam.StartTurn()
	r.PathGrid = common.GeneratePathGrid(state.Scene.tiledGrid)
	r.StepPositions = nil
	state.Camera.Target(r.CurrentActor)
}

func (r *AiController) Update(delta int64, state *State) {
	if state.ActiveTeam != state.AiTeam {
		return
	}
	if r.StepPositions == nil {
		r.StepPositions = r.GenerateStepPositions(state, r.CurrentActor, numberEnemyMoveTiles)
		r.StepIndex = 0
		r.StepTimer = 0
	} else {
		r.StepTimer = r.StepTimer - delta
		if r.StepTimer < 0 {
			r.StepTimer = AiTurnTime
			r.CurrentActor.SetPos(r.StepPositions[r.StepIndex])
			r.StepIndex += 1
			if r.StepIndex == len(r.StepPositions) {
				r.CurrentActor.ActionTokensLeft = 0
				nextActor := state.AiTeam.GetNextActor(r.CurrentActor)
				if nextActor.ActionTokensLeft == 0 {
					state.PlayerController.StartTurn(state)
				} else {
					r.CurrentActor = nextActor
					state.Camera.Target(r.CurrentActor)
				}
				r.StepPositions = nil
			}
		}
	}

}

func (r *AiController) GenerateStepPositions(state *State, actor *Actor, tilesToMove int) []*common.Position {
	target := state.PlayerTeam.Actors[0]
	pos0 := common.WorldToTile(actor.Pos)
	pos1 := common.WorldToTile(target.Pos)

	from := r.PathGrid.GetMapTile(pos0.X, pos0.Y)
	to := r.PathGrid.GetMapTile(pos1.X, pos1.Y)
	p, _, _ := astar.Path(from, to)

	tileIndex := len(p) - tilesToMove
	if tileIndex < 0 {
		tileIndex = 0
	}

	var turnPositions []*common.Position
	index := len(p) - 1
	for index >= tileIndex {
		targetPathTile := p[index].(*common.PathTile)
		turnPositions = append(turnPositions, common.PositionFromInt(targetPathTile.X*common.TileSize, targetPathTile.Y*common.TileSize))
		index -= 1
	}
	return turnPositions
}
