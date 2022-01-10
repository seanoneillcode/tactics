package core

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"image"
	_ "image/png"
	"io/ioutil"
	"log"
)

type TileGrid struct {
	image  *ebiten.Image
	layers [][][]int
}

func NewTileGrid() *TileGrid {
	b, err := ioutil.ReadFile("res/tiles.png")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}
	return &TileGrid{
		image: ebiten.NewImageFromImage(img),
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
}

func (l *TileGrid) Update(delta int64) {

}

func (l *TileGrid) Draw(screen *ebiten.Image) {
	for _, layer := range l.layers {
		for iy, line := range layer {
			for ix, t := range line {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64((ix)*common.TileSize), float64((iy)*common.TileSize))
				op.GeoM.Scale(common.Scale, common.Scale)

				sx := (t % common.TileXNum) * common.TileSize
				sy := (t / common.TileXNum) * common.TileSize
				screen.DrawImage(l.image.SubImage(image.Rect(sx, sy, sx+common.TileSize, sy+common.TileSize)).(*ebiten.Image), op)
			}
		}
	}
}
