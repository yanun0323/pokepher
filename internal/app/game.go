package app

import (
	"fmt"
	"image/color"
	"main/internal/domain/entity"
	"main/internal/domain/service"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yanun0323/ebitenpkg"
)

type Game struct {
	background ebitenpkg.Image
	battle     service.Stage
}

func NewGame() *Game {
	backgroundImg := ebiten.NewImage(1366, 768)
	backgroundImg.Fill(color.White)
	background := ebitenpkg.NewImage(backgroundImg).Align(ebitenpkg.AlignTopLeading)

	return &Game{
		background: background,
	}
}

func (g *Game) Start() {
	cfg := entity.PokemonConfigTable[25]
	fmt.Printf("%+v\n", cfg)
}

func (g *Game) Update() error {
	ebitenpkg.GameUpdate()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.battle != nil {
		g.battle.Draw(screen)
	}

	g.background.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 1366, 768
}
