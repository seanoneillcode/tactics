package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"github.com/seanoneillcode/go-tactics/pkg/core"
	"log"
	"time"
)

type Game struct {
	keys  []ebiten.Key
	state *core.State

	lastUpdateCalled time.Time
}

func (g *Game) Update() error {
	delta := time.Now().Sub(g.lastUpdateCalled).Milliseconds()

	g.state.Level.Update(delta)
	g.state.Player.Update(delta, g.state)

	g.keys = inpututil.AppendPressedKeys(g.keys[:0])

	g.lastUpdateCalled = time.Now()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	g.state.Level.Draw(screen)
	g.state.Player.Draw(screen)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return common.ScreenWidth * common.Scale, common.ScreenHeight * common.Scale
}

func main() {
	g := &Game{
		lastUpdateCalled: time.Now(),
		state: &core.State{
			Player: core.NewCharacter(),
			Level:  core.NewTileGrid(),
		},
	}

	ebiten.SetWindowSize(common.ScreenWidth*common.Scale, common.ScreenHeight*common.Scale)
	ebiten.SetWindowTitle("Fantasy Game")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
