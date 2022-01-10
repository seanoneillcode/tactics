package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"github.com/seanoneillcode/go-tactics/pkg/core"
	"log"
	"strings"
	"time"
)

type Game struct {
	keys             []ebiten.Key
	level            *core.TiledGrid
	player           *core.Character
	lastUpdateCalled time.Time
}

func (g *Game) Update() error {
	delta := time.Now().Sub(g.lastUpdateCalled).Milliseconds()

	g.level.Update(delta)
	g.player.Update(delta)

	g.keys = inpututil.AppendPressedKeys(g.keys[:0])

	g.lastUpdateCalled = time.Now()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	g.level.Draw(screen)
	g.player.Draw(screen)

	var keyStrings []string
	for _, p := range g.keys {
		keyStrings = append(keyStrings, p.String())
	}
	ebitenutil.DebugPrint(screen, strings.Join(keyStrings, ", "))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return common.ScreenWidth * common.Scale, common.ScreenHeight * common.Scale
}

func main() {
	g := &Game{
		lastUpdateCalled: time.Now(),
		player:           core.NewCharacter(),
		level:            core.NewTileGrid(),
	}

	ebiten.SetWindowSize(common.ScreenWidth*common.Scale, common.ScreenHeight*common.Scale)
	ebiten.SetWindowTitle("Fantasy Game")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
