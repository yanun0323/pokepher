package component

import (
	"fmt"
	"image/color"
	"main/internal/domain/entity"
	"main/internal/domain/usecase"
	"main/resource"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yanun0323/ebitenpkg"
)

type PokemonBattleStatusUI struct {
	Pokemon    *entity.Pokemon
	Background ebitenpkg.Image
	Gender     ebitenpkg.Image
	Name       ebitenpkg.Text
	Level      ebitenpkg.Text
	HP         ebitenpkg.Text
	HPLabel    ebitenpkg.Text
	HPBar      ebitenpkg.Image
	ExpLabel   ebitenpkg.Text
	ExpBar     ebitenpkg.Image
}

func NewPokemonBattleStatusUI(pokemon *entity.Pokemon) usecase.View {
	if pokemon == nil {
		return nil
	}

	ui := &PokemonBattleStatusUI{
		Pokemon: pokemon,
	}

	ui.Background = ebitenpkg.NewImage(resource.Battle.UIStatus).
		Align(ebitenpkg.AlignTopLeading)
	ui.Gender = ebitenpkg.NewImage(pokemon.Gender.Image()).
		Align(ebitenpkg.AlignTopLeading).
		Attach(ui.Background)
	ui.Name = ebitenpkg.NewText(pokemon.Name, 20).
		Align(ebitenpkg.AlignTopLeading).
		SetColor(color.Black).
		Attach(ui.Background)
	ui.Level = ebitenpkg.NewText(fmt.Sprintf("Lv.%d", pokemon.Level), 20).
		Align(ebitenpkg.AlignTopLeading).
		SetColor(color.Black).
		Attach(ui.Background)

	ui.HP = ebitenpkg.NewText(fmt.Sprintf("%d/%d", pokemon.HP, pokemon.MaxHP), 20).
		Align(ebitenpkg.AlignTopLeading).
		Attach(ui.Background)
	ui.HPLabel = ebitenpkg.NewText("HP", 20).
		Align(ebitenpkg.AlignTopLeading).
		SetColor(color.Black).
		Attach(ui.Background)
	ui.HPBar = ebitenpkg.NewImage(resource.NewImage(100, 20, color.RGBA{0, 255, 0, 255})).
		Align(ebitenpkg.AlignTopLeading).
		Attach(ui.Background)

	ui.ExpLabel = ebitenpkg.NewText("EXP", 20).
		Align(ebitenpkg.AlignTopLeading).
		SetColor(color.Black).
		Attach(ui.Background)
	ui.ExpBar = ebitenpkg.NewImage(resource.NewImage(100, 20, color.RGBA{0, 0, 255, 255})).
		Align(ebitenpkg.AlignTopLeading).
		Attach(ui.Background)

	return ui
}

func (b *PokemonBattleStatusUI) Update() error {
	return nil
}

func (b *PokemonBattleStatusUI) Draw(screen *ebiten.Image) {
	b.Background.Draw(screen)
	b.Gender.Draw(screen)
	b.Name.Draw(screen)
	b.Level.Draw(screen)
	b.HP.Draw(screen)
	b.HPLabel.Draw(screen)
	b.HPBar.Draw(screen)
	b.ExpLabel.Draw(screen)
	b.ExpBar.Draw(screen)
}
