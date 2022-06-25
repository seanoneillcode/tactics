package main

import (
	"errors"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/seanoneillcode/go-tactics/pkg/fight"
	"github.com/seanoneillcode/go-tactics/pkg/gui/dialog"
	"github.com/seanoneillcode/go-tactics/pkg/gui/equipment"
	"github.com/seanoneillcode/go-tactics/pkg/gui/inventory"
	"github.com/seanoneillcode/go-tactics/pkg/gui/menu"
	"github.com/seanoneillcode/go-tactics/pkg/input"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"github.com/seanoneillcode/go-tactics/pkg/explore"
	"github.com/seanoneillcode/go-tactics/pkg/gui"
)

func main() {
	g := &Game{
		lastUpdateCalled: time.Now(),
		state: &explore.State{
			Control:          &explore.Control{},
			Player:           explore.NewPlayer(),
			Map:              explore.NewMap(),
			Shop:             explore.NewShop(),
			UI:               explore.NewUI(), // todo replace with concept of active element or 'mode' and a unified controller for currently active element
			TeamState:        explore.NewTeamState(),
			TotalElapsedTime: 12 * 1000 * 60,
			DialogHandler:    explore.NewDialogHandler(),
			ModeManager:      explore.NewModeManager(),
			Camera:           explore.NewCamera(),
		},
		fightState: &fight.State{
			NextMode:         common.NoneMode,
			ActiveTeam:       nil,
			PlayerController: fight.NewPlayerController(),
			PlayerTeam:       nil,
			AiController:     fight.AiController{},
			AiTeam:           nil,
			Camera:           fight.NewCamera(),
		},
		dialog:      dialog.NewUi(),
		shopUI:      gui.NewShopUi(),
		inventoryUI: inventory.NewUi(),
		menuUI:      menu.NewUI(),
		equipmentUI: equipment.NewUI(),
		mode:        common.ExploreMode,
	}
	g.state.Map.LoadLevel("test-level-a")
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
	state            *explore.State
	fightState       *fight.State

	dialog      gui.UI
	shopUI      *gui.ShopUI
	inventoryUI gui.UI
	menuUI      gui.UI
	equipmentUI gui.UI
	mode        common.Mode
}

func (g *Game) Update() error {
	// calculate delta
	delta := time.Now().Sub(g.lastUpdateCalled).Milliseconds()
	g.lastUpdateCalled = time.Now()
	g.state.TotalElapsedTime = g.state.TotalElapsedTime + delta

	switch g.mode {
	case common.ExploreMode:
		// update state
		g.state.Update(delta)

		// update UI
		g.dialog.Update(delta, g.state)
		g.shopUI.Update(delta, g.state)
		g.inventoryUI.Update(delta, g.state)
		g.menuUI.Update(delta, g.state)
		g.equipmentUI.Update(delta, g.state)

		// check for mode change
		if g.state.ModeManager.NextMode == common.FightMode {
			g.state.ModeManager.NextMode = common.NoneMode
			g.mode = common.FightMode
			g.StartFightMode()
		}

	case common.FightMode:
		// update state
		g.fightState.Update(delta)

		// check for mode change
		if g.fightState.NextMode == common.ExploreMode {
			g.fightState.NextMode = common.NoneMode
			g.mode = common.ExploreMode
		}
	}

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
	switch g.mode {
	case common.ExploreMode:
		g.state.Map.Level.Draw(g.state.Camera)
		g.state.Player.Draw(g.state.Camera)
		g.state.Camera.DrawBuffer(screen)
		g.dialog.Draw(screen)
		g.shopUI.Draw(screen)
		g.inventoryUI.Draw(screen)
		g.menuUI.Draw(screen)
		g.equipmentUI.Draw(screen)
	case common.FightMode:
		g.fightState.Draw(g.fightState.Camera)
		g.fightState.Camera.DrawBuffer(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return common.ScreenWidth * common.Scale, common.ScreenHeight * common.Scale
}

func (g *Game) StartFightMode() {
	playerActors := []*fight.Actor{
		{},
	}
	enemyActors := []*fight.Actor{
		{},
	}
	g.fightState.StartFight(playerActors, enemyActors, "forest-scene")
}
