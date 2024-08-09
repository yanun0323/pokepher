package component

import (
	"image/color"
	"main/resource"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yanun0323/ebitenpkg"
)

type textBox struct {
	ebitenpkg.Text
	box ebitenpkg.Image
}

func NewTextBox(text string, w, h int, x, y float64) ebitenpkg.Text {
	padding := 50.0

	t := ebitenpkg.NewText(text, 50).
		Align(ebitenpkg.AlignTopLeading).
		SetFont(resource.Font.Regular).
		SetColor(color.Black).
		Move(padding, padding)

	box := NewBox(w, h, t).
		Move(x, y)

	return &textBox{
		Text: t,
		box:  box,
	}
}

func (t *textBox) Update() error {
	return nil
}

func (t *textBox) Draw(screen *ebiten.Image) {
	t.box.Draw(screen)
	t.Text.Draw(screen)
}
