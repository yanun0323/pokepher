package usecase

import "github.com/hajimehoshi/ebiten/v2"

type View interface {
	Update() error
	Draw(*ebiten.Image)
}
