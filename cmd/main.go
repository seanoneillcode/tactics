package main

import (
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"github.com/seanoneillcode/go-tactics/pkg/core"
	"github.com/seanoneillcode/go-tactics/pkg/gui"
	"log"
	"time"
)

var NormalEscapeError = errors.New("normal escape termination")

type Game struct {
	keys             []ebiten.Key
	state            *core.State
	lastUpdateCalled time.Time
	dialogBox        *gui.DialogueBox
	camera           *core.Camera
}

func (g *Game) Update() error {
	// calculate delta
	delta := time.Now().Sub(g.lastUpdateCalled).Milliseconds()
	g.lastUpdateCalled = time.Now()

	// update state
	g.state.Map.Update(delta, g.state)
	g.state.Player.Update(delta, g.state)
	if g.state.ActiveDialog != nil {
		g.state.ActiveDialog.Update(delta, g.state)
	}
	g.state.Shop.Update(delta, g.state)

	// update camera
	g.camera.Update(delta, g.state)

	// update UI
	g.dialogBox.Update(delta, g.state)

	// handle escape
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return NormalEscapeError
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.state.Map.Level.Draw(g.camera.GetBuffer())
	g.state.Player.Draw(g.camera.GetBuffer())
	g.camera.DrawBuffer(screen)
	g.state.Shop.Draw(screen)
	g.dialogBox.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return common.ScreenWidth * common.Scale, common.ScreenHeight * common.Scale
}

func main() {
	g := &Game{
		lastUpdateCalled: time.Now(),
		state: &core.State{
			Player: core.NewPlayer(),
			Map:    core.NewMap(),
			Shop:   core.NewShop(),
		},
		dialogBox: gui.NewDialogueBox(),
		camera:    core.NewCamera(),
	}
	g.state.Map.LoadLevel("siopa")
	g.state.Player.EnterLevel(g.state.Map.Level)

	ebiten.SetWindowSize(common.ScreenWidth*common.Scale, common.ScreenHeight*common.Scale)
	ebiten.SetWindowTitle("Fantasy Game")
	err := ebiten.RunGame(g)
	if err != nil {
		if errors.Is(err, NormalEscapeError) {
			log.Println("exiting normally")
		} else {
			log.Fatal(err)
		}
	}
}
