package component

import (
	"image"
	"image/color"
	"main/resource"

	"github.com/yanun0323/ebitenpkg"
)

func NewBox(w, h int, children ...ebitenpkg.Attachable) ebitenpkg.Image {
	s := 2.0

	w = int(float64(w) / s)
	h = int(float64(h) / s)

	pokeballBounds := resource.Global.Pokeball.Bounds()
	pokeballBounds = image.Rect(0, 0, pokeballBounds.Dx()-2, pokeballBounds.Dy()-2)
	pokeballWidth := float64(pokeballBounds.Dx() - 1)
	pokeballHeight := float64(pokeballBounds.Dy() - 1)

	pokeballCoordX := float64(w) - pokeballWidth - 1.5
	pokeballCoordY := float64(h) - pokeballHeight - 1.5

	return ebitenpkg.NewCanvas(w, h, color.RGBA{0, 0, 0, 64}).
		DrawRectOn(w-pokeballBounds.Dx()*2, h-pokeballBounds.Dy()*2, color.RGBA{255, 255, 255, 50}, pokeballWidth, pokeballHeight).
		DrawImageOn(resource.Global.Pokeball, 0, 0).
		DrawImageOn(resource.Global.Pokeball, pokeballCoordX, 0).
		DrawImageOn(resource.Global.Pokeball, 0, pokeballCoordY).
		DrawImageOn(resource.Global.Pokeball, pokeballCoordX, pokeballCoordY).
		Image(children...).
		Align(ebitenpkg.AlignTopLeading).
		Scale(s, s)
}
