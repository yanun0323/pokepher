package service

import (
	"main/internal/domain"
	"main/internal/domain/entity"
	"main/internal/domain/usecase"
	"main/internal/service/component"
	"main/internal/utils"
	"main/resource"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yanun0323/ebitenpkg"
	"github.com/yanun0323/pkg/logs"
)

type battleStage int

const (
	_battleStageEntry battleStage = iota
	_battleStageSelect
	_battleStageWin
	_battleStageLose
)

type battle struct {
	GameTime    int
	BattleStage *utils.Value[battleStage]
	TextBox     ebitenpkg.Text

	Player   *entity.Character
	Opponent *entity.Character

	PlayerPokemon   []*entity.Pokemon
	OpponentPokemon []*entity.Pokemon

	PlayerPokemonImage          ebitenpkg.Image
	PlayerPokemonBattleStatusUI ebitenpkg.Image

	OpponentPokemonImage          ebitenpkg.Image
	OpponentPokemonBattleStatusUI ebitenpkg.Image
}

func NewBattleStage(Player, Opponent *entity.Character, PlayerPokemon, OpponentPokemon []*entity.Pokemon) usecase.Stage {
	wWidth, wHeight := domain.DefaultWindowsBounds()
	textBoxHeight := 200
	battle := &battle{
		GameTime:        ebitenpkg.CurrentGameTime(),
		BattleStage:     utils.NewValue(_battleStageEntry),
		PlayerPokemon:   PlayerPokemon,
		OpponentPokemon: OpponentPokemon,
		TextBox:         component.NewTextBox("Hello World", wWidth, textBoxHeight, 0, float64(wHeight-textBoxHeight)),
	}

	battleStatusUIWidth := 600
	battleStatusUIHeight := 200

	if len(PlayerPokemon) != 0 {
		img, err := resource.LoadPokemonImage(PlayerPokemon[0].ID.String())
		if err != nil {
			logs.Errorf("failed to load player pokemon image: %v", err)
		} else {
			battle.PlayerPokemonImage = ebitenpkg.NewImage(img.Back).
				Align(ebitenpkg.AlignBottom).
				Spriteable(ebitenpkg.SpriteSheetOption{
					SpriteWidthCount:  8,
					SpriteHeightCount: 8,
					SpriteHandler: func(fps, timestamp int, direction ebitenpkg.Direction) (indexX int, indexY int, scaleX int, scaleY int) {
						index := (timestamp / 6) % 60
						return index % 8, index / 8, 1, 1
					},
				}).
				Move(250, 950).
				Scale(6, 6)
		}

		battle.PlayerPokemonBattleStatusUI = component.NewPokemonBattleStatusUI(PlayerPokemon[0], battleStatusUIWidth, battleStatusUIHeight, true, func(i ebitenpkg.Image) {
			i.Move(float64(domain.DefaultWindowsWidth())-10, 550)
		})
	}

	if len(OpponentPokemon) != 0 {
		img, err := resource.LoadPokemonImage(OpponentPokemon[0].ID.String())
		if err != nil {
			logs.Errorf("failed to load opponent pokemon image: %v", err)
		} else {
			battle.OpponentPokemonImage = ebitenpkg.NewImage(img.Front).
				Align(ebitenpkg.AlignBottom).
				Spriteable(ebitenpkg.SpriteSheetOption{
					SpriteWidthCount:  8,
					SpriteHeightCount: 8,
					SpriteHandler: func(fps, timestamp int, direction ebitenpkg.Direction) (indexX int, indexY int, scaleX int, scaleY int) {
						index := ((timestamp / 6) + 20) % 60
						return index % 8, index / 8, 1, 1
					},
				}).
				Move(1000, 500).
				Scale(4, 4)
		}

		battle.OpponentPokemonBattleStatusUI = component.NewPokemonBattleStatusUI(OpponentPokemon[0], battleStatusUIWidth, battleStatusUIHeight, false, func(i ebitenpkg.Image) {
			i.Move(float64(battleStatusUIWidth)+10, 10)
		})
	}

	return battle
}

func (b *battle) RestoreGameTime() int {
	return b.GameTime
}

func (b *battle) Update() error {
	return nil
}

func (b *battle) Draw(screen *ebiten.Image) {
	if b.PlayerPokemonImage != nil {
		b.PlayerPokemonImage.Draw(screen)
	}

	if b.PlayerPokemonBattleStatusUI != nil {
		b.PlayerPokemonBattleStatusUI.Draw(screen)
	}

	if b.OpponentPokemonImage != nil {
		b.OpponentPokemonImage.Draw(screen)
	}

	if b.OpponentPokemonBattleStatusUI != nil {
		b.OpponentPokemonBattleStatusUI.Draw(screen)
	}

	b.TextBox.Draw(screen)
}
