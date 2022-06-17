package main

import (
	"errors"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/seanoneillcode/go-tactics/pkg/gui/dialog"
	"github.com/seanoneillcode/go-tactics/pkg/gui/equipment"
	"github.com/seanoneillcode/go-tactics/pkg/gui/inventory"
	"github.com/seanoneillcode/go-tactics/pkg/gui/menu"
	"github.com/seanoneillcode/go-tactics/pkg/input"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"github.com/seanoneillcode/go-tactics/pkg/core"
	"github.com/seanoneillcode/go-tactics/pkg/gui"
)

func main() {
	g := &Game{
		lastUpdateCalled: time.Now(),
		state: &core.State{
			Control:          &core.Control{},
			Player:           core.NewPlayer(),
			Map:              core.NewMap(),
			Shop:             core.NewShop(),
			UI:               core.NewUI(), // todo replace with concept of active element or 'mode' and a unified controller for currently active element
			TeamState:        core.NewTeamState(),
			TotalElapsedTime: 12 * 1000 * 60,
			DialogHandler:    core.NewDialogHandler(),
		},
		dialog:      dialog.NewUi(),
		shopUI:      gui.NewShopUi(),
		camera:      core.NewCamera(),
		inventoryUI: inventory.NewUi(),
		menuUI:      menu.NewUI(),
		equipmentUI: equipment.NewUI(),
	}
	g.state.Map.LoadLevel("pub-level")
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

var NormalEscapeError = errors.New("normal escape termination")

type Game struct {
	lastUpdateCalled time.Time
	keys             []ebiten.Key
	state            *core.State
	camera           *core.Camera
	dialog           gui.UI
	shopUI           *gui.ShopUI
	inventoryUI      gui.UI
	menuUI           gui.UI
	equipmentUI      gui.UI
}

func (g *Game) Update() error {
	// calculate delta
	delta := time.Now().Sub(g.lastUpdateCalled).Milliseconds()
	g.lastUpdateCalled = time.Now()
	g.state.TotalElapsedTime = g.state.TotalElapsedTime + delta

	// update state
	g.state.Map.Update(delta, g.state)
	g.state.Player.Update(delta, g.state)
	g.state.Shop.Update(delta, g.state)

	// update camera
	g.camera.Update(delta, g.state)

	// update UI
	g.dialog.Update(delta, g.state)
	g.shopUI.Update(delta, g.state)
	g.inventoryUI.Update(delta, g.state)
	g.menuUI.Update(delta, g.state)
	g.equipmentUI.Update(delta, g.state)

	// handle escape
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || g.state.Control.Command == "exit" {
		return NormalEscapeError
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}
	input.Update()
	//fps := ebiten.CurrentFPS()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.state.Map.Level.Draw(g.camera.GetBuffer())
	g.state.Player.Draw(g.camera.GetBuffer())
	g.camera.DrawBuffer(screen)
	g.dialog.Draw(screen)
	g.shopUI.Draw(screen)
	g.inventoryUI.Draw(screen)
	g.menuUI.Draw(screen)
	g.equipmentUI.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return common.ScreenWidth * common.Scale, common.ScreenHeight * common.Scale
}
