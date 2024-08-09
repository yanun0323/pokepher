package app

import (
	"fmt"
	"image/color"
	"main/internal/domain"
	"main/internal/domain/entity"
	"main/internal/domain/usecase"
	"main/internal/service"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yanun0323/ebitenpkg"
)

type Game struct {
	background ebitenpkg.Image
	battle     usecase.Stage
}

func NewGame() *Game {
	background := ebitenpkg.NewRectangle(domain.DefaultWindowsWidth(), domain.DefaultWindowsHeight(), color.White).
		Align(ebitenpkg.AlignTopLeading)

	playerPokemon := entity.NewPokemon(entity.PokemonOption{
		ID:               "025",
		Level:            10,
		AbilityIdx:       0,
		HiddenAbility:    false,
		Gender:           entity.Male,
		Nature:           entity.NatureAdamant,
		IndividualValues: entity.StatsRandomIV(),
	})

	opponentPokemon := entity.NewPokemon(entity.PokemonOption{
		ID:               "025",
		Level:            10,
		AbilityIdx:       1,
		HiddenAbility:    false,
		Gender:           entity.Female,
		Nature:           entity.NatureAdamant,
		IndividualValues: entity.StatsRandomIV(),
	})

	return &Game{
		background: background,
		battle:     service.NewBattleStage(nil, nil, []*entity.Pokemon{playerPokemon}, []*entity.Pokemon{opponentPokemon}),
	}
}

func (g *Game) Start() {
	cfg := entity.PokemonConfigTable["025"]
	fmt.Printf("%+v\n", cfg)
}

func (g *Game) Update() error {
	ebitenpkg.GameUpdate()
	g.battle.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.background.Draw(screen)

	if g.battle != nil {
		g.battle.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return domain.DefaultWindowsBounds()
}
