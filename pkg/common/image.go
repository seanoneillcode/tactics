package common

import (
	"bytes"
	"image"
	"io/ioutil"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func LoadImage(imageFileName string) *ebiten.Image {
	b, err := ioutil.ReadFile("res/" + imageFileName)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}
	return ebiten.NewImageFromImage(img)
}
