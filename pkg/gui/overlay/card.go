package overlay

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"github.com/seanoneillcode/go-tactics/pkg/fight"
)

type Card struct {
	portrait  *common.Sprite
	turnToken *common.Sprite
}

func (c Card) Draw(screen *ebiten.Image) {

}

func (c Card) Update(state *fight.State) {

}
