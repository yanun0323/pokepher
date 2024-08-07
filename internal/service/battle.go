package service

import (
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
	GameTime int
	Stage    battleStage

	BattleStage *utils.Value[battleStage]

	Player   *entity.Character
	Opponent *entity.Character

	PlayerPokemon   []*entity.Pokemon
	OpponentPokemon []*entity.Pokemon

	PlayerPokemonImage          ebitenpkg.Image
	PlayerPokemonBattleStatusUI usecase.View

	OpponentPokemonImage          ebitenpkg.Image
	OpponentPokemonBattleStatusUI usecase.View
}

func NewBattleStage(Player, Opponent *entity.Character, PlayerPokemon, OpponentPokemon []*entity.Pokemon) usecase.Stage {
	battle := &battle{
		GameTime:        ebitenpkg.CurrentGameTime(),
		BattleStage:     utils.NewValue(_battleStageEntry),
		PlayerPokemon:   PlayerPokemon,
		OpponentPokemon: OpponentPokemon,
	}

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
						index := ((timestamp / 6) + 20) % 60
						return index % 8, index / 8, 1, 1
					},
				})
		}

		battle.PlayerPokemonBattleStatusUI = component.NewPokemonBattleStatusUI(PlayerPokemon[0])
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
				})
		}

		// battle.OpponentPokemonBattleStatusUI = component.NewPokemonBattleStatusUI(OpponentPokemon[0])
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
}
