package service

import "github.com/hajimehoshi/ebiten/v2"

type Stage interface {
	Update() error
	Draw(screen *ebiten.Image)
}
