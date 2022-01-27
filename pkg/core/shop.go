package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"log"
)

type Shop struct {
	data     *ShopData
	isActive bool
	// rendering
	x       float64
	y       float64
	bgImage *ebiten.Image
}

func NewShop() *Shop {
	s := &Shop{
		bgImage: LoadImage("shop-bg.png"),
	}
	return s
}

func (s *Shop) Open(data *ShopData) {
	s.isActive = true
	s.data = data
}

func (s *Shop) Update(delta int64, state *State) {
	if !s.isActive {
		return
	}
	if s.data == nil {
		log.Fatalf("opened a shop with no data")
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
		// p.ActiveShop.Reset() ?
		s.isActive = false // close shop
		state.Player.Activate()
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		log.Printf("shop is open, pressed accept")
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		log.Printf("shop is open, moved left")
		//s.navigateLeft()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		log.Printf("shop is open, moved right")
		//s.navigateRight()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		log.Printf("shop is open, moved up")
		//s.navigateUp()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		log.Printf("shop is open, moved down")
		//s.navigateDown()
	}
}

func (s *Shop) Draw(screen *ebiten.Image) {
	if !s.isActive {
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(common.Scale, common.Scale)

	screen.DrawImage(s.bgImage, op)
}
