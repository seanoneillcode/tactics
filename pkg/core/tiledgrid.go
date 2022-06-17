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
	resourceDirectory = "res/levels/"
)

type TiledGrid struct {
	//image             *ebiten.Image
	Layers            []*Layer            `json:"layers"`
	TileSetReferences []*TileSetReference `json:"tilesets"`
	TileSet           []*TileSet
}

type Layer struct {
	Data    []int         `json:"Data"`
	Height  int           `json:"height"`
	Width   int           `json:"width"`
	Objects []TiledObject `json:"objects"`
}

type TiledObject struct {
	Name       string            `json:"Name"`
	Type       string            `json:"type"`
	X          int               `json:"x"`
	Y          int               `json:"y"`
	Properties []*TileConfigProp `json:"properties"`
}

type TileSetReference struct {
	Source   string `json:"source"`
	FirstGid int    `json:"firstgid"`
}

type TileSet struct {
	ImageFileName string `json:"image"`
	ImageWidth    int    `json:"imagewidth"`
	ImageHeight   int    `json:"imageheight"`
	numTilesX     int
	numTilesY     int
	FirstGid      int
	Tiles         []*TileConfig `json:"tiles"`
	image         *ebiten.Image
}

type TileConfig struct {
	Id         int               `json:"id"`
	Properties []*TileConfigProp `json:"properties"`
}

type TileConfigProp struct {
	Name  string      `json:"Name"`
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

func NewTileGrid(fileName string) *TiledGrid {
	var tiledGrid TiledGrid
	//var tileSet TileSet

	configFile, err := os.Open(filepath.Join(resourceDirectory, fileName))
	if err != nil {
		log.Fatal("opening config file", err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&tiledGrid); err != nil {
		log.Fatal("parsing config file", err.Error())
	}

	tiledGrid.TileSet = []*TileSet{}
	for _, ref := range tiledGrid.TileSetReferences {
		tiledGrid.TileSet = append(tiledGrid.TileSet, loadTileSet(ref))
	}

	return &tiledGrid
}

func loadTileSet(ref *TileSetReference) *TileSet {
	tileSetConfigFile, err := os.Open(filepath.Join(resourceDirectory, ref.Source))
	if err != nil {
		log.Fatal("opening config file", err.Error())
	}

	var tileSet TileSet
	jsonParser := json.NewDecoder(tileSetConfigFile)
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

	tileSet.image = ebiten.NewImageFromImage(img)
	tileSet.FirstGid = ref.FirstGid
	return &tileSet
}

func (tg *TiledGrid) Draw(screen *ebiten.Image) {
	for _, layer := range tg.Layers {
		for i, tileIndex := range layer.Data {
			if tileIndex == 0 {
				continue
			}

			ts := tg.getTileSetForIndex(tileIndex)

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(((i)%layer.Width)*common.TileSize), float64(((i)/layer.Height)*common.TileSize))
			op.GeoM.Scale(common.Scale, common.Scale)

			sx := ((tileIndex - ts.FirstGid) % ts.numTilesX) * common.TileSize
			sy := ((tileIndex - ts.FirstGid) / ts.numTilesX) * common.TileSize

			screen.DrawImage(ts.image.SubImage(image.Rect(sx, sy, sx+common.TileSize, sy+common.TileSize)).(*ebiten.Image), op)
		}
	}
}

func (tg *TiledGrid) getTileSetForIndex(index int) *TileSet {
	for i, tileSet := range tg.TileSet {
		if i == len(tg.TileSet)-1 || tg.TileSet[i+1].FirstGid > index {
			return tileSet
		}
	}
	// should never happen
	return nil
}

type ObjectData struct {
	name       string
	objectType string
	x          int
	y          int
	properties []*ObjectProperty
}

type ObjectProperty struct {
	name    string
	objType string
	value   interface{}
}

func (tg *TiledGrid) GetObjectData() []*ObjectData {
	var ods []*ObjectData
	for _, l := range tg.Layers {
		for _, obj := range l.Objects {
			od := &ObjectData{
				name:       obj.Name,
				objectType: obj.Type,
				x:          obj.X,
				y:          obj.Y,
				properties: []*ObjectProperty{},
			}
			for _, p := range obj.Properties {
				od.properties = append(od.properties, &ObjectProperty{
					name:    p.Name,
					objType: p.Type,
					value:   p.Value,
				})
			}
			ods = append(ods, od)
		}
	}
	return ods
}

type TileData struct {
	x       int
	y       int
	isBlock bool
}

func (tg *TiledGrid) GetTileData(x int, y int) *TileData {
	td := TileData{
		x: x,
		y: y,
	}
	index := (y * tg.Layers[0].Width) + x

	if index < 0 || index >= len(tg.Layers[0].Data) {
		// no tile here
		td.isBlock = true
		return &td
	}

	if x < 0 || y < 0 {
		// no tile here
		td.isBlock = true
		return &td
	}

	tileSetIndex := tg.Layers[0].Data[index]
	if tileSetIndex == 0 {
		td.isBlock = true
		return &td
	}

	ts := tg.getTileSetForIndex(tileSetIndex)

	for _, tile := range ts.Tiles {
		if tile.Id == tileSetIndex-ts.FirstGid {
			for _, prop := range tile.Properties {
				if prop.Name == "isBlock" && prop.Value != nil {
					td.isBlock = (prop.Value).(bool)
				}
				break
			}
			break
		}
	}

	return &td
}
