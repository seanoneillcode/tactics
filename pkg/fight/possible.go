package fight

import "github.com/seanoneillcode/go-tactics/pkg/common"

type PossibleMoves struct {
	moves    []common.Tile
	moveTile *common.Sprite
}

func NewPossibleMoves() *PossibleMoves {
	return &PossibleMoves{
		moveTile: common.NewSprite("move-tile-selection.png"),
	}
}

func (r *PossibleMoves) Draw(camera *Camera) {
	for _, m := range r.moves {
		r.moveTile.X = float64(m.X * common.TileSize)
		r.moveTile.Y = float64(m.Y * common.TileSize)
		r.moveTile.Draw(camera)
	}
}

func (r *PossibleMoves) GeneratePossibleMoves(state *State, numMoves int) {
	r.moves = []common.Tile{}
	pos := state.PlayerController.SelectedActor.Pos
	r.generateTileMoves(state, int(pos.X/common.TileSize), int(pos.Y/common.TileSize), numMoves)
}

func (r *PossibleMoves) generateTileMoves(state *State, x int, y int, numMoves int) {
	if numMoves < 0 {
		return
	}

	if x < 0 || x > state.Scene.tiledGrid.Layers[0].Width {
		return // outside boundary
	}
	if y < 0 || y > state.Scene.tiledGrid.Layers[0].Height {
		return // outside boundary
	}

	td := state.Scene.tiledGrid.GetTileData(x, y)

	if td.IsBlock {
		return // blocked
	}

	// it's valid to move here
	r.moves = append(r.moves, common.Tile{X: x, Y: y})
	numMoves = numMoves - 1

	r.generateTileMoves(state, x-1, y, numMoves)
	r.generateTileMoves(state, x+1, y, numMoves)
	r.generateTileMoves(state, x, y-1, numMoves)
	r.generateTileMoves(state, x, y+1, numMoves)
}

func (r *PossibleMoves) CanMove(x float64, y float64) bool {
	ix := int(x / common.TileSize)
	iy := int(y / common.TileSize)
	validPlace := false
	for _, m := range r.moves {
		if m.X == ix && m.Y == iy {
			validPlace = true
			break
		}
	}
	if !validPlace {
		return false
	}

	return true
}

type PossibleTargets struct {
	targets   []common.Tile
	selection *common.Sprite
}

func NewPossibleTargets() *PossibleTargets {
	return &PossibleTargets{
		selection: common.NewSprite("target-tile-selection.png"),
	}
}

func (r *PossibleTargets) CanTarget(f float64, y float64) bool {
	return true
}

func (r *PossibleTargets) GeneratePossibleTargets(state *State) {
	pos := state.PlayerController.SelectedActor.Pos
	px, py := int(pos.X/common.TileSize), int(pos.Y/common.TileSize)

	skill := state.PlayerController.SelectedActor.Skills[0]
	r.targets = skill.TargetPattern.GetPattern(common.Tile{X: px, Y: py}, state)

}

func (r *PossibleTargets) Draw(camera *Camera) {
	for _, m := range r.targets {
		r.selection.X = float64(m.X * common.TileSize)
		r.selection.Y = float64(m.Y * common.TileSize)
		r.selection.Draw(camera)
	}
}
