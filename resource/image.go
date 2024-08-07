package resource

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewImage(w, h int, clr ...color.Color) *ebiten.Image {
	img := ebiten.NewImage(w, h)
	if len(clr) != 0 {
		img.Fill(clr[0])
	}

	return img
}
