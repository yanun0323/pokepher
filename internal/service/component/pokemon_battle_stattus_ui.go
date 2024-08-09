package component

import (
	"fmt"
	"image/color"
	"main/internal/domain/entity"
	"main/resource"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yanun0323/ebitenpkg"
)

type pokemonBattleStatusUI struct {
	ebitenpkg.Image
	Pokemon *entity.Pokemon
}

func NewPokemonBattleStatusUI(pokemon *entity.Pokemon, w, h int, showExp bool, fn func(ebitenpkg.Image)) ebitenpkg.Image {
	if pokemon == nil {
		return nil
	}

	ui := &pokemonBattleStatusUI{
		Pokemon: pokemon,
	}

	hx := func(ratio float64) float64 {
		return float64(h) * ratio
	}

	wx := func(ratio float64) float64 {
		return float64(w) * ratio
	}

	padding := 20.0

	hpViewWidth := w - int(padding*2)
	hpViewHeight := int(hx(0.25))
	hpLabelWidth := int(hx(0.33))
	hpBarPadding := 25
	hpBarWidth := hpViewWidth - hpLabelWidth - hpBarPadding
	hpBarHeight := hpViewHeight - hpBarPadding
	hpView := ebitenpkg.NewRoundedRectangle(hpViewWidth, hpViewHeight, hpViewHeight/2, color.Black,
		// HpLabel
		ebitenpkg.NewText("H P", hx(0.15)).
			SetFont(resource.Font.Headline).
			Align(ebitenpkg.AlignTopTrailing).
			SetColor(color.White).
			Move(-float64(hpViewWidth-hpLabelWidth), float64(hpBarPadding)/2.2),
		// HpBarBackground
		ebitenpkg.NewRoundedRectangle(hpBarWidth, hpBarHeight, hpBarHeight/2, color.RGBA{0, 100, 0, 100},
			// HpBar
			ebitenpkg.NewRoundedRectangle(hpBarWidth/3, hpBarHeight, hpBarHeight/2, color.RGBA{50, 245, 50, 255}).
				Align(ebitenpkg.AlignTopLeading),
			// HpValue
			ebitenpkg.NewText(fmt.Sprintf("%d / %d", pokemon.HP, pokemon.MaxHP), hx(0.125)).
				SetFont(resource.Font.Headline).
				Align(ebitenpkg.AlignCenter).
				SetColor(color.White).
				Move(float64(hpBarWidth)/2, float64(hpBarHeight)/2),
		).Align(ebitenpkg.AlignTopLeading).
			Move(-float64(hpBarWidth)-float64(hpBarPadding/2)-2, float64(hpBarPadding/2)),
	).Align(ebitenpkg.AlignTopTrailing).
		Move(-padding, hx(0.4))

	expView := ebitenpkg.NewRectangle(1, 1, color.Transparent)
	if showExp {
		expViewWidth := w - int(padding) - int(wx(0.1))
		expViewHeight := int(hx(0.18))
		expLabelWidth := int(hx(0.28))
		expBarPadding := 25
		expBarWidth := expViewWidth - expLabelWidth - expBarPadding
		expBarHeight := expViewHeight - expBarPadding
		expView = ebitenpkg.NewRoundedRectangle(expViewWidth, expViewHeight, expViewHeight/2, color.Black,
			// ExpLabel
			ebitenpkg.NewText("E X P", hx(0.1)).
				SetFont(resource.Font.Headline).
				Align(ebitenpkg.AlignTopTrailing).
				SetColor(color.White).
				Move(-float64(expViewWidth-expLabelWidth), float64(expBarPadding)/2.2),
			// ExpBarBackground
			ebitenpkg.NewRoundedRectangle(expBarWidth, expBarHeight, expBarHeight/2, color.RGBA{10, 50, 128, 255},
				// ExpBar
				ebitenpkg.NewRoundedRectangle(expBarWidth/3, expBarHeight, expBarHeight/2, color.RGBA{0, 128, 255, 255}).
					Align(ebitenpkg.AlignTopLeading),
			).Align(ebitenpkg.AlignTopLeading).
				Move(-float64(expBarWidth)-float64(expBarPadding/2)-2, float64(expBarPadding/2)),
		).Align(ebitenpkg.AlignTopTrailing).
			Move(-padding, hx(0.7))
	}

	statusWidth := w - int(padding*2)
	statusHeight := int(hx(0.18))
	status := ebitenpkg.NewRectangle(statusWidth, statusHeight, color.Transparent,
		// Name
		ebitenpkg.NewText(strings.ToUpper(pokemon.Name), hx(0.2)).
			SetFont(resource.Font.Headline).
			Align(ebitenpkg.AlignTopLeading).
			SetColor(color.Black).
			Move(float64(-statusWidth), 0),
		// Gender
		ebitenpkg.NewImage(pokemon.Gender.Image()).
			Align(ebitenpkg.AlignTopTrailing).
			Move(-padding, 3).
			Scale(hx(0.018), hx(0.018)),
		// Level
		ebitenpkg.NewText(fmt.Sprintf("Lv.%d", pokemon.Level), hx(0.2)).
			SetFont(resource.Font.Headline).
			Align(ebitenpkg.AlignTopLeading).
			SetColor(color.Black).
			Move(-hx(0.8)-padding, 0),
		// Status
		ebitenpkg.NewRoundedRectangle(100, statusHeight, 7, color.RGBA{0xff, 0, 0xff, 0xff},
			// StatusText
			ebitenpkg.NewText(pokemon.Status.String(), hx(0.15)).
				SetFont(resource.Font.Headline).
				Align(ebitenpkg.AlignCenter).
				SetColor(color.White).
				Move(-50, float64(statusHeight/2)),
		).Align(ebitenpkg.AlignTopTrailing).
			Move(-hx(1.0)-padding, 0),
	).Align(ebitenpkg.AlignTopLeading).
		Move(0, 20)

	if !showExp {
		h = int(hx(0.8))
	}

	ui.Image = ebitenpkg.NewRoundedRectangle(w, h, 15, color.RGBA{0, 0, 0, 25},
		status,
		hpView,
		expView,
	).Align(ebitenpkg.AlignTopTrailing)

	fn(ui.Image)

	return ui
}

func (b *pokemonBattleStatusUI) Update() error {
	return nil
}

func (b *pokemonBattleStatusUI) Draw(screen *ebiten.Image) {
	b.Image.Draw(screen)
}
