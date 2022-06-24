package core

import "fmt"

const FightFadeTime = 300

type GameMode string

const (
	ExploreGameMode   = "explore"
	FightGameMode     = "fight"
	ShopGameMode      = "shop"
	InventoryGameMode = "inventory"
)

type ModeManager struct {
	currentMode GameMode
	oldMode     GameMode
}

func NewModeManager() *ModeManager {
	return &ModeManager{}
}

func (r *ModeManager) Update(delta int64, state *State) {
	if r.oldMode != r.currentMode {
		if r.currentMode == FightGameMode {
			state.Map.fader.FadeOutAndIn(FightFadeTime)
		}

		r.oldMode = r.currentMode
	}
}

func (r *ModeManager) SwitchToFightMode(enemy *Enemy) {
	r.currentMode = FightGameMode
	fmt.Println("starting fight mode with: ", enemy.name)
}

func (r *ModeManager) SwitchToExploreMode() {
	r.currentMode = ExploreGameMode
}
