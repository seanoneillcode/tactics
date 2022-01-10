package main

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"github.com/seanoneillcode/go-tactics/pkg/sprite"
	"image"
	_ "image/png"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	tilesImage *ebiten.Image
)

func init() {
	b, err := ioutil.ReadFile("res/tiles.png")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}
	tilesImage = ebiten.NewImageFromImage(img)
}

type Game struct {
	layers           [][][]int
	keys             []ebiten.Key
	player           *sprite.Character
	lastUpdateCalled time.Time
}

func (g *Game) Update() error {
	delta := time.Now().Sub(g.lastUpdateCalled).Milliseconds()

	g.player.Update(delta)

	g.keys = inpututil.AppendPressedKeys(g.keys[:0])

	g.lastUpdateCalled = time.Now()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	for _, l := range g.layers {
		for iy, line := range l {
			for ix, t := range line {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64((ix)*common.TileSize), float64((iy)*common.TileSize))
				op.GeoM.Scale(common.Scale, common.Scale)

				sx := (t % common.TileXNum) * common.TileSize
				sy := (t / common.TileXNum) * common.TileSize
				screen.DrawImage(tilesImage.SubImage(image.Rect(sx, sy, sx+common.TileSize, sy+common.TileSize)).(*ebiten.Image), op)
			}
		}
	}

	g.player.Draw(screen)

	keyStrs := []string{}
	for _, p := range g.keys {
		keyStrs = append(keyStrs, p.String())
	}
	ebitenutil.DebugPrint(screen, strings.Join(keyStrs, ", "))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return common.ScreenWidth * common.Scale, common.ScreenHeight * common.Scale
}

func main() {
	g := &Game{
		lastUpdateCalled: time.Now(),
		player:           sprite.NewCharacter(),
		layers: [][][]int{
			{
				[]int{243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243},
				[]int{243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243},
				[]int{243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243},
				[]int{243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 243, 244, 243, 243},
				[]int{243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243},

				[]int{243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243},
				[]int{243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243},
				[]int{243, 243, 244, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243},
				[]int{243, 243, 243, 243, 243, 243, 243, 243, 243, 219, 243, 243, 243, 219, 243, 243},
				[]int{243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243},

				[]int{243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243},
				[]int{243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243},
				[]int{243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243},
				[]int{243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 244, 243, 243, 243, 243},
				[]int{243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243},
			},
			{
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 26, 27, 28, 29, 30, 31, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 51, 52, 53, 54, 55, 56, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 76, 77, 78, 79, 80, 81, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 101, 102, 103, 104, 105, 106, 0, 0, 0, 0, 0},

				[]int{0, 0, 0, 0, 0, 126, 127, 128, 129, 130, 131, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 303, 303, 245, 242, 303, 303, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0, 0},

				[]int{0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0, 0},
			},
		},
	}

	ebiten.SetWindowSize(common.ScreenWidth*common.Scale, common.ScreenHeight*common.Scale)
	ebiten.SetWindowTitle("Fantasy Game")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
