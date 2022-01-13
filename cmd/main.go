package main

import (
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"github.com/seanoneillcode/go-tactics/pkg/core"
	"log"
	"time"
)

var NormalEscapeError = errors.New("normal escape termination")

type Game struct {
	keys             []ebiten.Key
	state            *core.State
	lastUpdateCalled time.Time
}

func (g *Game) Update() error {
	// calculate delta
	delta := time.Now().Sub(g.lastUpdateCalled).Milliseconds()
	g.lastUpdateCalled = time.Now()

	// update state
	g.state.Level.Update(delta, g.state)
	g.state.Player.Update(delta, g.state)

	// handle escape
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return NormalEscapeError
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.state.Level.Draw(screen)
	g.state.Player.Draw(screen)
	g.state.DialogueBox.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return common.ScreenWidth * common.Scale, common.ScreenHeight * common.Scale
}

func main() {
	g := &Game{
		lastUpdateCalled: time.Now(),
		state: &core.State{
			Player:      core.NewPlayer(),
			Level:       core.NewLevel("home.json"),
			DialogueBox: core.NewDialogueBox(),
		},
	}
	//g.state.DialogueBox.AddTextBox("The quick brown fox jumps over the moon, however I don't jump at all. Testing this to completion. Making sure there are no breaks in the line and they don't overrun.")

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
