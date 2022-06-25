package explore

import (
	"fmt"
	"github.com/seanoneillcode/go-tactics/pkg/common"
)

type ModeManager struct {
	NextMode common.Mode
}

func NewModeManager() *ModeManager {
	return &ModeManager{
		NextMode: common.NoneMode,
	}
}

func (r *ModeManager) Update(delta int64, state *State) {

}

func (r *ModeManager) ChangeMode(mode common.Mode, enemy *Enemy) {
	r.NextMode = mode
	fmt.Println("starting fight mode with: ", enemy.name)
}
