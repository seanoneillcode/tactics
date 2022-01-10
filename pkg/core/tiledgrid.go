package core

import (
	"bytes"
	"encoding/json"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"image"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const (
	resourceDirectory = "res/"
)

type TiledGrid struct {
	image             *ebiten.Image
	Layers            []*Layer            `json:"layers"`
	TileSetReferences []*TileSetReference `json:"tilesets"`
	TileSet           []*TileSet
}

type Layer struct {
	Data   []int `json:"data"`
	Height int   `json:"height"`
	Width  int   `json:"width"`
}

type TileSetReference struct {
	Source string `json:"source"`
}

type TileSet struct {
	ImageFileName string `json:"image"`
	ImageWidth    int    `json:"imagewidth"`
	ImageHeight   int    `json:"imageheight"`
	numTilesX     int
	numTilesY     int
}

func NewTileGrid() *TiledGrid {
	var tiledGrid TiledGrid
	var tileSet TileSet
	var tiledFileName = filepath.Join(resourceDirectory, "home.json")

	configFile, err := os.Open(tiledFileName)
	if err != nil {
		log.Fatal("opening config file", err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&tiledGrid); err != nil {
		log.Fatal("parsing config file", err.Error())
	}

	tileSetConfigFile, err := os.Open(filepath.Join(resourceDirectory, tiledGrid.TileSetReferences[0].Source))
	if err != nil {
		log.Fatal("opening config file", err.Error())
	}

	jsonParser = json.NewDecoder(tileSetConfigFile)
	if err = jsonParser.Decode(&tileSet); err != nil {
		log.Fatal("parsing config file", err.Error())
	}
	tileSet.numTilesX = tileSet.ImageWidth / common.TileSize
	tileSet.numTilesY = tileSet.ImageHeight / common.TileSize

	b, err := ioutil.ReadFile(filepath.Join(resourceDirectory, tileSet.ImageFileName))
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}
	tiledGrid.image = ebiten.NewImageFromImage(img)
	tiledGrid.TileSet = []*TileSet{&tileSet}
	return &tiledGrid
}

func (l *TiledGrid) Update(delta int64) {

}

func (l *TiledGrid) Draw(screen *ebiten.Image) {
	for _, layer := range l.Layers {
		for i, t := range layer.Data {

			if t == 0 {
				continue
			}

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((i%layer.Width)*common.TileSize), float64((i/layer.Height)*common.TileSize))
			op.GeoM.Scale(common.Scale, common.Scale)

			sx := (t%l.TileSet[0].numTilesX - 1) * common.TileSize
			sy := (t / l.TileSet[0].numTilesY) * common.TileSize

			screen.DrawImage(l.image.SubImage(image.Rect(sx, sy, sx+common.TileSize, sy+common.TileSize)).(*ebiten.Image), op)

		}
	}
}
