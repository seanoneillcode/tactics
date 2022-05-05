package menu

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/seanoneillcode/go-tactics/pkg/core"
	"github.com/seanoneillcode/go-tactics/pkg/gui/elem"
	"time"
)

var listPos = &elem.Pos{X: 220, Y: 16}

type ui struct {
	highlight *elem.Sprite
	bg        *elem.StaticImage
	options   []*elem.Text
	location  *elem.Text
	time      *elem.Text
	money     *elem.Text
	// state
	selectedListIndex int
	IsActive          bool
	cards             []*elem.CharacterCard
	isLoaded          bool
}

func NewUI() *ui {
	textOffset := &elem.Pos{X: 4, Y: 4}
	u := &ui{
		bg:        elem.NewStaticImage("uis/menu/menu-bg.png", 0, 0),
		highlight: elem.NewSprite("uis/menu/menu-highlight.png", float64(listPos.X), float64(listPos.Y)),
		options: []*elem.Text{
			elem.NewText(listPos.X+textOffset.X, 16+textOffset.Y, "Items"),
			elem.NewText(listPos.X+textOffset.X, 32+textOffset.Y, "Equipment"),
			elem.NewText(listPos.X+textOffset.X, 48+textOffset.Y, "Magic"),
			elem.NewText(listPos.X+textOffset.X, 64+textOffset.Y, "Settings"),
			elem.NewText(listPos.X+textOffset.X, 80+textOffset.Y, "File"),
			elem.NewText(listPos.X+textOffset.X, 96+textOffset.Y, "Exit"),
		},
		money:    elem.NewText(listPos.X+textOffset.X, 144+textOffset.Y, "money:"),
		location: elem.NewText(listPos.X+textOffset.X, 144+16+textOffset.Y, "location:"),
		time:     elem.NewText(listPos.X+textOffset.X, 144+32+textOffset.Y, "time:"),
	}
	return u
}

func (r *ui) Draw(screen *ebiten.Image) {
	if !r.IsActive {
		return
	}
	r.bg.Draw(screen)
	r.highlight.Draw(screen)
	for _, o := range r.options {
		o.Draw(screen)
	}
	for _, card := range r.cards {
		card.Draw(screen)
	}
	r.location.Draw(screen)
	r.time.Draw(screen)
	r.money.Draw(screen)
}

func (r *ui) Update(delta int64, state *core.State) {
	if !state.UI.IsMenuActive() {
		r.IsActive = false
		r.isLoaded = false
		return
	}
	r.IsActive = true
	if !r.isLoaded {
		r.isLoaded = true
		r.rebuild(state.TeamState.Characters)
		r.location.SetValue(fmt.Sprintf("location: %s", state.Map.Level.Name))
		r.money.SetValue(fmt.Sprintf("money: %s", fmt.Sprintf("%d", state.TeamState.Money)))
		return
	}
	r.handleInput(state)
	highlightPos := &elem.Pos{X: listPos.X, Y: listPos.Y}
	highlightPos.Y = listPos.Y + (16.0 * r.selectedListIndex)
	r.highlight.SetPos(highlightPos)
	duration := time.Duration(state.TotalElapsedTime) * time.Millisecond
	z := time.Unix(0, 0).UTC().Add(duration)
	r.time.SetValue(fmt.Sprintf("time: %s", z.Format("15:04:05")))
}

func (r *ui) handleInput(state *core.State) {
	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
		state.UI.Close()
		state.Player.Activate()
		r.reset()
		r.IsActive = false
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		state.UI.Close()
		state.Player.Activate()
		r.reset()
		r.IsActive = false
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		r.selectedListIndex = r.selectedListIndex - 1
		if r.selectedListIndex < 0 {
			r.selectedListIndex = 0
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		r.selectedListIndex = r.selectedListIndex + 1
		if r.selectedListIndex == 6 {
			r.selectedListIndex = r.selectedListIndex - 1
			if r.selectedListIndex < 0 {
				r.selectedListIndex = 0
			}
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		switch r.selectedListIndex {
		// items
		case 0:
			state.UI.Open(core.ItemsUI)
		case 1:
			state.UI.Open(core.EquipmentUI)
		case 2:
			state.UI.Open(core.MagicUI)
		case 3:
			state.UI.Open(core.SettingsUI)
		case 4:
			state.UI.Open(core.FileUI)
		case 5:
			state.Control.ExitGame()
		}
	}
}

func (r *ui) reset() {
	r.selectedListIndex = 0
}

func (r *ui) rebuild(characters []*core.CharacterState) {
	var cards []*elem.CharacterCard
	pos := elem.Pos{
		X: 12,
		Y: 8,
	}
	for _, c := range characters {
		cards = append(cards, elem.NewCharacterCard(c, pos))
		pos.Y = pos.Y + 48 + 8
	}
	r.cards = cards
}
