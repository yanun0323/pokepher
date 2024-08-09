package entity

import (
	"main/resource"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yanun0323/ebitenpkg"
)

type Gender int

const (
	Genderless Gender = iota
	Male
	Female
)

func (g Gender) Image() *ebiten.Image {
	switch g {
	case Male:
		return resource.Global.SymbolMale
	case Female:
		return resource.Global.SymbolFemale
	default:
		b := resource.Global.SymbolMale.Bounds()
		return ebitenpkg.NewEbitenImage(b.Dx(), b.Dy())
	}
}
